FROM golang:1.18

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download


#RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go test ./internal/app/saver ./internal/app/urlcut

RUN go test ./internal/app/urlcut

RUN go build -o UrlCutterApi ./cmd/main.go

#RUN goose -dir ./migrations postgres  "user=postgres password=medusa dbname=postgres sslmode=disable" up
CMD ["./UrlCutterApi"]
