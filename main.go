package geoip

import (
	"encoding/json"
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
	scheme = "https://"
	host   = "clean.myip.wtf"
	path   = "/json"
)

func Lookup() (*Response, error) {
	req, err := build()
	if err != nil {
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
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}
