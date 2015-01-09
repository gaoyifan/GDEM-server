# GDEM-server
The height data server for GDEM

# installation

1. install Golang Environment
2. install requirement via `go get ...`
3. run program `go run mapServer.go`

## environment installation specified in ubuntu 14.04

    # apt-get install golang
    # echo 'export GOPATH=$HOME/go' >> ~/.bashrc    #if you use bash
    # source ~/.bashrc
    # go get github.com/gorilla/mux
    # cd /path/to/project/src
    # editor GDEM-server.go   #modify prefix, port and etc as you like
    # go run GDEM-server.go
