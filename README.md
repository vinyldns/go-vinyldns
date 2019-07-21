[![Build Status](https://travis-ci.org/vinyldns/go-vinyldns.svg?branch=master)](https://travis-ci.org/vinyldns/go-vinyldns) [![Go Report Card](https://goreportcard.com/badge/github.com/vinyldns/go-vinyldns)](https://goreportcard.com/report/github.com/vinyldns/go-vinyldns)

# vinyldns

A Golang client for the [vinyldns](https://github.com/vinyldns/vinyldns) DNS as a service API.

## Usage

Basic usage requires instantiating a client and using the `vinyldns/api.go` methods to interact with `vinyldns`:

For example:

```golang
import "github.com/vinyldns/go-vinyldns/vinyldns"

client := vinyldns.NewClient(vinyldns.ClientConfiguration{
  "accessKey",
  "secretKey",
  "my-vinyldns-host.com",
  "my custom user agent",
})

// For example, fetch zones...
// returns vinyldns.Error, []vinyldns.Zone
zs, err := client.Zones()
```

Alternatively, `NewClientFromEnv` instantiates a client from the following environment variables:

```
VINYLDNS_ACCESS_KEY=
VINYLDNS_SECRET_KEY=
VINYLDNS_HOST=

# Optional; defaults to `go-vinyldns/<version>`
VINYLDNS_USER_AGENT=
```

```golang
import "github.com/vinyldns/go-vinyldns/vinyldns"

client := vinyldns.NewClientFromEnv()
```

See `vinyldns/${resource}_resources.go` files for the various `vinyldns` resource structs.

See `vinyldns/${resource}.go` files for the various `vinyldns` API methods.

## Development

Install dependencies:

```
make deps
```

Run tests w/ code coverage:

```
make test
```

Install:

```
make install
```
