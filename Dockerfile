FROM golang:1.22.1-alpine3.19

RUN apk add --no-cache --upgrade bash

RUN apk update && \
    apk add --no-cache openjdk11 ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ENV JAVA_HOME /usr/lib/jvm/java-11-openjdk
ENV PATH $PATH:$JAVA_HOME/bin

WORKDIR /go.tutorial

COPY . .

COPY go.mod go.sum ./
RUN go mod tidy

ARG BUILD_STAGE
ENV BUILD_STAGE $BUILD_STAGE

ARG PROJECT_NAME="go.tutorial"
ENV PROJECT_NAME ${PROJECT_NAME}

ENV TZ=Asia/Bangkok

EXPOSE 8080

RUN go build -o ${PROJECT_NAME}

RUN sed -i 's/\r$//' run.sh
RUN chmod +x 'run.sh'
