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

var apiUrl = "https://api.reg.ru/api/regru2/"
var apiFunc string
var postData = url.Values{}
var postFields map[string]string

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

func apiPost(reqUrl string, postFields map[string]string) (body []byte) {
	for k, v := range postFields {
		postData.Add(k, v)
	}
	res, err := http.PostForm(reqUrl, postData)
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
	domainName, domain_name_exist := os.LookupEnv("RR_API_DOMAIN_NAME")

	if !username_exist {

		// Print the value of the environment variable
		fmt.Println("RR_API_USERNAME variable is empty")
	} else if !password_exist {
		fmt.Println("RR_API_PASSWORD variable is empty")
	} else if !domain_name_exist {
		fmt.Println("RR_API_DOMAIN_NAME variable is empty")
	} else {
		//https://api.reg.ru/api/regru2/<имя_категории_функции>/<имя_функции>[?<HTTP_параметры_для_запросов_GET>]

		apiFunc = "zone/get_resource_records"

		/* API functions
		zone/get_resource_records
		zone/add_txt
		zone/remove_record
		*/
		postFields = make(map[string]string)
		postFields["username"] = username
		postFields["password"] = password
		postFields["domain_name"] = domainName

		// Now get resource records
		reqUrl := apiUrl + apiFunc
		var answer dnsRecords
		b := apiPost(reqUrl, postFields)
		err := json.Unmarshal(b, &answer)
		if err != nil {
			fmt.Printf("could not unmarshal json: %s\n", err)
		}
		fmt.Printf("%+v\n", answer)
		//fmt.Printf("%+v\n", answer.Answer.Domains[0].Rrs[0].Content)

		/*
			// Now create TXT resource record
			apiFunc = "zone/add_txt"
			reqUrl = apiUrl + apiFunc
			subdomain := "ghm"
			text_record := "foo"
			postData.Add("subdomain", subdomain)
			postData.Add("text", text_record)
			b = getZonesPos(reqUrl, postData)
			fmt.Println(string(b))

		*/
	}

}
