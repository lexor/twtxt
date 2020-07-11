.PHONY: dev build install image release test clean

CGO_ENABLED=0
COMMIT=$(shell git rev-parse --short HEAD)

all: dev

dev: build
	@./twtd -D

build: generate
	@go build -tags "netgo static_build" -installsuffix netgo \
		-ldflags "-w -X $(shell go list)/.GitCommit=$(COMMIT)" \
		./cmd/twtd/...

generate:
	@rice embed-go

install: build
	@go install

image:
	@docker build -t prologic/twtxt .

release:
	@./tools/release.sh

test:
	@go test -v -cover -race .

clean:
	@git clean -f -d -X
