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
	"time"
)

type rrsData struct {
	Content string
	Prio    int
	Rectype string
	State   string
	Subname string
}
type domainData struct {
	Dname      string
	Result     string
	Rrs        []rrsData
	Service_id string
	Servtype   string
	Soa        map[string]string
}
type answerDomains struct {
	Domains []domainData
}

type dnsRecords struct {
	Answer       answerDomains `json:"answer,omitempty"`
	Charset      string        `json:"charset,omitempty"`
	Messagestore string        `json:"messagestore,omitempty"`
	Result       string        `json:"result,omitempty"`
}

/*
https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go#parsing-json-using-a-struct
*/
type myJSON struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue"`
	NullIntValue    *int      `json:"nullIntValue"`
}

type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

/*
var rawData =
{
   "answer" : {
      "domains" : [
         {
            "dname" : "77699677.xyz",
            "result" : "success",
            "rrs" : [
               {
                  "content" : "194.58.112.174",
                  "prio" : 0,
                  "rectype" : "A",
                  "state" : "A",
                  "subname" : "@"
               },
               {
                  "content" : "ns1.reg.ru.",
                  "prio" : 0,
                  "rectype" : "NS",
                  "state" : "A",
                  "subname" : "@"
               },
               {
                  "content" : "ns2.reg.ru.",
                  "prio" : 1,
                  "rectype" : "NS",
                  "state" : "A",
                  "subname" : "@"
               },
               {
                  "content" : "194.58.112.174",
                  "prio" : 0,
                  "rectype" : "A",
                  "state" : "A",
                  "subname" : "www"
               }
            ],
            "service_id" : "77843261",
            "servtype" : "domain",
            "soa" : {
               "minimum_ttl" : "3h",
               "ttl" : "1d"
            }
         }
      ]
   },
   "charset" : "utf-8",
   "messagestore" : null,
   "result" : "success"
}
*/

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
	//fmt.Printf("%#v\n", res)
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
		//fmt.Println(postData)
		var answer dnsRecords
		b := getZonesPos(apiUrl, postData)
		err := json.Unmarshal([]byte(b), &answer)
		if err != nil {
			fmt.Printf("could not unmarshal json: %s\n", err)
		}
		//fmt.Printf("%+v\n", answer.Answer.Domains[0].Dname)
		fmt.Printf("%+v\n", answer)

	}

}
