/**
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
**/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type domains struct {
	dname      string
	result     string
	rss        map[string]string
	service_id string
	servtype   string
	soa        map[string]string
}

type dnsRecords struct {
	Answer       map[string][]map[string]string `json:"answer,omitempty"`
	Result       string                         `json:"result,omitempty"`
	Charset      string                         `json:"charset,omitempty"`
	Messagestore string                         `json:"messagestore,omitempty"`
	//Result       string   `json:"answer,omitempty"`
}

func getZonesPos(apiUrl string, postData url.Values) (body []byte) {
	res, err := http.PostForm(apiUrl, postData)
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", body)
	return body
}

func main() {

	// Get environment variables
	username, username_exist := os.LookupEnv("RR_API_USERNAME")
	password, password_exist := os.LookupEnv("RR_API_PASSWORD")
	dbname, dbname_exist := os.LookupEnv("RR_API_DBNAME")

	if !username_exist {

		// Print the value of the environment variable
		fmt.Println("RR_API_USERNAME variable is empty")
	} else if !password_exist {
		fmt.Println("RR_API_PASSWORD variable is empty")
	} else if !dbname_exist {
		fmt.Println("RR_API_DBNAME variable is empty")
	} else {
		apiUrl := "https://api.reg.ru/api/regru2/zone/get_resource_records"
		postData := url.Values{}

		postData.Add("username", username)
		postData.Add("password", password)
		postData.Add("dname", dbname)
		fmt.Println(postData)
		var answer dnsRecords
		b := getZonesPos(apiUrl, postData)
		err := json.Unmarshal([]byte(b), &answer)
		if err != nil {
			fmt.Printf("could not unmarshal json: %s\n", err)
		}
		fmt.Printf("%+v\n", answer)

	}
}
