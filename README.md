# Go Plugin

> [!CAUTION]
> go-plugin is now part of the [go-foundation](https://github.com/mirkobrombin/go-foundation) framework. The v1.0.0 release mirrors go-plugin v0.1.0, but future versions may introduce breaking changes. Please migrate your project.

A structured **plugin registry** for Go with deterministic lifecycle management and lightweight discovery helpers.

## Features

- **Deterministic Registry:** Preserve insertion order for predictable startup and shutdown.
- **Factory Support:** Register deferred constructors for dynamic plugin creation.
- **Reflection Discovery:** Discover plugin implementations from runtime values.
- **Optional Router Integration:** Bridge route-providing plugins into `go-module-router`.

## Installation

```bash
go get github.com/mirkobrombin/go-plugin
```

## Quick Start

```go
package main

import "github.com/mirkobrombin/go-plugin/pkg/plugin"

type samplePlugin struct{}

func (samplePlugin) Name() string { return "sample" }
func (samplePlugin) Start() error { return nil }
func (samplePlugin) Stop() error  { return nil }

func main() {
    registry := plugin.NewRegistry()
    if err := registry.Register(samplePlugin{}); err != nil {
        panic(err)
    }

    _ = registry.StartAll()
    _ = registry.StopAll()
}
```

## Documentation

- [Getting Started](docs/getting-started.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
