FROM --platform=$BUILDPLATFORM golang:1.25.5-alpine AS builder

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

FROM alpine:3.23@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659

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
