#!/bin/sh
APP_NAME=go-ssrshare
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build && mv $APP_NAME ./dist

