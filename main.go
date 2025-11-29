package geoip

import (
	"encoding/json"
	"io"
	"net/http"
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

func Lookup() (*Response, error) {
	req, err := buildRequest()
	if err != nil {
		return nil, err
	}

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
