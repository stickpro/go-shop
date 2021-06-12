.PHONY:
.DEFAULT_GOAL := build

build:
	go build -o ./.bin/app ./cmd/app/main.go
