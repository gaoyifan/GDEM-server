FROM golang:1.5

MAINTAINER Yifan Gao "git@gaoyifan.com"

WORKDIR /srv

COPY main.go main.go

RUN go get github.com/gorilla/mux \
 && go get github.com/hashicorp/golang-lru \
 && go get gopkg.in/redis.v3 \
 && go build /srv/main.go

VOLUME /srv/map

EXPOSE 8000

CMD ./main
