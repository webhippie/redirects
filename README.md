# Redirects

[![Current Tag](https://img.shields.io/github/v/tag/webhippie/redirects?sort=semver)](https://github.com/webhippie/redirects) [![Build Status](https://github.com/webhippie/redirects/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/redirects/actions) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Docker Size](https://img.shields.io/docker/image-size/webhippie/redirects/latest)](https://hub.docker.com/r/webhippie/redirects) [![Docker Pulls](https://img.shields.io/docker/pulls/webhippie/redirects)](https://hub.docker.com/r/webhippie/redirects) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/redirects.svg)](https://pkg.go.dev/github.com/webhippie/redirects) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/redirects)](https://goreportcard.com/report/github.com/webhippie/redirects) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c33cdacbc6ed40beaf1f2a0a6d72718b)](https://www.codacy.com/gh/webhippie/redirects/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/redirects&amp;utm_campaign=Badge_Grade)

Redirects is a pretty simple pattern-based redirect server. It supports
different kinds of backends to store the patterns like JSON, YAML, TOML, Etcd,
Consul and Zookeeper. We are using it mostly as a default route for our reverse
proxy like [Træfɪk](https://traefik.io/).

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.19, at least that's the version we are using.

```console
git clone https://github.com/webhippie/redirects.git
cd redirects

make generate build

./bin/redirects -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2017 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/webhippie/redirects/releases
[dockerhub]: https://hub.docker.com/r/webhippie/redirects/tags/
[quay]: https://quay.io/repository/webhippie/redirects?tab=tags
[docs]: https://webhippie.github.io/redirects/#getting-started
[golang]: http://golang.org/doc/install.html
