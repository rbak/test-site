FROM golang

MAINTAINER Ryan Bak

WORKDIR /go

RUN go install -d github.com/rbak1/test-site

RUN go run /src/github.com/rbak1/test-site/build.go build

EXPOSE 4000

CMD /src/github.com/rbak1/test-site/bin/test-site
