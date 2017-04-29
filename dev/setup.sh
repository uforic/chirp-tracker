#!/bin/bash
rm -rf bin/
mkdir bin

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export GOOS=linux
export GOARCH=amd64

go get -d ./...
go build -o bin/goreman  github.com/mattn/goreman
go build -o bin/data-generator ${SCRIPT_DIR}/../cmd/data-generator.go

docker build -f docker/data-generator.yml -t data-generator .
docker build -f docker/web-server.yml -t web-server .
