package main

import (
	"github.com/0x1eef/geoip"
)

func main() {
	res, err := geoip.Lookup()
	if err != nil {
		panic(err)
	} else {
		println(res.IPAddress)
	}
}
