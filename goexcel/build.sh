#!/bin/bash

# build
GOOS=darwin GOARCH=amd64 go build -o goexcel ./main.go
GOOS=darwin GOARCH=amd64 go build -o goexcel_mac ./main.go
shasum -a 256 goexcel_mac
GOOS=linux GOARCH=amd64 go build -o goexcel_linux ./main.go
shasum -a 256 goexcel_linux
GOOS=windows GOARCH=amd64 go build -o goexcel_win ./main.go
shasum -a 256 goexcel_win

mv goexcel ${GOPATH}/bin
