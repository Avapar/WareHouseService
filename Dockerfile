FROM golang:1.20-alpine

ENV GOROOT /usr/local/go

COPY . /app

WORKDIR /app

RUN go mod download

RUN go build -o main

CMD ["./main"]
