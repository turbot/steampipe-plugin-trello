package main

import (
	"github.com/turbot/steampipe-plugin-trello/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: trello.Plugin})
}
