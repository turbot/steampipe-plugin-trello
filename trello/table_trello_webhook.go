package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloWebhook(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_webhook",
		Description: "Get details of the webhook.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AnyColumn([]string{"id"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           listWebhooks,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the webhook.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "active",
				Description: "Whether the webhook is active.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "callback_url",
				Description: "The callback url of the webhook.",
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

	id := d.EqualsQualString("id")

	// Return if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_webhook.listWebhooks", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	webhook, err := client.GetWebhook(id, args)
	if err != nil {
		logger.Error("trello_webhook.listWebhooks", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, webhook)

	return nil, nil
}
