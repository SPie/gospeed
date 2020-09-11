FROM golang:latest

RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y sqlite3 \
    libsqlite3-dev

RUN mkdir /speedtest

ADD . /speedtest

WORKDIR /speedtest

RUN go build .

CMD ["./gospeed"]