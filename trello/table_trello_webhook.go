package trello

import (
	"context"
	"path"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloWebhook(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_webhook",
		Description: "Get details of the webhooks.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AnyColumn([]string{"id_token"}),
			ShouldIgnoreError: isNotFoundError([]string{"400: invalid token"}),
			Hydrate:           listWebhooks,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"),
			ShouldIgnoreError: isNotFoundError([]string{"400: invalid id"}),
			Hydrate:           getWebhook,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the webhook.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "id_token",
				Description: "The id of the token the webhook belongs to.",
				Transform:   transform.FromQual("id_token"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active",
				Description: "Indicates whether the webhook is active.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "callback_url",
				Description: "The callback url for the webhook.",
				Transform:   transform.FromField("CallbackURL"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the webhook.",
				Transform:   transform.FromField("Desc"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_model",
				Description: "The id of the model of the webhook.",
				Transform:   transform.FromField("IDModel"),
				Type:        proto.ColumnType_STRING,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the webhook.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

//// LIST FUNCTION

func listWebhooks(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id_token")

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_webhook.listWebhooks", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var webhooks []*trello.Webhook

	path := path.Join("tokens", id, "webhooks")
	error := client.Get(path, args, &webhooks)
	if error != nil {
		logger.Error("trello_webhook.listWebhooks", "api_error", error)
		return nil, error
	}

	for _, webhook := range webhooks {
		d.StreamListItem(ctx, webhook)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getWebhook(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Return nil if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_webhook.getWebhook", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	webhook, err := client.GetWebhook(id, args)
	if err != nil {
		logger.Error("trello_webhook.getWebhook", "api_error", err)
		return nil, err
	}

	return webhook, nil
}
