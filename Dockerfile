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

# ARG VERSION
# ARG BUILD_DATE
# ARG VCS_REF

# LABEL org.label-schema.version=$VERSION
# LABEL org.label-schema.build-date=$BUILD_DATE
# LABEL org.label-schema.vcs-ref=$VCS_REF
LABEL org.label-schema.vcs-url="https://github.com/tboerger/redirects.git"
LABEL org.label-schema.name="Redirects"
LABEL org.label-schema.vendor="Thomas Boerger"
LABEL org.label-schema.schema-version="1.0"
