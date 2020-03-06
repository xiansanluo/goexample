FROM alpine:3.2

MAINTAINER xiansan Luo<xiansan.luo@gmail.com>

LABEL name="strsvc" \
      description="example for service" \
      owner="xiansan.luo@gmail.com"

COPY ./server /server

WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/server"]
