---
title: "Getting Started"
date: 2022-05-04T00:00:00+00:00
anchor: "getting-started"
weight: 20
---

## Installation

So far we are offering only a few different variants for the installation. You
can choose between [Docker][docker] or pre-built binaries which are stored on
our download mirror and GitHub releases. Maybe we will also provide system
packages for the major distributions later if we see the need for it.

### Docker

Generally we are offering the images through
[quay.io/webhippie/redirects][quay] and [webhippie/redirects][dockerhub], so
feel free to choose one of the providers. Maybe we will come up with Kustomize
manifests or some Helm chart.

### Binaries

Simply download a binary matching your operating system and your architecture
from our [downloads][downloads] or the GitHub releases and place it within your
path like `/usr/local/bin` if you are using macOS or Linux.

## Configuration

We provide overall three different variants of configuration. The variant based
on environment variables and commandline flags are split up into global values
and command-specific values.

### Envrionment variables

If you prefer to configure the service with environment variables you can see
the available variables below.

#### Global

REDIRECTS_CONFIG_FILE
: Path to optional config file

REDIRECTS_LOG_LEVEL
: Set logging level, defaults to `info`

REDIRECTS_LOG_COLOR
: Enable colored logging, defaults to `true`

REDIRECTS_LOG_PRETTY
: Enable pretty logging, defaults to `true`

REDIRECTS_YAML_ENABLED
: Enable YAML storage, defaults to `false`

REDIRECTS_YAML_FILE
: Define YAML storage file, defaults to `file://storage/redirects.yaml`

REDIRECTS_JSON_ENABLED
: Enable JSON storage, defaults to `false`

REDIRECTS_JSON_FILE
: Define JSON storage file, defaults to `file://storage/redirects.json`

REDIRECTS_TOML_ENABLED
: Enable TOML storage, defaults to `false`

REDIRECTS_TOML_FILE
: Define TOML storage file, defaults to `file://storage/redirects.toml`

REDIRECTS_ZK_ENABLED
: Enable Zookeeper storage, defaults to `false`

REDIRECTS_ZK_ENDPOINT
: Used Zookeeper endpoints, defaults to `127.0.0.1:2181`, comma-separated for multiple endpoints

REDIRECTS_ZK_TIMEOUT
: Connection timeout for Zookeeper storage, defaults to `10s`

REDIRECTS_ZK_PREFIX
: Define Zookeeper storage prefix, defaults to `/redirects`

REDIRECTS_ETCD_ENABLED
: Enable Etcd storage, defaults to `false`

REDIRECTS_ETCD_ENDPOINT
: Used Etcd endpoints, defaults to `127.0.0.1:2379`, comma-separated for multiple endpoints

REDIRECTS_ETCD_TIMEOUT
: Connection timeout for Etcd storage, defaults to `10s`

REDIRECTS_ETCD_PREFIX
: Define Etcd storage prefix, defaults to `/redirects`

REDIRECTS_ETCD_USERNAME
: Username to access Etcd

REDIRECTS_ETCD_PASSWORD
: Password to access Etcd

REDIRECTS_ETCD_CA
: Path to CA cert to access Etcd

REDIRECTS_ETCD_CERT
: Path to SSL certificate to access Etcd

REDIRECTS_ETCD_KEY
: Path to SSL key to access Etcd

REDIRECTS_ETCD_SKIP_VERIFY
: Skip SSL verification for Etcd, defaults to `false`

REDIRECTS_CONSUL_ENABLED
: Enable Consul storage, defaults to `false`

REDIRECTS_CONSUL_ENDPOINT
: Used Consul endpoints, defaults to `127.0.0.1:8500`, somma-separated for multiple endpoints

REDIRECTS_CONSUL_TIMEOUT
: Connection timeout for Consul storage, defaults to `10s`

REDIRECTS_CONSUL_PREFIX
: Define Consul storage prefix, defaults to `/redirects`

REDIRECTS_CONSUL_CA
: Path to CA cert to access Consul

REDIRECTS_CONSUL_CERT
: Path to SSL certificate to access Consul

REDIRECTS_CONSUL_KEY
: Path to SSL key to access Consul

REDIRECTS_CONSUL_SKIP_VERIFY
: Skip SSL verification for Consul, defaults to `false`

#### Server

REDIRECTS_METRICS_ADDR
: Address to bind the metrics, defaults to `0.0.0.0:8081`

REDIRECTS_METRICS_TOKEN
: Token to make metrics secure

REDIRECTS_SERVER_ADDR
: Address to bind the server, defaults to `0.0.0.0:8080`

REDIRECTS_SERVER_PPROF
: Enable pprof debugging, defaults to `false`

REDIRECTS_SERVER_ROOT
: Root path of the server, defaults to `/`

REDIRECTS_SERVER_HOST
: External access to server, defaults to `http://localhost:8080`

REDIRECTS_SERVER_CERT
: Path to cert for SSL encryption

REDIRECTS_SERVER_KEY
: Path to key for SSL encryption

REDIRECTS_SERVER_STRICT_CURVES
: Use strict SSL curves, defaults to `false`

REDIRECTS_SERVER_STRICT_CIPHERS
: Use strict SSL ciphers, defaults to `false`

REDIRECTS_SERVER_TEMPLATES
: Folder for custom templates

#### Health

REDIRECTS_METRICS_ADDR
: Address to bind the metrics, defaults to `0.0.0.0:8081`

### Commandline flags

If you prefer to configure the service with commandline flags you can see the
available variables below.

#### Global

--config-file
: Path to optional config file

--log-level
: Set logging level, defaults to `info`

--log-color
: Enable colored logging, defaults to `true`

--log-pretty
: Enable pretty logging, defaults to `true`

--yaml-enabled
: Enable YAML storage, defaults to `false`

--yaml-file
: Define YAML storage file, defaults to `file://storage/redirects.yaml`

--json-enabled
: Enable JSON storage, defaults to `false`

--json-file
: Define JSON storage file, defaults to `file://storage/redirects.json`

--toml-enabled
: Enable TOML storage, defaults to `false`

--toml-file
: Define TOML storage file, defaults to `file://storage/redirects.toml`

--zk-enabled
: Enable Zookeeper storage, defaults to `false`

--zk-endpoint
: Used Zookeeper endpoints, defaults to `127.0.0.1:2181`, repeat for multiple endpoints

--zk-timeout
: Connection timeout for Zookeeper storage, defaults to `10s`

--zk-prefix
: Define Zookeeper storage prefix, defaults to `/redirects`

--etcd-enabled
: Enable Etcd storage, defaults to `false`

--etcd-endpoint
: Used Etcd endpoints, defaults to `127.0.0.1:2379`, repeat for multiple endpoints

--etcd-timeout
: Connection timeout for Etcd storage, defaults to `10s`

--etcd-prefix
: Define Etcd storage prefix, defaults to `/redirects`

--etcd-username
: Username to access Etcd

--etcd-password
: Password to access Etcd

--etcd-ca
: Path to CA cert to access Etcd

--etcd-cert
: Path to SSL certificate to access Etcd

--etcd-key
: Path to SSL key to access Etcd

--etcd-skip-verify
: Skip SSL verification for Etcd, defaults to `false`

--consul-enabled
: Enable Consul storage, defaults to `false`

--consul-endpoint
: Used Consul endpoints, defaults to `127.0.0.1:8500`, repeat for multiple endpoints

--consul-timeout
: Connection timeout for Consul storage, defaults to `10s`

--consul-prefix
: Define Consul storage prefix, defaults to `/redirects`

--consul-ca
: Path to CA cert to access Consul

--consul-cert
: Path to SSL certificate to access Consul

--consul-key
: Path to SSL key to access Consul

--consul-skip-verify
: Skip SSL verification for Consul, defaults to `false`

#### Server

--metrics-addr
: Address to bind the metrics, defaults to `0.0.0.0:8081`

--metrics-token
: Token to make metrics secure

--server-addr
: Address to bind the server, defaults to `0.0.0.0:8080`

--server-pprof
: Enable pprof debugging, defaults to `false`

--server-root
: Root path of the server, defaults to `/`

--server-host
: External access to server, defaults to `http://localhost:8080`

--server-cert
: Path to cert for SSL encryption

--server-key
: Path to key for SSL encryption

--strict-curves
: Use strict SSL curves, defaults to `false`

--strict-ciphers
: Use strict SSL ciphers, defaults to `false`

--templates-path
: Folder for custom templates

#### Health

--metrics-addr
: Address to bind the metrics, defaults to `0.0.0.0:8081`

### Configuration file

So far we support multiple file formats like `json` or `yaml`, if you want to
get a full example configuration just take a look at [our repository][repo],
there you can always see the latest configuration format. These example configs
include all available options and the default values. The configuration file
will be automatically loaded if it's placed at
`/etc/redirects/config.yml`, `${HOME}/.redirects/config.yml` or
`$(pwd)/redirects/config.yml`.

## Usage

The program provides a few sub-commands on execution. The available config
methods have already been mentioned above. Generally you can always see a
formated help output if you execute the binary similar to something like
 `redirects --help`.

[docker]: https://www.docker.com/
[quay]: https://quay.io/repository/webhippie/redirects
[dockerhub]: https://hub.docker.com/r/webhippie/redirects
[downloads]: https://dl.webhippie.de/#redirects/
[repo]: https://github.com/webhippie/redirects/tree/master/config
