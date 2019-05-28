FROM  golang:1.12

MAINTAINER nickchou "nickchou"

WORKDIR $GOPATH/src/github.com/nickchou/go-hello

ADD . $GOPATH/src/github.com/nickchou/go-hello

RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]


