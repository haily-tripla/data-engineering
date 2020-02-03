package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

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

func HandleRequest(ctx context.Context, config AthenaConfig) (*athena.GetQueryResultsOutput, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc.Identity.CognitoIdentityPoolID)
	log.Print(lc.Identity.CognitoIdentityID)

	log.Println("athena config", config)
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Region)})
	check(err)
	fmt.Println("create aws session", awsSession)
	svc := athena.New(awsSession)

	fmt.Println("create svc", svc)
	var a athena.StartQueryExecutionInput
	a.SetQueryString(config.QueryString)

	var q athena.QueryExecutionContext
	q.SetDatabase(config.Database)
	a.SetQueryExecutionContext(&q)
	fmt.Println("setting QueryExecutionContext")

	var r athena.ResultConfiguration
	r.SetOutputLocation(config.OutputLocation)
	a.SetResultConfiguration(&r)
	fmt.Println("setting ResultConfiguration")
	result, err := svc.StartQueryExecution(&a)
	check(err)

	fmt.Println("startQueryExecution result:")
	fmt.Println(result.GoString())

	var qri athena.GetQueryExecutionInput
	qri.SetQueryExecutionId(*result.QueryExecutionId)

	var qrop *athena.GetQueryExecutionOutput
	duration := time.Duration(2) * time.Second // Pause for 2 seconds

	for {
		qrop, err = svc.GetQueryExecution(&qri)
		check(err)

		if *qrop.QueryExecution.Status.State != "RUNNING" {
			break
		}
		fmt.Println("waiting.")
		time.Sleep(duration)

	}
	if *qrop.QueryExecution.Status.State == "SUCCEEDED" {

		var ip athena.GetQueryResultsInput
		ip.SetQueryExecutionId(*result.QueryExecutionId)

		op, err := svc.GetQueryResults(&ip)
		check(err)

		fmt.Println("QueryResults", *op)

		//athenaResultSet := *op.ResultSet.Rows[1].Data[0].VarCharValue
		//fmt.Println("result set: ", athenaResultSet)
		return op, nil
	} else {
		fmt.Println(*qrop.QueryExecution.Status.State)
	}
	return nil, err
}

func main() {
	lambda.Start(HandleRequest)
}
