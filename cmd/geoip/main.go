package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/geoip"
)

func main() {
	if res, err := geoip.Lookup(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	} else {
		fmt.Printf("%-7s %20s\n", "IP", res.IPAddress)
		fmt.Printf("%-7s %20s\n", "ISP", res.ISP)
		fmt.Printf("%-7s %20s\n", "City", res.City)
		fmt.Printf("%-7s %20s\n", "Country", res.Country)
	}
}
