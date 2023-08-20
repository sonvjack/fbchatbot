FROM golang:1.21.0-alpine3.18 AS base
WORKDIR /fbchatbot

COPY . /fbchatbot

RUN export GOPROXY=https://goproxy.cn && go build

FROM alpine:latest

WORKDIR /fbchatbot
COPY --from=base /fbchatbot/fbchatbot /fbchatbot/
COPY --from=base /fbchatbot/config.yml /fbchatbot/
CMD ["./fbchatbot"]