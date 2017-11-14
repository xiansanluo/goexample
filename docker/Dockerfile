FROM alpine:3.2

MAINTAINER xiansan Luo<xiansan.luo@liulishuo.com>

LABEL name="strsvc" \
      description="example for service" \
      owner="xiansan.luo@liulishuo.com"

COPY ./server /server

WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/server"]
