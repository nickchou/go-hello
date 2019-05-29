FROM scratch

MAINTAINER nickchou "nickchou"

ADD go-hello /go-hello

EXPOSE 8081

CMD ["./go-hello"]

