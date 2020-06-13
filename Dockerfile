FROM golang:alpine as builder

RUN apk add --no-cache git build-base
WORKDIR /src
COPY . /src
RUN git checkout master && \
    go mod download && \
    make linux-amd64 && \
    mv ./bin/bot-linux-amd64 /bot

FROM alpine:latest

COPY --from=builder /bot /
ENTRYPOINT ["/bot"]