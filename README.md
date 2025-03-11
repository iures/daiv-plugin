# DAIV Plugin System

This is a Go library for creating and managing plugins for the DAIV application ecosystem.

## Installation

To use this library in your main application or plugin, add it as a dependency:

```bash
go get github.com/iures/daiv-plugin
```

## Usage

### In the Main Application

To use the plugin system in your main application:

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	plugin "github.com/iures/daiv-plugin"
)

func main() {
	// Set up plugin directory
	homeDir, _ := os.UserHomeDir()
	pluginsDir := filepath.Join(homeDir, ".daiv", "plugins")
	
	// Create plugin manager
	manager, err := plugin.NewPluginManager(pluginsDir)
	if err != nil {
		fmt.Printf("Error creating plugin manager: %v\n", err)
		return
	}
	
	// Get plugin registry
	registry := plugin.GetRegistry()
	
	// Load all available plugins
	if err := registry.LoadExternalPlugins(); err != nil {
		fmt.Printf("Error loading plugins: %v\n", err)
	}
	
	// Use all standup plugins
	standupPlugins := registry.GetStandupPlugins()
	for _, p := range standupPlugins {
		// Initialize the plugin
		if err := plugin.Initialize(p); err != nil {
			fmt.Printf("Error initializing plugin %s: %v\n", p.Name(), err)
			continue
		}
		
		// Use the plugin...
	}
}
```

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

// This function is required for Go to load your plugin
func main() {}
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

### Key Structs

- `PluginManager`: Handles downloading, installing, and loading plugins
- `Registry`: Manages the registration and access of plugins
- `PluginManifest`: Defines the configuration requirements for a plugin

## License

[Add your license information here] 
