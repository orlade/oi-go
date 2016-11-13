package oi

import (
	"log"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// COMMANDS

// Command is the interface that we're exposing as a plugin.
//
// Plugins should implement the Command interface like so:
//
// type FooCommand struct{}
// func (FooCommand) Invoke(args ...string) string {
//     doSomething()
//     return "todo: result"
// }
type Command interface {
	Invoke(args ...string) string
}

// CommandConfig defines the details of a Command.
type CommandConfig struct {
	Name        string
	Description string
}

// CommandRPC implements Command by invoking a remote plugin server.
type CommandRPC struct {
	client *rpc.Client
}

// Invoke calls a server over the CommandRPC's client.
func (c *CommandRPC) Invoke(args ...string) string {
	var resp string
	err := c.client.Call("Plugin.Invoke", args, &resp)
	if err != nil {
		log.Println(err)
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// CommandRPCServer is the server that CommandRPC talks to, conforming to the net/rpc requirements.
type CommandRPCServer struct {
	// This is the real implementation
	Impl Command
}

// Invoke delegates the plugin invocation to the Command's implementation.
func (s *CommandRPCServer) Invoke(args []string, resp *string) error {
	*resp = s.Impl.Invoke(args...)
	return nil
}

// CLIENT PLUGIN

// ClientPlugin provides a plugin client to call the remote plugin implementation.
type ClientPlugin struct{}

// Server returns nothing, since a ClientPlugin doesn't serve anything.
func (ClientPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, nil
}

// Client returns an RPC implementation of Command that will call a remote plugin.
func (ClientPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &CommandRPC{client: c}, nil
}

// SERVER PLUGIN

// ServerPlugin is a base struct for plugin servers to implement.
type ServerPlugin struct {
	Command
}

// Server returns a CommandRPCServer that will call Invoke() on the ServerPlugin's Command.
func (s ServerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &CommandRPCServer{Impl: s.Command}, nil
}

// Client returns nothing, since ServerPlugins are never used as clients.
func (ServerPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return nil, nil
}
