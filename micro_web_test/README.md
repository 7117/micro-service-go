# Micro_web_test Service

This is the Micro_web_test service

Generated with

```
micro new micro_web_test --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.micro_web_test
- Type: web
- Alias: micro_web_test

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./micro_web_test-web
```

Build a docker image
```
make docker
```