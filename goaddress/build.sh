#!/bin/bash

# unit test
if ! go test ./lib; then
    echo "test failed"
    exit 1
fi

# build
GOOS=darwin GOARCH=amd64 go build -o goaddress ./main.go
GOOS=darwin GOARCH=amd64 go build -o goaddress_mac ./main.go
shasum -a 256 goaddress_mac
GOOS=linux GOARCH=amd64 go build -o goaddress_linux ./main.go
shasum -a 256 goaddress_linux
GOOS=windows GOARCH=amd64 go build -o goaddress_win ./main.go
shasum -a 256 goaddress_win

mv goaddress "${GOPATH}/bin"
