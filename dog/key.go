package dog

import (
	"fmt"
	"infinity-dog/network"
)

func CheckKey() {
	jsonString := network.DoGet("/api/v1/validate")
	fmt.Println(jsonString)
}

func CreateApplicationKey() {
	input := `{
  "data": {
    "type": "application_keys",
    "attributes": {
      "name": "aa_usage_read8",
			"scopes": ["usage_read", "timeseries_query","events_read","incident_read"]
    }
  }
}`
	jsonString := network.DoPost("/api/v2/current_user/application_keys", []byte(input))
	fmt.Println(jsonString)
}

func CreateApiKey() {
	input := `{
  "data": {
    "type": "api_keys",
    "attributes": {
      "name": "aa_usage_read",
			"scopes": ["usage_read"]
    }
  }
}`
	jsonString := network.DoPost("/api/v2/api_keys", []byte(input))
	fmt.Println(jsonString)
}
