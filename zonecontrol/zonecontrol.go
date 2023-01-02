package zonecontrol

import (
	"io"
	"log"
	"net/url"

	"../connector"
)

//var apiFunc string
//var postFields map[string]string

//Make any POST request with default API settings and return bytes
func ApiRequest(reqUrl string, postFields map[string]string) (body []byte) {
	postData := url.Values{}
	for k, v := range postFields {
		postData.Add(k, v)
	}
	// Take request over connector (may inclde proxy, tmeout settings and so on...)
	c := connector.NewConnection()
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
