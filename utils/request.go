package utils

import (
	"net/http"
	"time"
)

var clientRequest *http.Client

func GetClientHttp() *http.Client {
	if clientRequest == nil {
		var tr = &http.Transport{
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 100,
			TLSHandshakeTimeout: 1000 * time.Second,
		}

		clientRequest = &http.Client{Transport: tr, Timeout: 5 * time.Second}
		GetLogger().Info("Request client initiated.")
	}

	return clientRequest

}
