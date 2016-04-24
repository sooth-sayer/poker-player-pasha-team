FROM golang:1.4.2

RUN apt-get update -qq
RUN apt-get install -y build-essential pkg-config libaio1 libaio-dev alien

WORKDIR /app

ENV GOBIN $GOPATH/bin

ADD . /app

RUN go get github.com/tools/godep
RUN make deps

EXPOSE 8080

CMD make
