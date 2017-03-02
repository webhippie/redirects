FROM webhippie/alpine:latest
MAINTAINER Thomas Boerger <thomas@webhippie.de>

EXPOSE 8080
VOLUME ["/var/lib/redirects"]

RUN apk update && \
  apk add \
    ca-certificates \
    bash && \
  rm -rf \
    /var/cache/apk/* && \
  addgroup \
    -g 1000 \
    redirects && \
  adduser -D \
    -h /var/lib/redirects \
    -s /bin/bash \
    -G redirects \
    -u 1000 \
    redirects

COPY redirects /usr/bin/

USER redirects
ENTRYPOINT ["/usr/bin/redirects"]
CMD ["server"]
