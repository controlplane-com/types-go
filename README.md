# Control Plane Corp. Go Schema

This is a collection of Go types designed to streamline and standardize the data structures used across projects that use our [OpenAPI](https://console.cpln.io/openapi/core).

## Installation

```bash
go get github.com/controlplane-com/types-go
```

## Usage

Import the types from `github.com/controlplane-com/types-go/pkg/<package>` and use them in your Go projects:

```go
package main

import (
    "github.com/controlplane-com/types-go/pkg/base"
    "github.com/controlplane-com/types-go/pkg/gvc"
)

func example() {
    // Use base types
    kind := base.KindGvc

    // Use GVC types
    gvcData := &gvc.Gvc{
        // ... your data
    }
}
```
