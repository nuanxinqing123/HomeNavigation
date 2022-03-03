FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

MAINTAINER HomeNavigation "nuanxinqing@gmail.com"

WORKDIR $GOPATH/src/Gin_HomeNavigation

COPY . .

ADD . ./

#增加缺失的包，移除没用的包
RUN go mod tidy

RUN go build -o main .

FROM alpine:latest

COPY --from=builder go/src/Gin_HomeNavigation/conf /app/conf
COPY --from=builder go/src/Gin_HomeNavigation/views /app/views
COPY --from=builder go/src/Gin_HomeNavigation/main /app
COPY --from=builder go/src/Gin_HomeNavigation/start.sh /

RUN mkdir /datas&&cp -ra /app/* /datas&&rm -rf /app

VOLUME /app
RUN chmod +x /start.sh

ENTRYPOINT  ["/start.sh"]