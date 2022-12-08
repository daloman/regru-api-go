package client

import (
	"net/http"
	"net/url"
	"time"
)

// Use fromEnv
// Const is for test purposes only
const proxyUrl = "http://192.168.1.100:3128"

// Init Client connection over httpProxy
func NewClient() *http.Client {
	proxy, _ := url.Parse(proxyUrl)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		Proxy:              http.ProxyURL(proxy),
	}
	client := &http.Client{Transport: tr}
	return client
}
