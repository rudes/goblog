FROM golang

ADD . /go/src/github.com/rudes/otherletters.net
RUN cd /go/src/github.com/rudes/otherletters.net; go get
RUN go install github.com/rudes/otherletters.net

EXPOSE 8080

ENTRYPOINT /go/bin/otherletters.net
