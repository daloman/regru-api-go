package regruapigo

import (
	"encoding/json"
	"fmt"

	"github.com/daloman/regru-api-go/pkg/zonecontrol"
)

const apiUrl = "https://api.reg.ru/api/regru2/"
const zoneGetRrs = "zone/get_resource_records"
const zoneAddTxt = "zone/add_txt"
const zoneRemoveRrs = "zone/remove_record"

var username, password, domainName string

/*
func main() {

	// Get environment variables
	// Make test them (Can use tes
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

		//GetZones(username, password, domainName)

		// Add TXT record example:
		//addTxtRr(username, password, domainName, "_acme_oops_002.77699677.xyz.", "_acme_oops_body_002")
		//
		// Remove TXT resource record example
		//rmTxtRr(username, password, domainName, "_acme_oops", "TXT")
	}
}
*/

func GetZones(username, password, domainName string) {
	// Now get resource records
	apiFunc := zoneGetRrs
	reqUrl := apiUrl + apiFunc

	postFields := make(map[string]string)
	postFields["username"] = username
	postFields["password"] = password
	postFields["domain_name"] = domainName

	answer := zonecontrol.ApiRequest(reqUrl, postFields)
	b := zonecontrol.Response{}
	err := json.Unmarshal(answer, &b)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
	}
	fmt.Printf("The answer is: %+v\n", b)
	return
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
	// Print "raw" default api request in json format
	//fmt.Println(string(answer))
	b := zonecontrol.Response{}
	err := json.Unmarshal(answer, &b)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
	}
	fmt.Printf("The answer is: %+v\n", b)
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
	// Print "raw" default api request in json format
	//fmt.Println(string(answer))
	b := zonecontrol.Response{}
	err := json.Unmarshal(answer, &b)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
	}
	fmt.Printf("The answer is: %+v\n", b)
}
