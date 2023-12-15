FROM golang:1.20-alpine3.16


LABEL maintainer="raudra.nitt@gmail.com"

RUN apk add build-base bash git


WORKDIR /feed-events

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /feed-events

RUN go mod tidy

RUN go build -o /usr/local/bin/feed-events

CMD feed-events

