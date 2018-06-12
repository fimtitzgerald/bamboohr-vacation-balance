package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Resp1 struct {
	TimeOffType string `json:"timeOffType`
	Balance     string `json:"balance"`
}

var authKey = os.Getenv("MY_BB_API")
var employeeId = os.Getenv("BB_ID")

func main() {
	base_url := fmt.Sprintf("https://api.bamboohr.com/api/gateway.php/unbounce/v1/employees/%s/time_off/calculator/?end=2018-06-11", employeeId)

	req, _ := http.NewRequest("GET", base_url, nil)

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(authKey, "")
	req.Header.Add("Cache-Control", "no-cache")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	content, _ := ioutil.ReadAll(resp.Body)

	var record []Resp1

	err = json.Unmarshal(content, &record)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(record[0].TimeOffType)
	fmt.Println(record[0].Balance)
}
