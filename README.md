[![Build and Release](https://github.com/vinyldns/go-vinyldns/actions/workflows/go.yml/badge.svg)](https://github.com/vinyldns/go-vinyldns/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/vinyldns/go-vinyldns)](https://goreportcard.com/report/github.com/vinyldns/go-vinyldns) [![Godoc](https://godoc.org/github.com/vinyldns/go-vinyldns/vinyldns?status.svg)](https://godoc.org/github.com/vinyldns/go-vinyldns/vinyldns)
![GitHub](https://img.shields.io/github/license/vinyldns/go-vinyldns)

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

Go 1.27.3 or later is required. CI reads the required version from `go.mod`.

Run tests w/ code coverage:

```
make test
```

Install:

```
make install
```

## Release notes

### 0.9.19

- Require Go 1.27.3 and use the same toolchain in CI.
- Update the AWS SDK and credential provider, including the Smithy dependency.
- Preserve the existing client API and request-signing behavior.

Consumers must update this module and rebuild their applications with a patched
Go toolchain. Updating this library does not change existing CLI binaries.

The client uses Go's default TLS certificate verification and the host's trust
configuration; it does not bundle a certificate store or implement certificate
revocation checks. A toolchain update alone does not establish remediation for a
compromised certificate authority. Apply any required trust-store updates under
your organization's certificate guidance. Certificate-advisory validation and
replacement CLI binaries remain separate release requirements.
