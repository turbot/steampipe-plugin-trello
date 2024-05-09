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
		Name:               pluginName,
		DefaultTransform:   transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		// Member ID would be same per connection.
		// API key and API Token is specific to a member.
		// A member can have multiple organizations, workspaces, boards, etc...
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "member_id",
				Hydrate: getMemberId,
			},
		},
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"trello_board":           tableTrelloBoard(ctx),
			"trello_card":            tableTrelloCard(ctx),
			"trello_list":            tableTrelloList(ctx),
			"trello_member":          tableTrelloMember(ctx),
			"trello_my_board":        tableTrelloMyBoard(ctx),
			"trello_my_member":       tableTrelloMyMember(ctx),
			"trello_my_notification": tableTrelloMyNotification(ctx),
			"trello_my_organization": tableTrelloMyOrganization(ctx),
			"trello_organization":    tableTrelloOrganization(ctx),
			"trello_search_board":    tableTrelloSearchBoard(ctx),
			"trello_search_card":     tableTrelloSearchCard(ctx),
			"trello_token":           tableTrelloToken(ctx),
			"trello_webhook":         tableTrelloWebhook(ctx),
		},
	}

	return p
}
