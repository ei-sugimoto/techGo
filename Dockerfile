FROM golang:1.27

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy

