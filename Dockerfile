FROM golang:1.5

MAINTAINER Yifan Gao "git@gaoyifan.com"

WORKDIR /srv

RUN git clone https://github.com/gaoyifan/GDEM-server.git repo \
    && go get github.com/gorilla/mux \
    && go get github.com/hashicorp/golang-lru \
    && go get gopkg.in/redis.v3 \
    && go build /srv/repo/main.go

VOLUME /srv/map

EXPOSE 8000

CMD ./main
