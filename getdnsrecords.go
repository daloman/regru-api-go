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
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getZones() {
	res, err := http.Get("https://api.reg.ru/api/regru2/zone/get_resource_records")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	//fmt.Println(res.StatusCode)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

func getZonesPos(apiUrl string, postData url.Values) {
	res, err := http.PostForm(apiUrl, postData)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	//fmt.Println(res.StatusCode)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
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
		//getZones()
		apiUrl := "https://api.reg.ru/api/regru2/zone/get_resource_records"
		postData := url.Values{}

		postData.Add("username", username)
		postData.Add("password", password)
		postData.Add("dname", dbname)
		fmt.Println(postData)
		getZonesPos(apiUrl, postData)
	}
}
