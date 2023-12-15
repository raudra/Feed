FROM golang:1.20-alpine3.16


LABEL maintainer="raudra.nitt@gmail.com"

RUN apk add build-base bash git


WORKDIR /feed-manager

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /feed-manager

RUN go mod tidy

RUN go build -o /usr/local/bin/feed-manager

CMD feed-manager

