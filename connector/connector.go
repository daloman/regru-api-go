package connector

import (
	"net/http"
	"net/url"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const proxyUrl = "LOCAL_PROXY"

// NewConnection init Client connection over httpProxy if env variable is defined
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
			log.Errorf("Can't set proxy, due to URL parse error: %v", err)
		} else {
			log.Infof("Using proxy %s=%v", proxyUrl, proxy)
			tr.Proxy = http.ProxyURL(proxy)
		}
	}

	return &http.Client{Transport: tr, Timeout: 10 * time.Second}
}
