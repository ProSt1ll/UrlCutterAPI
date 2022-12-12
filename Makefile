.PHONY: build
build:
		go build -o UrlCutterApi ./cmd/main.go

.DEFAULT_GOAL := build