FROM golang:1.12.0

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1

ENV PATH $PATH:/www/app/vendor/bin

RUN \
  apt update -y && \
  apt install -y \
  autoconf-dev \
  automake \
  libtool \
  curl \
  make \
  g++ \
  unzip && \
  ln -f -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN git clone https://github.com/google/protobuf.git
WORKDIR /www/app
