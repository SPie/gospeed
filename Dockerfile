FROM golang:1.14

RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y sqlite3 \
    libsqlite3-dev

RUN mkdir /speedtest

ADD . /speedtest

WORKDIR /speedtest/cmd/speedtest

RUN go build .