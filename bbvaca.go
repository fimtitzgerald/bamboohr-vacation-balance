package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//Resp1 is response object for each type of time off
type Resp1 struct {
	TimeOffType string `json:"timeOffType"`
	Balance     string `json:"balance"`
	TimeOffName string `json:"name"`
}

var authKey = os.Getenv("MY_BB_API")
var employeeID = os.Getenv("BB_ID")

func main() {
	currentTime := time.Now().Local()
	baseURL := fmt.Sprintf("https://api.bamboohr.com/api/gateway.php/unbounce/v1/employees/%s/time_off/calculator/?end=%v", employeeID, currentTime.Format("2006-01-02"))

	req, _ := http.NewRequest("GET", baseURL, nil)

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
	for i := range record {
		fmt.Println("Your", record[i].TimeOffName, "balance is:", record[i].Balance, "days")
	}
	//fmt.Println("Your vacation balance is:", record[0].Balance, "days")
}
