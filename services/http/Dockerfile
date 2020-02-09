FROM golang:1.13-alpine AS builder

RUN apk add --no-cache git openssh gcc musl-dev

WORKDIR /zah-caddy

COPY ./main.go main.go
COPY ./go.mod go.mod

RUN go get -d -v \
  && go build -o /go/bin/zah-caddy

FROM alpine:edge

COPY --from=builder /go/bin/zah-caddy /go/bin/zah-caddy

CMD ["/go/bin/zah-caddy", "--conf", "/config/Caddyfile", "--log", "stdout", "--agree=true"]

