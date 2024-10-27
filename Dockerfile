FROM golang:1.22-alpine AS builder

WORKDIR /usr/src/service
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o build/main cmd/service/main.go
RUN go build -o build/kafka cmd/workers/kafka/main.go

FROM alpine:latest

WORKDIR /app

COPY templates ./templates

COPY --from=builder /usr/src/service/build/main .
COPY --from=builder /usr/src/service/build/kafka .

CMD ./main & ./kafka