FROM golang:alpine AS builder

WORKDIR /go/src/github.com/brcodingdev/chat-service

COPY . .
RUN apk add --no-cache git
RUN go mod tidy
RUN go build -o chatservice ./cmd


FROM alpine:3.16

RUN apk update \
    && apk upgrade

WORKDIR /app
COPY .env /app/.env

ENV DB_HOST=host.docker.internal
ENV RABBIT_HOST=host.docker.internal

COPY --from=builder /go/src/github.com/brcodingdev/chat-service/chatservice .

CMD ["./chatservice"]

EXPOSE 9010

