FROM --platform=$BUILDPLATFORM golang:1.27.3-alpine@sha256:d200d2dd76b03a99b52651fd1325aff7dcd35422c35c01bac56571b0507f3deb AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/redirects
COPY . /go/src/redirects/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 redirects && \
    adduser -D -u 1337 -h /var/lib/redirects -G redirects redirects

EXPOSE 8080 8081
VOLUME ["/var/lib/redirects"]
ENTRYPOINT ["/usr/bin/redirects"]
CMD ["server"]
HEALTHCHECK CMD ["/usr/bin/redirects", "health"]

COPY --from=builder /go/src/redirects/bin/redirects /usr/bin/redirects
WORKDIR /var/lib/redirects
USER redirects
