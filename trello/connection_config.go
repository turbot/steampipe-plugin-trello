package trello

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type trelloConfig struct {
	APIKey *string `cty:"api_key"`
	Token  *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"token": {
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
