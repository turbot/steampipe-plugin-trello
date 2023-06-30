package trello

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-trello"

// Plugin creates this (trello) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"trello_board":               tableTrelloBoard(ctx),
			"trello_card":                tableTrelloCard(ctx),
			"trello_list":                tableTrelloList(ctx),
			"trello_member":              tableTrelloMember(ctx),
			"trello_my_board":            tableTrelloMyBoard(ctx),
			"trello_my_member":           tableTrelloMyMember(ctx),
			"trello_my_notification":     tableTrelloMyNotification(ctx),
			"trello_my_organization":     tableTrelloMyOrganization(ctx),
			"trello_organization":        tableTrelloOrganization(ctx),
			"trello_search_board":        tableTrelloSearchBoard(ctx),
			"trello_search_card":         tableTrelloSearchCard(ctx),
			"trello_search_member":       tableTrelloSearchMember(ctx),
			"trello_search_organization": tableTrelloSearchOrganization(ctx),
			"trello_webhook":             tableTrelloWebhook(ctx),
		},
	}

	return p
}
