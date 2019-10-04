// Package plugins provides the default plugins available to be used by the dotgo
// executable.
package plugins

type Plugin interface {
	CanHandle(directive string) bool
	Handle(directive string, data map[string]interface{}) error

	// Plugins in Go are hard, this is a _BIG_ todo to actually figure out how
	// to register plugins and use them accordingly.
	// Could use https://github.com/containous/yaegi to facilitate something,
	// but would require a weird sort of validation to ensure that once plugins are loaded
	// that they provide the required functions.
	// Essentially an interface but for a package?
}

// defaultPlugins should contain a list of the available plugins from this package.
var defaultPlugins = []Plugin{
	newClean(),
}

// GetPlugins is used to allow us to call the `Load` function on ourselves.
func GetPlugins() []Plugin {
	return defaultPlugins
}
