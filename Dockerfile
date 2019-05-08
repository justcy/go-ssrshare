FROM alpine:latest
MAINTAINER justcy <justxcy@gmail.com>

ADD ./dist/go-ssrshare /bin

RUN \
  chmod 0770 /bin/go-ssrshare

EXPOSE 8922

CMD ["go-ssrshare"]

