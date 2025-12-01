package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/0x1eef/myip/pkg/myip"
)

var (
	showVersion bool
)

func main() {
	if showVersion {
		fmt.Printf("myip: v%s\n", myip.Version)
	} else {
		if res, err := myip.Lookup(); err != nil {
			fmt.Fprintf(os.Stderr, "myip: %v\n", err)
		} else {
			fmt.Printf("%-7s %35s\n", "IP", res.IPAddress)
			fmt.Printf("%-7s %35s\n", "ISP", res.ISP)
			fmt.Printf("%-7s %35s\n", "City", res.City)
			fmt.Printf("%-7s %35s\n", "Country", res.Country)
		}
	}
}

func init() {
	flag.BoolVar(&showVersion, "v", false, "Show version information")
	flag.Parse()
}
