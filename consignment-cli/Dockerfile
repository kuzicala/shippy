FROM golang:latest as builder

WORKDIR /go/src/shippy/consignment-cli
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN export http_proxy=http://192.168.20.128:1087 && export https_proxy=http://192.168.20.128:1087 && dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# 若运行环境是 Linux 则需把 alpine 换成 debian
# 使用最新版 alpine 作为基础镜像
FROM alpine:latest

RUN apk --no-cache add ca-certificates
# 在容器的根目录下创建 app 目录
RUN mkdir /app
# 将工作目录切换到 /app 下
WORKDIR /app
# 将微服务的服务端运行文件拷贝到 /app 下
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/shippy/consignment-cli/consignment-cli .
# 运行服务端
CMD ["./consignment-cli"]