package trello

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type trelloConfig struct {
	TrelloAPIKey *string `cty:"trello_api_key"`
	TrelloToken  *string `cty:"trello_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"trello_api_key": {
		Type: schema.TypeString,
	},
	"trello_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &trelloConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) trelloConfig {
	if connection == nil || connection.Config == nil {
		return trelloConfig{}
	}
	config, _ := connection.Config.(trelloConfig)
	return config
}
