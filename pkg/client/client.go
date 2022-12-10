package client

import (
	"encoding/json"
	"fmt"

	"github.com/daloman/regru-api-go/pkg/zonecontrol"
)

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

type response dnsRecords

const apiUrl = "https://api.reg.ru/api/regru2/"
const zoneGetRrs = "zone/get_resource_records"
const zoneAddTxt = "zone/add_txt"
const zoneRemoveRrs = "zone/remove_record"

//var username, password, domainName string

func GetZones(username, password, domainName string) {
	// Now get resource records
	apiFunc := zoneGetRrs
	reqUrl := apiUrl + apiFunc
	// Create data map for POST request
	postFields := make(map[string]string)
	postFields["username"] = username
	postFields["password"] = password
	postFields["domain_name"] = domainName

	answer := zonecontrol.ApiRequest(reqUrl, postFields)
	unmarshalRsponse(answer)
}

func AddTxtRr(username, password, domainName, subdomain, textBody string) {
	// Now get resource records
	apiFunc := zoneAddTxt
	reqUrl := apiUrl + apiFunc

	postFields := make(map[string]string)
	postFields["username"] = username
	postFields["password"] = password
	postFields["domain_name"] = domainName
	postFields["subdomain"] = subdomain
	postFields["text"] = textBody

	answer := zonecontrol.ApiRequest(reqUrl, postFields)
	unmarshalRsponse(answer)
}

func RmTxtRr(username, password, domainName, subdomain, resourceRecordType string) {
	apiFunc := zoneRemoveRrs
	reqUrl := apiUrl + apiFunc

	postFields := make(map[string]string)
	postFields["username"] = username
	postFields["password"] = password
	postFields["domain_name"] = domainName
	postFields["subdomain"] = subdomain
	postFields["record_type"] = resourceRecordType

	answer := zonecontrol.ApiRequest(reqUrl, postFields)
	unmarshalRsponse(answer)
}

func unmarshalRsponse(rawData []byte) {
	// Print "raw" default api request in json format
	//fmt.Println(string(answer))
	b := response{}
	err := json.Unmarshal(rawData, &b)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
	}
	fmt.Printf("The answer is: %+v\n", b)
}
