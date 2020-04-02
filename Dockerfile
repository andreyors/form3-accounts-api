FROM golang:1.14-alpine

ENV CGO_ENABLED 0

RUN set -xe \
    && mkdir /app

COPY go.* /app/

WORKDIR /app

RUN set -xe \
    && go mod download

CMD ["go", "test", "-v", "./..."]
