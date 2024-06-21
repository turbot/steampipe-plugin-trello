package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloMyNotification(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_my_notification",
		Description: "Get details of your notifications.",
		List: &plugin.ListConfig{
			Hydrate: listMyNotifications,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for notification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "date",
				Description: "The timestamp of when the notification was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "date_read",
				Description: "The timestamp of when the notification was read.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "id_member_creator",
				Description: "The id of the member who created the notification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_action",
				Description: "The id of the action of the notification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of notification.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "unread",
				Description: "Indicates whether the notification is unread.",
				Type:        proto.ColumnType_BOOL,
			},

			// JSON fields
			{
				Name:        "data",
				Description: "Additional data related to the notification, which varies depending on the notification type.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the notification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listMyNotifications(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_my_notification.listMyNotifications", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	notifications, err := client.GetMyNotifications(args)
	if err != nil {
		logger.Error("trello_my_notification.listMyNotifications", "api_error", err)
		return nil, err
	}

	for _, notification := range notifications {
		d.StreamListItem(ctx, notification)
	}

	return nil, nil
}
