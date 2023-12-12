package trello

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type trelloConfig struct {
	APIKey *string `hcl:"api_key"`
	Token  *string `hcl:"token"`
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
