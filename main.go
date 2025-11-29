package geoip

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Response struct {
	IPAddress   string `json:"YourIPaddress"`
	Hostname    string `json:"YourHostname"`
	Location    string `json:"YourLocation"`
	City        string `json:"YourCity"`
	Country     string `json:"YourCountry"`
	CountryCode string `json:"YourCountryCode"`
	ISP         string `json:"YourISP"`
	IsTorExit   bool   `json:"YourTorExit"`
}

const (
	scheme  = "https://"
	host    = "clean.myip.wtf"
	path    = "/json"
	version = "0.1.0"
)

var (
	headers = map[string]string{
		"Accept":     "application/json",
		"User-Agent": fmt.Sprintf("geoip/%s", version),
	}
)

func Lookup() (*Response, error) {
	if req, err := build(); err != nil {
		return nil, err
	} else {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var r Response
		err = json.Unmarshal(body, &r)
		if err != nil {
			return nil, err
		}
		return &r, nil
	}
}

func build() (*http.Request, error) {
	endpoint := strings.Join([]string{scheme, host, path}, "")
	if req, err := http.NewRequest("GET", endpoint, nil); err != nil {
		return nil, err
	} else {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		return req, nil
	}
}
