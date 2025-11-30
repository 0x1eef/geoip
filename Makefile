fmt:
	go fmt ./...

build:
	go build -o bin/geoip cmd/geoip/main.go

release: fmt build
	strip bin/geoip

run:
	go run cmd/geoip/main.go