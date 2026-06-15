package filter_plugin

import "image"

// FilterPlugin is the interface that all filter plugins must implement.
type FilterPlugin interface {
	Name() string
	Description() string
	Version() string
	Apply(input image.Image) (image.Image, error)
}

// TODO: use a struct or interface?
//
//	type FilterPluginFactory struct {
//		Name        string
//		Description string
//		Version     string
//		Create      func() FilterPlugin
//	}

// FilterPluginFactory is a function type that creates a new instance of a FilterPlugin. "CreatePlugin" here.
type FilterPluginFactory func() FilterPlugin
