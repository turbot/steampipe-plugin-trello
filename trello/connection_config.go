package trello

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
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
	if connection == nil || connection.GetConfig() == nil {
		return trelloConfig{}
	}
	config, _ := connection.GetConfig().(trelloConfig)
	return config
}
