#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

sh ${SCRIPT_DIR}/setup.sh

export GOOS=linux
export GOARCH=amd64

go build -o bin/viewer ${SCRIPT_DIR}/../cmd/viewer.go
docker build -f docker/web-server-production.yml -t web-server-prod .

docker tag web-server-prod uforic/web-server-prod:latest
docker tag generator uforic/data-generator:latest

docker push uforic/web-server-prod
docker push uforic/data-generator
