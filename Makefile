.PHONY: all build run gotool clean help

BINARY="gosible"

all: gotool buildmac

buildlinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  bin/${BINARY} cmd/main.go

buildmac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY} cmd/main.go

buildwin:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o  bin/${BINARY} cmd/main.go

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f bin/${BINARY} ] ; then rm bin/${BINARY} ; fi

help:
	@echo "make"
	@echo "make build"
	@echo "make clean"
	@echo "make gotool"