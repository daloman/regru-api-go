package connector

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const proxyUrl = "LOCAL_PROXY"

// Init Client connection over httpProxy
func NewConnection() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	proxyAddress, proxyDefined := os.LookupEnv(proxyUrl)
	if proxyDefined && proxyAddress != "" {
		proxy, err := url.Parse(proxyAddress)
		if err != nil {
			log.Printf("Can't set proxy, due to URL parse error: %v", err)
		} else {
			log.Printf("Using proxy %s=%v", proxyUrl, proxy)
			tr.Proxy = http.ProxyURL(proxy)
		}
	}

	return &http.Client{Transport: tr, Timeout: 10 * time.Second}
}
