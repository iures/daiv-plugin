package main

import (
	"fmt"
	"time"

	plug "github.com/iures/daivplug"
)

// BasicPlugin implements a simple example plugin
type BasicPlugin struct {
	config map[string]interface{}
}

// Name returns the unique identifier for this plugin
func (p *BasicPlugin) Name() string {
	return "basic-example-plugin"
}

// Manifest returns the configuration manifest
func (p *BasicPlugin) Manifest() *plug.PluginManifest {
	return &plug.PluginManifest{
		ConfigKeys: []plug.ConfigKey{
			{
				Key:         "username",
				Type:        plug.ConfigTypeString,
				Name:        "Username",
				Description: "Your username",
				Required:    true,
			},
			{
				Key:         "is_enabled",
				Type:        plug.ConfigTypeBoolean,
				Name:        "Enable Features",
				Description: "Enable additional features",
				Required:    false,
				Value:       false,
			},
		},
	}
}

// Initialize sets up the plugin with its configuration
func (p *BasicPlugin) Initialize(settings map[string]interface{}) error {
	p.config = settings
	return nil
}

// Shutdown performs cleanup
func (p *BasicPlugin) Shutdown() error {
	return nil
}

// GetStandupContext implements the StandupPlugin interface
func (p *BasicPlugin) GetStandupContext(timeRange plug.TimeRange) (plug.StandupContext, error) {
	username, _ := p.config["username"].(string)
	
	return plug.StandupContext{
		PluginName: p.Name(),
		Content:    fmt.Sprintf("Hello, %s! This is a basic plugin example.\nTime range: %s to %s", 
			username, 
			timeRange.Start.Format(time.RFC822), 
			timeRange.End.Format(time.RFC822)),
	}, nil
}

// This function is required for Go to load your plugin
func main() {}

// Export the plugin so it can be loaded
var Plugin BasicPlugin 
