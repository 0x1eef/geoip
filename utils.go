package geoip

import (
	"net/http"
	"strings"
)

const (
	scheme = "https://"
	host   = "clean.myip.wtf"
	path   = "/json"
)

func buildRequest() (*http.Request, error) {
	endpoint := strings.Join([]string{scheme, host, path}, "")
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}
