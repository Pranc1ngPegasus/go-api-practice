FROM golang:1.15-alpine

WORKDIR /app
ARG APP_PORT
ENV GO111MODULE=on

RUN apk update \
 && apk upgrade \
 && apk add ca-certificates alpine-sdk bash gettext \
 && apk --update add tzdata \
 && rm -rf /var/cache/apk/*

COPY . .

EXPOSE $APP_PORT
