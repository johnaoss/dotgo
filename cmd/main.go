package main

import (
	"os"
	"path/filepath"

	"github.com/johnaoss/dotgo"
	"github.com/johnaoss/dotgo/messenger"

	"github.com/alexflint/go-arg"
)

const (
	version = "0.0.1"
)

var (
	args Arguments
	log  = messenger.Messenger()
)

// Arguments is the struct representing the command line arguments of the program
type Arguments struct {
	SuperQuiet bool `arg:"-Q, --super-quiet" help:"suppress almost all output"`
	Quiet      bool `arg:"-q" help:"suppress most output"`
	Verbose    bool `arg:"-v" help:"enable verbose output"`
	// TODO: Fix this one to use the `BASEDIR` metavar thing.
	BaseDirectory string `arg:"-d, --base-directory" help:"execute commands from within BASEDIR"`
	// TODO: Fix this one to use the `CONFIGFILE` metavar.
	Configfile            string   `arg:"-c, --config-file" help:"run commands given in CONFIGFILE"`
	Plugins               []string `arg:"-p, --plugin" help:"load PLUGIN as a plugin"`
	DisableBuiltInPlugins bool     `arg:"--disable-built-in-plugins" help:"disable built-in plugins"`
	// TODO: Fix this one to use the `PLUGIN_DIR` metavar.
	PluginDirs []string `arg:"--plugin-dir" help:"load all plugins in PLUGIN_DIR"`
	NoColor    bool     `arg:"--no-color" help:"disable color output"`
}

func (a Arguments) Version() string {
	return "dotgo version " + version
}

func main() {
	arg.MustParse(&args)

	configPath, err := getConfigPath()
	if err != nil {
		log.Error("Failed to parse configfile path, given error: " + err.Error())
		os.Exit(1)
	}

	_, err = dotgo.ReadConfig(configPath)
	if err != nil {
		log.Error("Failed to parse config file, given error: " + err.Error())
		os.Exit(1)
	}
}

func getConfigPath() (string, error) {
	if args.Configfile != "" {
		return args.Configfile, nil
	}
	log.Info("No ConfigFile specified, switching to one at $HOME/.config/dotgo")
	homedir, err := os.UserHomeDir()
	return filepath.Clean(homedir + "/.config/dotgo"), err

}
