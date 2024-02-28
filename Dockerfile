FROM golang:1.18.10 AS builder

WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download

RUN go build -o loghacker

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/loghacker /app/

CMD ["sleep","infinity"]

