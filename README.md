## About

The geoip library is a Go library and application that provides a simple
interface for obtaining your public IP address, geographical location,
and other related information. The library uses the free-to-use HTTP
service [clean.myip.wtf](https://clean.myip.wtf).

## Example

```go
package main

import (
  "fmt"

  "github.com/0x1eef/geoip"
)

func main() {
  res, err := geoip.Lookup()
  if err != nil {
    panic(err)
  } else {
    fmt.Printf("%-7s %20s\n", "IP", res.IPAddress)
    fmt.Printf("%-7s %20s\n", "ISP", res.ISP)
    fmt.Printf("%-7s %20s\n", "City", res.City)
    fmt.Printf("%-7s %20s\n", "Country", res.Country)
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


