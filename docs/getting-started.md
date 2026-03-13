# Getting Started

`go-plugin` focuses on deterministic plugin registration and lifecycle management for applications that want an explicit, testable plugin model.

## Main components

- `plugin.Registry` stores plugin instances in insertion order.
- `plugin.FactoryRegistry` stores named plugin factories.
- `plugin.DiscoverFromValues` discovers values implementing the `Plugin` interface.
- `plugin.RegisterRouteProviders` bridges `Routes()` providers into `go-module-router`.

## Extension points

The package also includes optional helpers for `.so` loading and process-based sandboxes. Those helpers stay small and can be adopted incrementally when your plugin model needs them.
