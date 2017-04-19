# Redirects

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/redirects/status.svg)](http://github.dronehippie.de/webhippie/redirects)
[![Go Doc](https://godoc.org/github.com/webhippie/redirects?status.svg)](http://godoc.org/github.com/webhippie/redirects)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/redirects)](http://goreportcard.com/report/github.com/webhippie/redirects)
[![](https://images.microbadger.com/badges/image/webhippie/redirects.svg)](http://microbadger.com/images/webhippie/redirects "Get your own image badge on microbadger.com")
[![Join the chat at https://gitter.im/webhippie/general](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/webhippie/general)
[![Stories in Ready](https://badge.waffle.io/webhippie/redirects.svg?label=ready&title=Ready)](http://waffle.io/webhippie/redirects)

**This project is under heavy development, it's not in a working state yet!**

Redirects is a pretty simple pattern-based redirect server. It supports different kinds of backends to store the patterns like JSON, YAML, TOML, Etcd, Consul and Zookeeper. We are using it mostly as a default route for our reverse proxy like [Træfɪk](https://traefik.io/).


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.webhippie.de/misc/redirects). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie). If you are missing an architecture just write us on our nice [Gitter](https://gitter.im/webhippie/general) chat. If you find a security issue please contact thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring of the dependencies and we are not exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go version `>= 1.6`. It is also possible to just simply execute the `go get github.com/webhippie/redirects` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/redirects
cd $GOPATH/src/github.com/webhippie/redirects
make clean build

./redirects -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2017 Thomas Boerger <thomas@webhippie.de>
```
