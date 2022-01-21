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

RUN go build -o Gin_HomeNavigation .

FROM scratch

COPY --from=builder $GOPATH/src/Gin_HomeNavigation/Gin_HomeNavigation /

ENTRYPOINT  ["./Gin_HomeNavigation"]