FROM golang:1.8

MAINTAINER Ryan Bak

RUN go get -v -d github.com/rbak1/test-site/...

WORKDIR /go/src/github.com/rbak1/test-site

RUN go run build.go build

EXPOSE 4000

CMD /go/src/github.com/rbak1/test-site/bin/test-site
