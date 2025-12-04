## About

The myip utility provides a simple command-line interface for obtaining your
public IP address, geographical location, and other related information. The
utility uses the privacy-focused and free-to-use HTTP service known as
[clean.myip.wtf](https://clean.myip.wtf) to fetch information about
your IP address.

## Features

* Retrieve public IP address
* Retrieve hostname associated with an IP address
* Retrieve ISP associated with an IP address
* Retrieve city, and country associated with an IP address

# Install

#### Binaries

Prebuilt binaries are made available via [GitHub actions](https://github.com/0x1eef/myip/actions/runs/19806477120):

* [myip-windows-amd64 (v0.1.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718983018)
* [myip-linux-amd64 (v0.1.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718980945)
* [myip-darwin-amd64 (v0.1.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718980702)
* [myip-freebsd-amd64 (v0.1.0)](https://github.com/0x1eef/myip/actions/runs/19806477120/artifacts/4718985996)

#### Package

    user@localhost$ go install github.com/0x1eef/myip/cmd/myip@latest
    user@localhost$ ~/go/bin/myip

#### Source

    user@localhost$ git clone https://github.com/0x1eef/myip.git
    user@localhost$ cd myip
    user@localhost$ make build
    user@localhost$ ./bin/myip

## Usage

#### Basics

    user@localhost$ myip
    IP                                  169.150.198.74
    Hostname         unn-169-150-198-74.datapacket.com
    ISP                               Datacamp Limited
    City                                     São Paulo
    Country                                     Brazil

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/myip)
* [gitlab.com/@0x1eef](https://gitlab.com/0x1eef/myip)
* [hardenedbsd.org/@0x1eef](https://hardenedbsd.org/@0x1eef/myip)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)


