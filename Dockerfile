FROM golang:latest

COPY ./* .

CMD "go run main.go"