package filter_plugin

import (
	"fmt"
)

// TODO: 1. add mutex  2. filter chain?
type PluginRegistry struct {
	plugins map[string]FilterPlugin
}

func (r *PluginRegistry) Register(plugin FilterPlugin) error {
	if r.plugins == nil {
		r.plugins = make(map[string]FilterPlugin)
	}
	if plugin == nil {
		return fmt.Errorf("cannot register nil plugin")
	}
	if _, exists := r.plugins[plugin.Name()]; exists {
		return fmt.Errorf("plugin with name %s already exists", plugin.Name())
	}
	r.plugins[plugin.Name()] = plugin
	return nil
}

func (r *PluginRegistry) Deregister(name string) error {
	if _, exists := r.plugins[name]; !exists {
		return fmt.Errorf("plugin with name %s does not exist", name)
	}
	delete(r.plugins, name)
	return nil
}

func (r *PluginRegistry) GetPlugin(name string) (FilterPlugin, error) {
	plugin, exists := r.plugins[name]
	if !exists {
		return nil, fmt.Errorf("plugin with name %s does not exist", name)
	}
	return plugin, nil
}

type ListPluginResult struct {
	Name        string
	Description string
	Version     string
}

func (r *PluginRegistry) ListPlugins() []ListPluginResult {
	results := make([]ListPluginResult, 0, len(r.plugins))
	for _, plugin := range r.plugins {
		results = append(results, ListPluginResult{
			Name:        plugin.Name(),
			Description: plugin.Description(),
			Version:     plugin.Version(),
		})
	}
	return results
}

var (
	globalRegistry = &PluginRegistry{}
)

func GetRegistry() *PluginRegistry {
	return globalRegistry
}
