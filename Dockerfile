FROM golang:1.5

MAINTAINER Yifan Gao "git@gaoyifan.com"

WORKDIR /srv

RUN git clone https://github.com/gaoyifan/GDEM-server.git repo \
    && go get github.com/gorilla/mux \
    && go build /srv/repo/src/GDEM-server/GDEM-server.go

VOLUME /srv/map

EXPOSE 8000

CMD ./GDEM-server
