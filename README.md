## About

The myip utility provides a simple command-line interface for obtaining your
public IP address, geographical location, and other related information. The
utility uses the free-to-use HTTP service
[clean.myip.wtf](https://clean.myip.wtf).

## Features

* Retrieve public IP address
* Retrieve the following information associated with an IP address:
	* Hostname
	* Internet Service Provider (ISP)
	* City
	* Country

# Install

#### Binaries

Prebuilt binaries are made available via [GItHub actions](https://github.com/0x1eef/myip/actions/runs/19806477120):

* [myip-windows-amd64 (v0.2.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718983018)
* [myip-linux-amd64 (v0.2.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718980945)
* [myip-darwin-amd64 (v0.2.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718980702)
* [myip-freebsd-amd64 (v0.2.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718985996)

#### Package

    user@localhost$ go install github.com/0x1eef/myip/cmd/myip@latest
    user@localhost$ ~/go/bin/myip golang.go

#### Source

    user@localhost$ git clone https://github.com/0x1eef/myip.git
    user@localhost$ cd myip
    user@localhost$ make build
    user@localhost$ ./bin/myip golang.go


## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/myip/pkg/myip"
)

func main() {
	if res, err := myip.Lookup(); err != nil {
		fmt.Fprintf(os.Stderr, "myip: %v\n", err)
	} else {
		fmt.Printf("%-7s %35s\n", "IP", res.IPAddress)
		fmt.Printf("%-7s %35s\n", "ISP", res.ISP)
		fmt.Printf("%-7s %35s\n", "City", res.City)
		fmt.Printf("%-7s %35s\n", "Country", res.Country)
	}
}
```

## Demo

<details>
  <summary>Play</summary>
  <img src="demo.gif" alt="demo" />
</details>

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)


