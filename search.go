package oi

import (
	"fmt"
	"os/user"
	"path/filepath"

	"log"
)

// Get the current user to find their home directory.
var usr, _ = user.Current()

// TODO(orlade): Replace with environment variable configuration.
var oipath = usr.HomeDir + "/.oi/modules"

// PluginSearcher finds availble plugins.
type PluginSearcher interface {
	// Searches for plugins in some location and returns a map of names to paths.
	Search() map[string]string
}

// HomePluginSearcher searches for plugins in the user's home directory. This is where they will be
// if they were installed through Oi itself.
type HomePluginSearcher struct{}

// Search searches for plugin commands
func (HomePluginSearcher) Search() map[string]string {
	plugins := make(map[string]string)

	files, err := filepath.Glob(fmt.Sprintf("%s/*", oipath))
	if err != nil {
		log.Fatalf("Error searching for plugins in %s: %s", oipath, err)
	}

	// Strip the path and 'oi-' prefix from the commands.
	for _, file := range files {
		name := filepath.Base(file)
		if name[:3] == "oi-" {
			name = name[3:]
		}
		plugins[name] = file
	}
	return plugins
}
