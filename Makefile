.PHONY: build
build:
		go build -o UrlCutterApi ./cmd/main.go
		go test ./internal/app/saver ./internal/app/urlcut
.DEFAULT_GOAL := build