package zonecontrol

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/daloman/regru-api-go/client"
)

type rrsData struct {
	Content string
	Prio    int
	Rectype string
	State   string
	Subname string
}
type domainData struct {
	Dname       string
	ErrorCode   string
	ErrorText   string
	ErrorParams map[string]string
	Result      string
	Rrs         []rrsData
	ServiceId   string
	Servtype    string
	Soa         map[string]string
}
type answerDomains struct {
	Domains []domainData
}

type dnsRecords struct {
	Answer       answerDomains     `json:"answer,omitempty"`
	Charset      string            `json:"charset,omitempty"`
	Messagestore string            `json:"messagestore,omitempty"`
	Result       string            `json:"result,omitempty"`
	ErrorCode    string            `json:"error_code,omitempty"`
	ErrorText    string            `json:"error_text,omitempty"`
	ErrorParams  map[string]string `json:"error_params,omitempty"`
}

type response dnsRecords

const apiUrl = "https://api.reg.ru/api/regru2/"
const zoneGetRrs = "zone/get_resource_records"
const zoneAddTxt = "zone/add_txt"
const zoneRemoveRrs = "zone/remove_record"

func GetZones(username, password, domainName string) {
	// Now get resource records
	apiFunc := zoneGetRrs
	reqUrl := apiUrl + apiFunc
	// Create data map for POST request
	postFields := make(map[string]string)
	postFields["username"] = username
	postFields["password"] = password
	postFields["domain_name"] = domainName

	answer := client.ApiRequest(reqUrl, postFields)
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

	answer := client.ApiRequest(reqUrl, postFields)
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

	answer := client.ApiRequest(reqUrl, postFields)
	unmarshalRsponse(answer)
}

func unmarshalRsponse(rawData []byte) {
	b := response{}
	err := json.Unmarshal(rawData, &b)
	if err != nil {
		log.Warnf("Could not unmarshal json: %s\n", err)
	}
	log.Printf("The answer is: %+v\n", b)
}
