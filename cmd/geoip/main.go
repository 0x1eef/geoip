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
		fmt.Println(res.IPAddress)
	}
}
