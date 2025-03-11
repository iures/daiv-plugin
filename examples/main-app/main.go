package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	plugin "github.com/iures/daiv-plugin"
)

func main() {
	// Set up plugin directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}
	
	pluginsDir := filepath.Join(homeDir, ".daiv", "plugins")
	
	// Create plugin manager
	manager, err := plugin.NewPluginManager(pluginsDir)
	if err != nil {
		fmt.Printf("Error creating plugin manager: %v\n", err)
		return
	}
	
	// Use the manager variable to avoid unused variable error
	fmt.Printf("Plugin manager initialized at: %p\n", manager)
	
	// Example of installing a plugin (commented out for safety)
	/*
	err = manager.InstallFromGitHub("username/repo-name", "main")
	if err != nil {
		fmt.Printf("Error installing plugin: %v\n", err)
	}
	*/
	
	// Get plugin registry
	registry := plugin.GetRegistry()
	
	// Load all available plugins
	err = registry.LoadExternalPlugins()
	if err != nil {
		fmt.Printf("Error loading plugins: %v\n", err)
	}
	
	// Use all standup plugins
	standupPlugins := registry.GetStandupPlugins()
	fmt.Printf("Found %d standup plugins\n", len(standupPlugins))
	
	for _, p := range standupPlugins {
		// Initialize the plugin
		if err := plugin.Initialize(p); err != nil {
			fmt.Printf("Error initializing plugin %s: %v\n", p.Name(), err)
			continue
		}
		
		// Set time range for last 24 hours
		now := time.Now()
		yesterday := now.Add(-24 * time.Hour)
		timeRange := plugin.TimeRange{
			Start: yesterday,
			End:   now,
		}
		
		// Get standup context from the plugin
		context, err := p.GetStandupContext(timeRange)
		if err != nil {
			fmt.Printf("Error getting standup context from %s: %v\n", p.Name(), err)
			continue
		}
		
		// Display the plugin's output
		fmt.Println(context.String())
	}
	
	// Shutdown all plugins
	if err := registry.ShutdownAll(); err != nil {
		fmt.Printf("Error shutting down plugins: %v\n", err)
	}
} 
