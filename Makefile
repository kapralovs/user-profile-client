.PHONY: build
build:
	go build -v ./cmd/get ./cmd/post

.PHONY: run
run:
	go run ./cmd/post ./cmd/get
	
.DEFAULT_GOAL := build