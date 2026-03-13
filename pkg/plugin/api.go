package plugin

// Plugin defines the lifecycle and identity for plugins.
type Plugin interface {
	Name() string
	Start() error
	Stop() error
}

// Factory creates plugin instances on demand.
type Factory func() Plugin
