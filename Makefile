# Copyright (c) 2016-2018 Betalo AB

.PHONY: all
all: build

.PHONY: build
build:
	go build -v github.com/betalo-sweden/await

.PHONY: copyright
copyright:
	@find . -type f -name '*.go' -exec grep -H -m 1 . {} \; \
	    | grep -v '/vendor/' \
	    | (! grep -v "// Copyright (C) .*$$(date +%Y) Betalo AB") \
	    | (! grep -v "// Code generated by .*")

.PHONY: deps
deps:
	go get -u github.com/golang/dep
	dep ensure

.PHONY: lint
lint:
	@if [ $$(gofmt -l . | wc -l) != 0 ]; then \
	    echo "gofmt: code not formatted"; \
	    gofmt -l . | grep -v vendor/; \
	    exit 1; \
	fi

	@gometalinter \
	             --vendor \
	             --tests \
	             --disable=gocyclo \
	             --disable=dupl \
	             --disable=deadcode \
	             --disable=gotype \
	             --disable=maligned \
	             --disable=interfacer \
	             --disable=varcheck \
	             --disable=gas \
	             --disable=megacheck \
	             ./...

.PHONY: test
test:
	go test -v

.PHONY: rel
rel:
	GOOS=darwin GOARCH=amd64 go build -o await-darwin-amd64 github.com/betalo-sweden/await
	GOOS=linux  GOARCH=amd64 go build -o await-linux-amd64  github.com/betalo-sweden/await
