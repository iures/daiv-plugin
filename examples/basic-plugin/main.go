package main

import (
	"fmt"
	"time"

	plugin "github.com/iures/daiv-plugin"
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
func (p *BasicPlugin) Manifest() *plugin.PluginManifest {
	return &plugin.PluginManifest{
		ConfigKeys: []plugin.ConfigKey{
			{
				Key:         "username",
				Type:        plugin.ConfigTypeString,
				Name:        "Username",
				Description: "Your username",
				Required:    true,
			},
			{
				Key:         "is_enabled",
				Type:        plugin.ConfigTypeBoolean,
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
func (p *BasicPlugin) GetStandupContext(timeRange plugin.TimeRange) (plugin.StandupContext, error) {
	username, _ := p.config["username"].(string)
	
	return plugin.StandupContext{
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
