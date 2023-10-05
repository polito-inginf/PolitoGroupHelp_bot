FROM golang:1.21.0

WORKDIR /service

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY entrypoints/telegramConnection/main.go main/main.go
COPY utils utils
COPY .env .env
COPY telegramConnection telegramConnection

ENTRYPOINT ["go", "run", "/service/main/main.go"]
