#!/bin/bash

echo "Setting go env for alpine linux"
export GOOS=linux 
export GOARCH=amd64

echo "Getting dependencies for go-check-certs"
go get
go build .

echo "Building docker container"
docker build . -t derfoh/go-check-certs

echo "Cleaning up binary file"
rm go-check-certs