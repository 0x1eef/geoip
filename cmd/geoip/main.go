package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/geoip"
)

func main() {
	if res, err := geoip.Lookup(); err != nil {
		fmt.Fprintf(os.Stderr, "geoip: %v\n", err)
	} else {
		fmt.Printf("%-7s %35s\n", "IP", res.IPAddress)
		fmt.Printf("%-7s %35s\n", "ISP", res.ISP)
		fmt.Printf("%-7s %35s\n", "City", res.City)
		fmt.Printf("%-7s %35s\n", "Country", res.Country)
	}
}
