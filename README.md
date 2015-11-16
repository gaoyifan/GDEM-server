# GDEM-server
The height data server for GDEM

# installation

1. install Golang Environment
2. install requirement via `go get ...`
3. run program `go run main.go`

## environment installation specified in ubuntu 14.04

    # apt-get install golang
    # echo 'export GOPATH=$HOME/go' >> ~/.bashrc    #if you use bash
    # source ~/.bashrc
    # go get github.com/gorilla/mux
    # cd /path/to/project
    # editor main.go   #modify prefix, port and etc as you like
    # go run main.go

# Download Map Data

You can downlaod the GDEM v2 Map Data(global data) which is converted to png format:

    http://pan.baidu.com/s/1pJBT1qz

Or you can just downlaod the data you need.
