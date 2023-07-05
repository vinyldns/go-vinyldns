[![Build and Release](https://github.com/vinyldns/go-vinyldns/actions/workflows/go.yml/badge.svg)](https://github.com/vinyldns/go-vinyldns/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/vinyldns/go-vinyldns)](https://goreportcard.com/report/github.com/vinyldns/go-vinyldns) [![Godoc](https://godoc.org/github.com/vinyldns/go-vinyldns/vinyldns?status.svg)](https://godoc.org/github.com/vinyldns/go-vinyldns/vinyldns)
![GitHub](https://img.shields.io/github/license/vinyldns/go-vinyldns)

# Vinyldns

A Golang client for the [vinyldns](https://github.com/vinyldns/vinyldns) DNS as a service API.

## Usage

Install the client using the command: `go get github.com/vinyldns/go-vinyldns/vinyldns`.

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

## Local Development and Testing

### Install:

```
make install
```

Start the vinyldns local development server following the [quickstart instructions](https://github.com/vinyldns/vinyldns#quickstart). Download the credentials from the portal to use it in the client.

Create a file (for ex: `test.go`) in root directory (`go-vinyldns`). You can then check the local changes made to the `vinyldns` package from the file as follows:

```golang
package main

import "github.com/vinyldns/go-vinyldns/vinyldns"

func main() {
  client := vinyldns.NewClient(vinyldns.ClientConfiguration{
  "accessKey",
  "secretKey",
  "my-vinyldns-host.com",
  "my custom user agent", //optional
  })

  // For example, fetch zones...
  zs, err := client.Zones()
}
```
Run the file with the command: `go run test.go`

### Run tests w/ code coverage:

```
make test
```
