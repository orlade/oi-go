package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/orlade/oi"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "oi"
	app.Usage = "Automate development tasks"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable debug mode",
		},
		cli.StringFlag{
			Name:  "load, l",
			Value: "~/.oi/config.toml",
			Usage: "Load configuration from `FILE`",
		},
	}
	// TODO(orlade): Allow loading from ~/.oi file.
	// app.Before = altsrc.InitInputSourceWithContext(app.Flags, altsrc.NewYamlSourceFromFlagFunc("load"))

	// Discover available plugins and create a command for each one.
	nameCommands := new(oi.HomePluginSearcher).Search()

	for name, binary := range nameCommands {
		app.Commands = []cli.Command{
			{
				Name:  name,
				Usage: "Executes an action on " + name,
				Action: func(c *cli.Context) error {

					// Connect via RPC
					rpcClient, err := createClient(name, binary).Client()
					if err != nil {
						log.Printf("Command %s for plugin %s not found",
							color.MagentaString(name),
							color.CyanString(name))
						log.Fatalln(err)
					}

					// Request the plugin
					raw, err := rpcClient.Dispense(name)
					if err != nil {
						log.Printf("Plugin %s not found", name)
						log.Fatalln(err)
					}

					// We should have a Command now! This feels like a normal interface
					// implementation but is in fact over an RPC connection.
					Command := raw.(oi.Command)
					result := Command.Invoke(os.Args[2:]...)
					fmt.Println(result)
					return nil
				},
			},
		}
	}

	// TODO(orlade): Builtin commands like "install".
	// TODO(orlade): Allow plugins to define categories for their top-level commands.

	app.Run(os.Args)
}

// createClient is a helper function to create a ClientPlugin for an Oi plugin.
func createClient(name string, binaryPath string) (client *plugin.Client) {
	pluginMap := make(map[string]plugin.Plugin)
	pluginMap[name] = new(oi.ClientPlugin)

	client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(binaryPath),
	})
	defer client.Kill()
	return
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}
