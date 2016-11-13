package oi

import (
	"fmt"
	"os/user"
	"path/filepath"

	"log"
)

var usr, _ = user.Current()
var OIPATH = usr.HomeDir + "/.oi/modules"

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

	files, err := filepath.Glob(fmt.Sprintf("%s/*", OIPATH))
	if err != nil {
		log.Fatalf("Error searching for plugins in %s: %s", OIPATH, err)
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
