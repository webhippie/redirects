FROM --platform=$BUILDPLATFORM golang:1.26.2-alpine@sha256:27f829349da645e287cb195a9921c106fc224eeebbdc33aeb0f4fca2382befa6 AS builder

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

FROM alpine:3.23@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

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
