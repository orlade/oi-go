# Oi

A CLI tool for automating development tasks.

## Design

Oi's command-line interface is structured like that of Git. After the initial `oi` command follows
the name of a `module`, and after that is an `action` to perform on the module. Options can be
specified at any point after the `oi` command.

### Plugins

Oi is designed to be a framework with which modular, extensible CLI tools can be built. Using 
Hashicorp's [go-plugin][go-plugin] library, Oi scans for plugins that define modules, each of which
registers its own command and associated actions. New project-specific tools can then be built by 
aggregating the commands for existing tools, and extending them for particular use cases.

### Creating a Plugin

An Oi plugin is a Go program that implements a `go-plugin` `Server`.

TODO(orlade): More details.

## Components

`oi/cli` builds the main `oi` command. Most development automation will be implemented as plugins,
so the Oi CLI exposes those plugins through a consistent user interface.

The files in the root of the `oi` repository make up the Oi core library, mainly specifying and 
invoking plugins.  


[go-plugin]: https://github.com/hashicorp/go-plugin
