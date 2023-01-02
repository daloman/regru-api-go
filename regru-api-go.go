package regru-api-go

import (
	"fmt"
	"os"

	"github.com/daloman/regru-api-go/pkg/client"
)

//const apiUrl = "https://api.reg.ru/api/regru2/"
//const zoneGetRrs = "zone/get_resource_records"
//const zoneAddTxt = "zone/add_txt"
//const zoneRemoveRrs = "zone/remove_record"

//var username, password, domainName string

func zoneInfo() {

	// Get environment variables
	// Make test them (Can use tes
	username, username_exist := os.LookupEnv("RR_API_USERNAME")
	password, password_exist := os.LookupEnv("RR_API_PASSWORD")
	domainName, domain_name_exist := os.LookupEnv("RR_API_DOMAIN_NAME")

	// Check essential variables defined
	if !username_exist {
		fmt.Println("RR_API_USERNAME variable is empty")
	} else if !password_exist {
		fmt.Println("RR_API_PASSWORD variable is empty")
	} else if !domain_name_exist {
		fmt.Println("RR_API_DOMAIN_NAME variable is empty")
	} else {
		// Simple request to take zones
		client.GetZones(username, password, domainName)
	}

}
