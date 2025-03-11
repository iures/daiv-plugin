# DAIV Plugin System

This is a Go library for creating and managing plugins for the DAIV application ecosystem.

## Installation

To use this library in your main application or plugin, add it as a dependency:

```bash
go get github.com/iures/daiv-plugin
```

## Usage

### Creating a Plugin

To create a plugin that works with this system:

```go
package main

import (
	plugin "github.com/iures/daiv-plugin"
)

// MyPlugin implements the Plugin interface
type MyPlugin struct {
	config map[string]interface{}
}

// Name returns the unique identifier for this plugin
func (p *MyPlugin) Name() string {
	return "my-plugin"
}

// Manifest returns the configuration manifest
func (p *MyPlugin) Manifest() *plugin.PluginManifest {
	return &plugin.PluginManifest{
		ConfigKeys: []plugin.ConfigKey{
			{
				Key:         "api_key",
				Type:        plugin.ConfigTypePassword,
				Name:        "API Key",
				Description: "Your API key for the service",
				Required:    true,
				Secret:      true,
				EnvVar:      "MY_PLUGIN_API_KEY",
			},
		},
	}
}

// Initialize sets up the plugin with its configuration
func (p *MyPlugin) Initialize(settings map[string]interface{}) error {
	p.config = settings
	return nil
}

// Shutdown performs cleanup
func (p *MyPlugin) Shutdown() error {
	return nil
}

// Exporting the plugin
var Plugin MyPlugin

```

## Building and Distributing Plugins

To build your plugin as a shared library:

```bash
go build -buildmode=plugin -o my-plugin.so
```

Place this .so file in the plugins directory, and the main application will load it.

## Documentation

### Core Interfaces

- `Plugin`: The base interface all plugins must implement
- `StandupPlugin`: Interface for plugins that generate standup reports
