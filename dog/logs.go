package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"infinity-dog/network"
	"infinity-dog/util"
	"time"
)

var utc, _ = time.LoadLocation("UTC")

func golangTimeToDogTime(s string) string {
	dateString := s[0:10]
	timeString := s[11:19]
	return dateString + "T" + timeString
}

func Logs(hours int, query string) {

	utcNow := time.Now().In(utc)
	// we seem to be off by about 1 hour
	utcNow = utcNow.Add(time.Minute * 55)
	utcString := fmt.Sprintf("%v", utcNow.Add(time.Hour*time.Duration(hours*-1)))
	from := golangTimeToDogTime(utcString)
	utcString = fmt.Sprintf("%v", utcNow.Add(time.Second))
	to := golangTimeToDogTime(utcString)

	cursor := ""
	startTime := time.Now().Unix()
	hits := 0
	for {
		fmt.Println(from, to, cursor)
		payloadString := makePayload(query, from, to, cursor)
		// 300 requests per hour (aka 5 per minute)
		jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))
		hits++
		if hits == 300 {
			for {
				delta := time.Now().Unix() - startTime
				if delta > 3600 {
					break
				}
				fmt.Println("at 300", delta)
				time.Sleep(time.Second * 1)
			}
			startTime = time.Now().Unix()
			hits = 0
		}

		files.SaveFile(fmt.Sprintf("samples/%s.json", util.PseudoUuid()), jsonString)

		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)

		/*
			now := time.Now().Unix()
			for _, d := range logResponse.Data {
				delta := now - d.Attributes.Timestamp.Unix()
				tsFloat := float64(delta) / 60.0
				fmt.Printf("%.2f %s\n", tsFloat, d.Attributes.Service)
				fmt.Printf("%s\n", d.Attributes.Message)
				fmt.Printf("%s\n", d.Attributes.SubAttributes.Msg)
				fmt.Printf("%s\n", d.Attributes.SubAttributes.Exception)
			}*/

		cursor = logResponse.Meta.Page.After

		if cursor == "" {
			break
		}
	}
}

func makePayload(query, from, to, cursor string) string {
	payload := `{
  "filter": {
    "query": "%s",
    "indexes": [
      "main"
    ],
		"from": "%s+01:00",
    "to": "%s+01:00"
  },
  "sort": "timestamp",
  "page": {
	  "cursor": %s,
    "limit": 1000
  }
}`
	cursorString := "null"
	if cursor != "" {
		cursorString = fmt.Sprintf(`"%s"`, cursor)
	}
	return fmt.Sprintf(payload, query, from, to, cursorString)
}
