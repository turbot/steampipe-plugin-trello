package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-trello/trello"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: trello.Plugin})
}
