package zonecontrol

import (
	"io"
	"log"
	"net/url"
	"regru-api-go/pkg/client"
)

var apiFunc string
var postFields map[string]string

type rrsData struct {
	Content string
	Prio    int
	Rectype string
	State   string
	Subname string
}
type domainData struct {
	Dname        string
	Error_code   string
	Error_text   string
	Error_params map[string]string
	Result       string
	Rrs          []rrsData
	Service_id   string
	Servtype     string
	Soa          map[string]string
}
type answerDomains struct {
	Domains []domainData
}

type dnsRecords struct {
	Answer       answerDomains     `json:"answer,omitempty"`
	Charset      string            `json:"charset,omitempty"`
	Messagestore string            `json:"messagestore,omitempty"`
	Result       string            `json:"result,omitempty"`
	Error_code   string            `json:"error_code,omitempty"`
	Error_text   string            `json:"error_text,omitempty"`
	Error_params map[string]string `json:"error_params,omitempty"`
}

const apiUrl = "https://api.reg.ru/api/regru2/"
const zoneGetRrs = "zone/get_resource_records"
const zoneAddTxt = "zone/add_txt"

type Response dnsRecords

//Make any POST request with default API settings and return bytes
func ApiRequest(reqUrl string, postFields map[string]string) (body []byte) {
	postData := url.Values{}
	for k, v := range postFields {
		postData.Add(k, v)
	}

	c := client.NewClient()
	res, err := c.PostForm(reqUrl, postData)
	if err != nil {
		log.Fatal(err.Error())
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
