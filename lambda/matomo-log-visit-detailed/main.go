package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pkg/aws/secretsmanager"
	"time"

	"text/template"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	servicelambda "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "github.com/go-sql-driver/mysql"
	avro "gopkg.in/avro.v0"
)

const AWS_REGION = "ap-northeast-1"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type AthenaConfig struct {
	Region         string `json:"region"`
	QueryString    string `json:"querystring"`
	Database       string `json:"database"`
	OutputLocation string `json:"outputlocation"`
}

func getLatestIdvisit(session *session.Session) string {

	fmt.Println("getLatestIdvisit start")
	payload := AthenaConfig{
		Region:         AWS_REGION,
		QueryString:    "select max(idvisit) as max_idvisit from sampledb.matomo_log_visit_detailed",
		Database:       "sampledb",
		OutputLocation: "s3://biwako-stg/tmp/athena-results/matomo/log_visit_detailed/",
	}
	jsonBytes, _ := json.Marshal(payload)

	input := &servicelambda.InvokeInput{
		FunctionName:   aws.String("arn:aws:lambda:ap-northeast-1:944180824657:function:runAthenaQuery"),
		Payload:        jsonBytes,
		InvocationType: aws.String("RequestResponse"),
		LogType:        aws.String("None"),
	}
	svc := servicelambda.New(session)
	req, resp := svc.InvokeRequest(input)
	err := req.Send()
	check(err)

	fmt.Println("lambda: status code", resp.StatusCode)
	var op athena.GetQueryResultsOutput

	err = json.Unmarshal(resp.Payload, &op)
	check(err)

	latestVisitorId := *op.ResultSet.Rows[1].Data[0].VarCharValue
	fmt.Println("Max visitor id", latestVisitorId)
	return latestVisitorId

}

func geDetailedLogVisitSQL(sqlTemplateFile string, latestIdvisit string) string {
	var sqlBuf bytes.Buffer
	sqlText, err := ioutil.ReadFile(sqlTemplateFile)
	check(err)

	t := template.New("athena-matomo-log-visit")
	t, _ = t.Parse(string(sqlText))

	vars := make(map[string]interface{})
	vars["idvisit"] = latestIdvisit

	t.Execute(&sqlBuf, vars)

	sqlQuery := sqlBuf.String()

	log.Print(sqlQuery)
	return sqlQuery
}

func generateLogVisitAvroData(dbClient *sql.DB, outputFile string, sqlQuery string, avroSchemaFile string) {

	f, err := os.Create(outputFile)
	check(err)
	defer f.Close()

	avroSchema, err := avro.ParseSchemaFile(avroSchemaFile)
	check(err)

	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(avroSchema)

	testWriter, err := avro.NewDataFileWriter(f, avroSchema, writer)
	check(err)

	fmt.Println("query", sqlQuery)
	rows, err := dbClient.Query(sqlQuery)
	check(err)

	columns, err := rows.Columns()
	check(err)

	values := make([]sql.RawBytes, len(columns))
	args := make([]interface{}, len(values))
	for i := range values {
		args[i] = &values[i]
	}

	var buf bytes.Buffer
	for rows.Next() {
		f.Sync()
		var lv = &logVisit{}
		err = rows.Scan(args...)
		check(err)

		copyFromLine(lv, values)

		encoder := avro.NewBinaryEncoder(&buf)
		err = writer.Write(lv, encoder)

		check(err)
		testWriter.Write(lv)
		testWriter.Flush()
	}
	testWriter.Close()

	log.Print("Generating Log Visit Avro data")
}

func uploadLogVisitToS3(bucket string, outputFile string, key string, awsSession *session.Session) {

	currentTime := time.Now()
	todayskey := fmt.Sprintf(key, currentTime.Year(), int(currentTime.Month()), currentTime.Day(), currentTime.Format("2006-01-02T15:04:05.999999"))

	fmt.Println(bucket, todayskey)

	f, err := os.Open(outputFile)
	check(err)

	s3Session, err := s3.New(awsSession).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(todayskey),
		Body:   f,
	})
	check(err)

	fmt.Printf("file uploaded to, %s\n", s3Session)

}

func LambdaHandler() (int, error) {
	// init config

	awsCredentials := credentials.NewEnvCredentials()
	fmt.Println("awsCredentials")

	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: awsCredentials,
	})

	secretName := "prod/matomo"
	config := secretsmanager.GetSecretRds(awsSession, secretName)

	latestIdvisit := getLatestIdvisit(awsSession)

	avroSchemaFile := "/tmp/matomo_log_visit_detailed.avsc"
	schemaLocation := "code/matomo/matomo-log-visit-detailed/template/matomo_log_visit_detailed.avsc"
	outputFile := "/tmp/matomo_log_visit_detailed.avro"
	sqlLocation := "code/matomo/matomo-log-visit-detailed/sql/log_visit_detailed.sql"
	sqlFile := "/tmp/log_visit_detailed.sql"

	f, err := os.Create(sqlFile)
	if err != nil {
		fmt.Println("failed to create file")
		panic(err)
	}
	defer f.Close()

	downloader := s3manager.NewDownloader(awsSession)
	numBytes, err := downloader.Download(f,
		&s3.GetObjectInput{
			Bucket: aws.String("biwako-stg"),
			Key:    aws.String(sqlLocation)})
	if err != nil {
		fmt.Println("unable to download item ", sqlLocation, err)
		panic(err)
	}

	fmt.Println("downloaded from s3", sqlLocation, numBytes, "bytes", sqlFile)

	sqlQuery := geDetailedLogVisitSQL(sqlFile, latestIdvisit)

	f2, err := os.Create(avroSchemaFile)
	if err != nil {
		fmt.Println("failed to create file")
		panic(err)
	}
	defer f2.Close()
	downloader2 := s3manager.NewDownloader(awsSession)
	numBytes, err = downloader2.Download(f2,
		&s3.GetObjectInput{
			Bucket: aws.String("biwako-stg"),
			Key:    aws.String(schemaLocation)})

	if err != nil {
		fmt.Println("Unable to download item ", schemaLocation, err)
		panic(err)
	}
	fmt.Println("Downloaded from s3", schemaLocation, numBytes, "bytes", avroSchemaFile)

	// Create the MySQL DNS string for the DB connection
	authToken, err := rdsutils.BuildAuthToken(config.Host+":"+config.Port,
		AWS_REGION,
		"lambda",
		awsCredentials)

	check(err)
	fmt.Println("built auth token")

	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowCleartextPasswords=true&tls=true",
		"lambda",
		authToken,
		config.Host,
		config.Port,
		config.Dbname,
	)

	fmt.Println("connection string:", connectStr)
	dbClient, err := sql.Open("mysql", connectStr)
	check(err)
	defer dbClient.Close()

	fmt.Println("db connected")
	generateLogVisitAvroData(dbClient, outputFile, sqlQuery, avroSchemaFile)

	bucketname := "biwako-stg"
	s3OutputLocation := "data/processed/matomo/matomo_log_visit/%d/%d/%d/detailed_log_visit_%s.avro"

	uploadLogVisitToS3(bucketname, outputFile, s3OutputLocation, awsSession)

	return 0, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
