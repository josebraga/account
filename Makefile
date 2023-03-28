
GIT_ROOT ?= $(shell git rev-parse --show-toplevel)

build:
	go mod download
	go mod tidy
	mkdir -p ${GIT_ROOT}/output
	go build -o ${GIT_ROOT}/output ./cmd/...
	cp -r configs ${GIT_ROOT}/output
