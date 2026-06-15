package engine

import (
	"image"
	"os"
	"plugin"
	"strings"

	"github.com/GrapefruitCat030/gfc_plugo_filter/pkg/filter_plugin"
)

type Engine struct{}

func (e *Engine) LoadPlugin(pluso_path string) error {
	// get the .so file
	pluso, err := plugin.Open(pluso_path)
	if err != nil {
		return err
	}
	factorySymbol, err := pluso.Lookup("CreatePlugin")
	if err != nil {
		return err
	}
	pluFactory := factorySymbol.(filter_plugin.FilterPluginFactory)
	pluIns := pluFactory()
	// register the plugin
	registry := filter_plugin.GetRegistry()
	return registry.Register(pluIns)
}

func (e *Engine) LoadAllPlugins(pluso_dir string) error {
	entries, err := os.ReadDir(pluso_dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".so") {
			continue
		}
		pluso_path := entry.Name()
		if err := e.LoadPlugin(pluso_path); err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) UnloadPlugin(filter_name string) error {
	registry := filter_plugin.GetRegistry()
	return registry.Deregister(filter_name)
}

// just apply one filter now
func (e *Engine) ApplyFilter(input image.Image, filter_name string) (image.Image, error) {
	registry := filter_plugin.GetRegistry()
	plugin, err := registry.GetPlugin(filter_name)
	if err != nil {
		return nil, err
	}
	return plugin.Apply(input)
}

var (
	globalEngine = &Engine{}
)

func GetEngine() *Engine {
	return globalEngine
}
