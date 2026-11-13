package trello

import (
	"context"
	"errors"
	"os"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func connectTrello(ctx context.Context, d *plugin.QueryData) (*trello.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "trello"
	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return cachedData.(*trello.Client), nil
	}

	// Default to using env vars (#2)
	apiKey := os.Getenv("TRELLO_API_KEY")
	token := os.Getenv("TRELLO_TOKEN")

	// But prefer the config (#1)
	trelloConfig := GetConfig(d.Connection)
	if trelloConfig.APIKey != nil {
		apiKey = *trelloConfig.APIKey
	}
	if trelloConfig.Token != nil {
		token = *trelloConfig.Token
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, errors.New("api_key must be configured")
	}
	if token == "" {
		return nil, errors.New("token must be configured")
	}

	client := trello.NewClient(apiKey, token)

	if client != nil {
		if err := d.ConnectionCache.Set(ctx, cacheKey, client); err != nil {
			plugin.Logger(ctx).Warn("connectTrello", "cache_set_error", err)
		}
	}

	return client, nil
}
