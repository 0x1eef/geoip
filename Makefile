fmt:
	go fmt ./...

build:
	go build -o bin/myip cmd/myip/main.go

release: fmt build
	strip bin/myip

run:
	go run cmd/myip/main.go