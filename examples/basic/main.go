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
