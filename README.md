# Vital Go Library

[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/fern-vital/vital-go)
[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)

The Vital Go library provides convenient access to the Vital API from applications written in Go.

## Documentation

API reference documentation is available [here](https://docs.tryvital.io/home/welcome).

## Installation

Run the following command to use the Vital Go library in your Go module:

```sh
go get github.com/fern-vital/vital-go
```
## Requirements

This module requires Go version >= 1.19.

## Usage

```go
import (
  "context"
  "fmt"

  vital "github.com/fern-vital/vital-go"
  vitalclient "github.com/fern-vital/vital-go/client"
)

client := vitalclient.NewClient(
  vitalclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
)

response, err := client.LabTests.Get(context.TODO(), "OrderId")
if err != nil {
  return err
}
fmt.Printf("Received lab test %s\n", response)
```

## Client Options
A variety of client options are included to adapt the behavior of the library, which includes configuring authorization tokens to be sent on every request, or providing your own instrumented *http.Client. Both of these options are shown below:

```go
client := vitalclient.NewClient(
  vitalclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  vitalclient.ClientWithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Handling Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
response, err := client.LabTests.Get(
  context.TODO(),
  "OrderId",
)
if err != nil {
  if apiErr, ok := err.(*core.APIError); ok && apiErr.StatusCode == http.StatusBadRequest {
    // Do something with the bad request ...
  }
  return err
}
```

## Handling Errors

When you sign up to Vital you get access to two environments, Sandbox and Production.

| Environment URLs |                            |
| ---------------- | -------------------------- |
| `production`     | api.tryvital.io            |
| `production-eu`  | api.eu.tryvital.io         |
| `sandbox`        | api.sandbox.tryvital.io    |
| `sandbox-eu`     | api.sandbox.eu.tryvital.io |

By default, the SDK uses the `production` environment. See the snippet below
for an example on how ot change the environment.

```go
import (
  "context"

  vital "github.com/fern-vital/vital-go"
  vitalclient "github.com/fern-vital/vital-go/client"
)

client := vitalclient.NewClient(
  vitalclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  vitalclient.ClientWithBaseURL(vital.Environments.Sandbox)
)
```

## Beta status

This SDK is in beta, and there may be breaking changes between versions without a major version update.
Therefore, we recommend pinning the package version to a specific version in your `go.mod` file. This way,
you can install the same version each time without breaking changes unless you are intentionally looking
for the latest version.

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. Additions
made directly to this library would have to be moved over to our generation code, otherwise they would be
overwritten upon the next generated release. Feel free to open a PR as a proof of concept, but know that we
will not be able to merge it as-is. We suggest opening an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
