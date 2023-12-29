NAME := jurassic-park-api

.PHONY: test up

default: up

test:
	go fmt ./...
	go test -v ./...

up:
	docker-compose up --build
