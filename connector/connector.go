package connector

import (
	"net/http"
	"time"
)

// NewConnection init Client connection using httpProxy
// if env variable is defined.
func NewConnection() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		Proxy:              http.ProxyFromEnvironment,
	}

	return &http.Client{Transport: tr, Timeout: 10 * time.Second}
}
