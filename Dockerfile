FROM golang

ADD . /go/src/github.com/rudes/otherletter/
RUN go install github.com/rudes/otherletter
ENTRYPOINT /go/bin/otherletter

EXPOSE 8080
