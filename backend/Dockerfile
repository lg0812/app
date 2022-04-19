FROM golang:latest as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src
COPY . .

RUN go build -o backend .

FROM ubuntu:latest

WORKDIR /backend

COPY --from=builder /go/src/backend .
ENV TZ=Asia/ShanghaiENV TZ=Asia/Shanghai
EXPOSE 8080

ENTRYPOINT ./backend