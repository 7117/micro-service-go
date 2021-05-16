
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:

	go build -o micro_web_test-web main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t micro_web_test-web:latest
