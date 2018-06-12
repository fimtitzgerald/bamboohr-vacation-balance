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

// global variables
var authKey = os.Getenv("MY_BB_API")
var employeeID = os.Getenv("BB_ID")

func main() {

	inArgs := os.Args[1:]

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

	if len(inArgs) > 0 {
		for _, v := range inArgs {
			if v == "vaca" {
				fmt.Println("Your vacation balance is:", record[0].Balance, "days")
			} else if v == "lieu" {
				fmt.Println("Your lieu balance is:", record[1].Balance, "days")
			} else if v == "pers" {
				fmt.Println("Your", record[3].TimeOffName, "balance is:", record[3].Balance, "days")
				fmt.Println("Your", record[4].TimeOffName, "balance is:", record[4].Balance, "days")
				fmt.Println("Your", record[5].TimeOffName, "balance is:", record[5].Balance, "days")
			}
		}
	} else {
		for i := range record {
			fmt.Println("Your", record[i].TimeOffName, "balance is:", record[i].Balance, "days")
		}
	}

}
