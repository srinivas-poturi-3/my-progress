FROM golang:latest

ADD . .
RUN go mod tidy
RUN go build -o app main.go

EXPOSE [":9999"]
ENTRYPOINT ["./app"]
