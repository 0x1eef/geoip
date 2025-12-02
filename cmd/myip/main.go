package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/0x1eef/myip/pkg/myip"
)

var (
	showHelp    bool
	showVersion bool
)

func main() {
	if showHelp {
		flag.Usage()
	} else if showVersion {
		fmt.Printf("myip: v%s\n", myip.Version)
	} else {
		if res, err := myip.Lookup(); err != nil {
			fmt.Fprintf(os.Stderr, "myip: %v\n", err)
		} else {
			length := padding(res)
			fmt.Printf("%-8s %*s\n", "IP", length, res.IPAddress)
			fmt.Printf("%-8s %*s\n", "Hostname", length, res.Hostname)
			fmt.Printf("%-8s %*s\n", "ISP", length, res.ISP)
			fmt.Printf("%-8s %*s\n", "City", length, res.City)
			fmt.Printf("%-8s %*s\n", "Country", length, res.Country)
		}
	}
}

func padding(res *myip.Response) int {
	base := 8
	lenx := 0
	parts := []string{
		res.IPAddress, res.Hostname,
		res.ISP, res.City,
		res.Country,
	}
	for _, part := range parts {
		if len(part) > lenx {
			lenx = len(part)
		}
	}
	return lenx + base
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "Show help information")
	flag.BoolVar(&showVersion, "v", false, "Show version information")
	flag.Parse()
}
