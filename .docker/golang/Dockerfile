FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/gl-proxy

WORKDIR /go/src/github.com/Sharykhin/gl-proxy

RUN go get .

RUN go install github.com/Sharykhin/gl-proxy

ENTRYPOINT /go/bin/gl-proxy

#CMD tail -f /dev/null

EXPOSE 8888