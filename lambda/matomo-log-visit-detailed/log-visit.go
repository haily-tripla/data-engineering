package main

import (
	"database/sql"

	strconvCustom "pkg/util-custom"
)

type logVisit struct {
	Idvisit                   int32  `avro:"idvisit"`
	Idsite                    int32  `avro:"idsite"`
	Site_name                 string `avro:"site_name"`
	Site_main_url             string `avro:"site_main_url"`
	Idvisitor                 string `avro:"idvisitor"`
	Visit_last_action_time    string `avro:"visit_last_action_time"`
	Location_ip               string `avro:"location_ip"`
	User_id                   string `avro:"user_id"`
	Visit_first_action_time   string `avro:"visit_first_action_time"`
	Visit_goal_buyer          bool   `avro:"visit_goal_buyer"`
	Visit_goal_converted      bool   `avro:"visit_goal_converted"`
	Visitor_days_since_first  int32  `avro:"visitor_days_since_first"`
	Visitor_days_since_order  int32  `avro:"visitor_days_since_order"`
	Visitor_returning         bool   `avro:"visitor_returning"`
	Visitor_count_visits      int32  `avro:"visitor_count_visits"`
	Visit_entry_idaction_name int32  `avro:"visit_entry_idaction_name"`
	Visit_entry_idaction_url  int32  `avro:"visit_entry_idaction_url"`
	Visit_exit_idaction_name  int32  `avro:"visit_exit_idaction_name"`
	Visit_exit_idaction_url   int32  `avro:"visit_exit_idaction_url"`
	Visit_total_actions       int32  `avro:"visit_total_actions"`
	Visit_total_interactions  int32  `avro:"visit_total_interactions"`
	Visit_total_searches      int32  `avro:"visit_total_searches"`
	Referer_keyword           string `avro:"referer_keyword"`
	Referer_name              string `avro:"referer_name"`
	Referer_type              int32  `avro:"referer_type"`
	Referer_type_name         string `avro:"referer_type_name"`
	Referer_url               string `avro:"referer_url"`
	Location_browser_lang     string `avro:"location_browser_lang"`
	Browser_lang              string `avro:"browser_lang"`
	Config_browser_engine     string `avro:"config_browser_engine"`
	Config_browser_name       string `avro:"config_browser_name"`
	Config_browser_label      string `avro:"config_browser_label"`
	Config_browser_version    string `avro:"config_browser_version"`
	Config_device_brand       string `avro:"config_device_brand"`
	Config_device_brand_name  string `avro:"config_device_brand_name"`
	Config_browser_family     string `avro:"config_browser_family"`
	Config_device_model       string `avro:"config_device_model"`
	Config_device_type        int32  `avro:"config_device_type"`
	Config_device_type_code   string `avro:"config_device_type_code"`
	Config_device_type_name   string `avro:"config_device_type_name"`
	Config_os                 string `avro:"config_os"`
	Config_os_name            string `avro:"config_os_name"`
	Config_os_family          string `avro:"config_os_family"`
	Config_os_version         string `avro:"config_os_version"`
	Visit_total_events        int32  `avro:"visit_total_events"`
	Visitor_localtime         string `avro:"visitor_localtime"`
	Visitor_days_since_last   int32  `avro:"visitor_days_since_last"`
	Config_resolution         string `avro:"config_resolution"`
	Config_cookie             bool   `avro:"config_cookie"`
	Config_director           bool   `avro:"config_director"`
	Config_flash              bool   `avro:"config_flash"`
	Config_gears              bool   `avro:"config_gears"`
	Config_java               bool   `avro:"config_java"`
	Config_pdf                bool   `avro:"config_pdf"`
	Config_quicktime          bool   `avro:"config_quicktime"`
	Config_realplayer         bool   `avro:"config_realplayer"`
	Config_silverlight        bool   `avro:"config_silverlight"`
	Config_windowsmedia       bool   `avro:"config_windowsmedia"`
	Visit_total_time          int32  `avro:"visit_total_time"`
	Location_city             string `avro:"location_city"`
	Location_country          string `avro:"location_country"`
	Location_latitude         string `avro:"location_latitude"`
	Location_longitude        string `avro:"location_longitude"`
	Location_region           string `avro:"location_region"`
	Custom_var_k1             string `avro:"custom_var_k1"`
	Custom_var_v1             string `avro:"custom_var_v1"`
	Custom_var_k2             string `avro:"custom_var_k2"`
	Custom_var_v2             string `avro:"custom_var_v2"`
	Custom_var_k3             string `avro:"custom_var_k3"`
	Custom_var_v3             string `avro:"custom_var_v3"`
	Custom_var_k4             string `avro:"custom_var_k4"`
	Custom_var_v4             string `avro:"custom_var_v4"`
	Custom_var_k5             string `avro:"custom_var_k5"`
	Custom_var_v5             string `avro:"custom_var_v5"`
	Last_idlink_va            string `avro:"last_idlink_va"`
	Custom_dimension_1        string `avro:"custom_dimension_1"`
	Custom_dimension_2        string `avro:"custom_dimension_2"`
	Custom_dimension_3        string `avro:"custom_dimension_3"`
	Custom_dimension_4        string `avro:"custom_dimension_4"`
	Custom_dimension_5        string `avro:"custom_dimension_5"`
	Campaign_content          string `avro:"campaign_content"`
	Campaign_id               string `avro:"campaign_id"`
	Campaign_keyword          string `avro:"campaign_keyword"`
	Campaign_medium           string `avro:"campaign_medium"`
	Campaign_name             string `avro:"campaign_name"`
	Campaign_source           string `avro:"campaign_source"`
	EtlCreateDateString       string `avro:"etl_create_date_string"`
	EtlUpdateDateString       string `avro:"etl_update_date_string"`
	EtlCreateDateUnix         int64  `avro:"etl_create_date_unix"`
	EtlUpdateDateUnix         int64  `avro:"etl_update_date_unix"`
}

func copyFromLine(lvptr *logVisit, row []sql.RawBytes) {

	idvisit := string(row[0])
	idsite := string(row[1])
	site_name := string(row[2])
	site_main_url := string(row[3])
	idvisitor := string(row[4])
	visit_last_action_time := string(row[5])
	location_ip := string(row[6])
	user_id := string(row[7])
	visit_first_action_time := string(row[8])
	visit_goal_buyer := string(row[9])
	visit_goal_converted := string(row[10])
	visitor_days_since_first := string(row[11])
	visitor_days_since_order := string(row[12])
	visitor_returning := string(row[13])
	visitor_count_visits := string(row[14])
	visit_entry_idaction_name := string(row[15])
	visit_entry_idaction_url := string(row[16])
	visit_exit_idaction_name := string(row[17])
	visit_exit_idaction_url := string(row[18])
	visit_total_actions := string(row[19])
	visit_total_interactions := string(row[20])
	visit_total_searches := string(row[21])
	referer_keyword := string(row[22])
	referer_name := string(row[23])
	referer_type := string(row[24])
	referer_type_name := string(row[25])
	referer_url := string(row[26])
	location_browser_lang := string(row[27])
	browser_lang := string(row[28])
	config_browser_engine := string(row[29])
	config_browser_name := string(row[30])
	config_browser_label := string(row[31])
	config_browser_version := string(row[32])
	config_device_brand := string(row[33])
	config_device_brand_name := string(row[34])
	config_browser_family := string(row[35])
	config_device_model := string(row[36])
	config_device_type := string(row[37])
	config_device_type_code := string(row[38])
	config_device_type_name := string(row[39])
	config_os := string(row[40])
	config_os_name := string(row[41])
	config_os_family := string(row[42])
	config_os_version := string(row[43])
	visit_total_events := string(row[44])
	visitor_localtime := string(row[45])
	visitor_days_since_last := string(row[46])
	config_resolution := string(row[47])
	config_cookie := string(row[48])
	config_director := string(row[49])
	config_flash := string(row[50])
	config_gears := string(row[51])
	config_java := string(row[52])
	config_pdf := string(row[53])
	config_quicktime := string(row[54])
	config_realplayer := string(row[55])
	config_silverlight := string(row[56])
	config_windowsmedia := string(row[57])
	visit_total_time := string(row[58])
	location_city := string(row[59])
	location_country := string(row[60])
	location_latitude := string(row[61])
	location_longitude := string(row[62])
	location_region := string(row[63])
	custom_var_k1 := string(row[64])
	custom_var_v1 := string(row[65])
	custom_var_k2 := string(row[66])
	custom_var_v2 := string(row[67])
	custom_var_k3 := string(row[68])
	custom_var_v3 := string(row[69])
	custom_var_k4 := string(row[70])
	custom_var_v4 := string(row[71])
	custom_var_k5 := string(row[72])
	custom_var_v5 := string(row[73])
	last_idlink_va := string(row[74])
	custom_dimension_1 := string(row[75])
	custom_dimension_2 := string(row[76])
	custom_dimension_3 := string(row[77])
	custom_dimension_4 := string(row[78])
	custom_dimension_5 := string(row[79])
	campaign_content := string(row[80])
	campaign_id := string(row[81])
	campaign_keyword := string(row[82])
	campaign_medium := string(row[83])
	campaign_name := string(row[84])
	campaign_source := string(row[85])
	etl_create_date_string := string(row[86])
	etl_update_date_string := string(row[87])
	etl_create_date_unix := string(row[88])
	etl_update_date_unix := string(row[89])

	if idvisit != "" {
		lvptr.Idvisit = strconvCustom.Atoint32(idvisit)
	}
	if idsite != "" {
		lvptr.Idsite = strconvCustom.Atoint32(idsite)
	}
	if site_name != "" {
		lvptr.Site_name = site_name
	}
	if site_main_url != "" {
		lvptr.Site_main_url = site_main_url
	}
	if idvisitor != "" {
		lvptr.Idvisitor = idvisitor
	}
	if visit_last_action_time != "" {
		lvptr.Visit_last_action_time = visit_last_action_time
	}
	if location_ip != "" {
		lvptr.Location_ip = location_ip
	}
	if user_id != "" {
		lvptr.User_id = user_id
	}
	if visit_first_action_time != "" {
		lvptr.Visit_first_action_time = visit_first_action_time
	}
	if visit_goal_buyer != "" {
		lvptr.Visit_goal_buyer = strconvCustom.Atobool(visit_goal_buyer)
	}
	if visit_goal_converted != "" {
		lvptr.Visit_goal_converted = strconvCustom.Atobool(visit_goal_converted)
	}
	if visitor_days_since_first != "" {
		lvptr.Visitor_days_since_first = strconvCustom.Atoint32(visitor_days_since_first)
	}
	if visitor_days_since_order != "" {
		lvptr.Visitor_days_since_order = strconvCustom.Atoint32(visitor_days_since_order)
	}
	if visitor_returning != "" {
		lvptr.Visitor_returning = strconvCustom.Atobool(visitor_returning)
	}
	if visitor_count_visits != "" {
		lvptr.Visitor_count_visits = strconvCustom.Atoint32(visitor_count_visits)
	}
	if visit_entry_idaction_name != "" {
		lvptr.Visit_entry_idaction_name = strconvCustom.Atoint32(visit_entry_idaction_name)
	}
	if visit_entry_idaction_url != "" {
		lvptr.Visit_entry_idaction_url = strconvCustom.Atoint32(visit_entry_idaction_url)
	}
	if visit_exit_idaction_name != "" {
		lvptr.Visit_exit_idaction_name = strconvCustom.Atoint32(visit_exit_idaction_name)
	}
	if visit_exit_idaction_url != "" {
		lvptr.Visit_exit_idaction_url = strconvCustom.Atoint32(visit_exit_idaction_url)
	}
	if visit_total_actions != "" {
		lvptr.Visit_total_actions = strconvCustom.Atoint32(visit_total_actions)
	}
	if visit_total_interactions != "" {
		lvptr.Visit_total_interactions = strconvCustom.Atoint32(visit_total_interactions)
	}
	if visit_total_searches != "" {
		lvptr.Visit_total_searches = strconvCustom.Atoint32(visit_total_searches)
	}
	if referer_keyword != "" {
		lvptr.Referer_keyword = referer_keyword
	}
	if referer_name != "" {
		lvptr.Referer_name = referer_name
	}
	if referer_type != "" {
		lvptr.Referer_type = strconvCustom.Atoint32(referer_type)
	}
	if referer_type_name != "" {
		lvptr.Referer_type_name = referer_type_name
	}
	if referer_url != "" {
		lvptr.Referer_url = referer_url
	}
	if location_browser_lang != "" {
		lvptr.Location_browser_lang = location_browser_lang
	}
	if browser_lang != "" {
		lvptr.Browser_lang = browser_lang
	}
	if config_browser_engine != "" {
		lvptr.Config_browser_engine = config_browser_engine
	}
	if config_browser_name != "" {
		lvptr.Config_browser_name = config_browser_name
	}
	if config_browser_label != "" {
		lvptr.Config_browser_label = config_browser_label
	}
	if config_browser_version != "" {
		lvptr.Config_browser_version = config_browser_version
	}
	if config_device_brand != "" {
		lvptr.Config_device_brand = config_device_brand
	}
	if config_device_brand_name != "" {
		lvptr.Config_device_brand_name = config_device_brand_name
	}
	if config_browser_family != "" {
		lvptr.Config_browser_family = config_browser_family
	}
	if config_device_model != "" {
		lvptr.Config_device_model = config_device_model
	}
	if config_device_type != "" {
		lvptr.Config_device_type = strconvCustom.Atoint32(config_device_type)
	}
	if config_device_type_code != "" {
		lvptr.Config_device_type_code = config_device_type_code
	}
	if config_device_type_name != "" {
		lvptr.Config_device_type_name = config_device_type_name
	}
	if config_os != "" {
		lvptr.Config_os = config_os
	}
	if config_os_name != "" {
		lvptr.Config_os_name = config_os_name
	}
	if config_os_family != "" {
		lvptr.Config_os_family = config_os_family
	}
	if config_os_version != "" {
		lvptr.Config_os_version = config_os_version
	}
	if visit_total_events != "" {
		lvptr.Visit_total_events = strconvCustom.Atoint32(visit_total_events)
	}
	if visitor_localtime != "" {
		lvptr.Visitor_localtime = visitor_localtime
	}
	if visitor_days_since_last != "" {
		lvptr.Visitor_days_since_last = strconvCustom.Atoint32(visitor_days_since_last)
	}
	if config_resolution != "" {
		lvptr.Config_resolution = config_resolution
	}
	if config_cookie != "" {
		lvptr.Config_cookie = strconvCustom.Atobool(config_cookie)
	}
	if config_director != "" {
		lvptr.Config_director = strconvCustom.Atobool(config_director)
	}
	if config_flash != "" {
		lvptr.Config_flash = strconvCustom.Atobool(config_flash)
	}
	if config_gears != "" {
		lvptr.Config_gears = strconvCustom.Atobool(config_gears)
	}
	if config_java != "" {
		lvptr.Config_java = strconvCustom.Atobool(config_java)
	}
	if config_pdf != "" {
		lvptr.Config_pdf = strconvCustom.Atobool(config_pdf)
	}
	if config_quicktime != "" {
		lvptr.Config_quicktime = strconvCustom.Atobool(config_quicktime)
	}
	if config_realplayer != "" {
		lvptr.Config_realplayer = strconvCustom.Atobool(config_realplayer)
	}
	if config_silverlight != "" {
		lvptr.Config_silverlight = strconvCustom.Atobool(config_silverlight)
	}
	if config_windowsmedia != "" {
		lvptr.Config_windowsmedia = strconvCustom.Atobool(config_windowsmedia)
	}
	if visit_total_time != "" {
		lvptr.Visit_total_time = strconvCustom.Atoint32(visit_total_time)
	}
	if location_city != "" {
		lvptr.Location_city = location_city
	}
	if location_country != "" {
		lvptr.Location_country = location_country
	}
	if location_latitude != "" {
		lvptr.Location_latitude = location_latitude
	}
	if location_longitude != "" {
		lvptr.Location_longitude = location_longitude
	}
	if location_region != "" {
		lvptr.Location_region = location_region
	}
	if custom_var_k1 != "" {
		lvptr.Custom_var_k1 = custom_var_k1
	}
	if custom_var_v1 != "" {
		lvptr.Custom_var_v1 = custom_var_v1
	}
	if custom_var_k2 != "" {
		lvptr.Custom_var_k2 = custom_var_k2
	}
	if custom_var_v2 != "" {
		lvptr.Custom_var_v2 = custom_var_v2
	}
	if custom_var_k3 != "" {
		lvptr.Custom_var_k3 = custom_var_k3
	}
	if custom_var_v3 != "" {
		lvptr.Custom_var_v3 = custom_var_v3
	}
	if custom_var_k4 != "" {
		lvptr.Custom_var_k4 = custom_var_k4
	}
	if custom_var_v4 != "" {
		lvptr.Custom_var_v4 = custom_var_v4
	}
	if custom_var_k5 != "" {
		lvptr.Custom_var_k5 = custom_var_k5
	}
	if custom_var_v5 != "" {
		lvptr.Custom_var_v5 = custom_var_v5
	}
	if last_idlink_va != "" {
		lvptr.Last_idlink_va = last_idlink_va
	}
	if custom_dimension_1 != "" {
		lvptr.Custom_dimension_1 = custom_dimension_1
	}
	if custom_dimension_2 != "" {
		lvptr.Custom_dimension_2 = custom_dimension_2
	}
	if custom_dimension_3 != "" {
		lvptr.Custom_dimension_3 = custom_dimension_3
	}
	if custom_dimension_4 != "" {
		lvptr.Custom_dimension_4 = custom_dimension_4
	}
	if custom_dimension_5 != "" {
		lvptr.Custom_dimension_5 = custom_dimension_5
	}
	if campaign_content != "" {
		lvptr.Campaign_content = campaign_content
	}
	if campaign_id != "" {
		lvptr.Campaign_id = campaign_id
	}
	if campaign_keyword != "" {
		lvptr.Campaign_keyword = campaign_keyword
	}
	if campaign_medium != "" {
		lvptr.Campaign_medium = campaign_medium
	}
	if campaign_name != "" {
		lvptr.Campaign_name = campaign_name
	}
	if campaign_source != "" {
		lvptr.Campaign_source = campaign_source
	}
	if etl_create_date_string != "" {
		lvptr.EtlCreateDateString = etl_create_date_string
	}
	if etl_update_date_string != "" {
		lvptr.EtlUpdateDateString = etl_update_date_string
	}
	if etl_create_date_unix != "" {
		lvptr.EtlCreateDateUnix = strconvCustom.Atoint64(etl_create_date_unix)
	}
	if etl_update_date_unix != "" {
		lvptr.EtlUpdateDateUnix = strconvCustom.Atoint64(etl_update_date_unix)
	}

}
