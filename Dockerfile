FROM golang:1.17
MAINTAINER HomeNavigation "nuanxinqing@gmail.com"
WORKDIR $GOPATH/src/Gin_HomeNavigation
COPY . .
ADD . ./
#设置环境变量，开启go module和设置下载代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

#增加缺失的包，移除没用的包
RUN go mod tidy

RUN go build -o Gin_HomeNavigation .

CMD ["./Gin_HomeNavigation"]