FROM --platform=$BUILDPLATFORM golang:1.25.2-alpine3.21@sha256:acbf7658ccbcb6025a0c2cf446ec12d869ee1449ba9aa854453644e58c7d5913 AS builder

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

FROM alpine:3.22@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1

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
