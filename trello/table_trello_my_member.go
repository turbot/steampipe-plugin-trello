package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloMyMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_my_member",
		Description: "Get details of my trello member.",
		List: &plugin.ListConfig{
			Hydrate: listMyMembers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "username",
				Description: "The username of the member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "avatar_hash",
				Description: "The hash of the avatar of the member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The email address of the member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "full_name",
				Description: "The full name of the member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "initials",
				Description: "The initials of the member.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "id_boards",
				Description: "An array of board IDs that the member is on.",
				Transform:   transform.FromField("IDBoards"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id_organizations",
				Description: "An array of organization IDs that the member belongs to.",
				Transform:   transform.FromField("IDOrganizations"),
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FullName"),
			},
		},
	}
}

//// LIST FUNCTION

func listMyMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_my_member.listMyMembers", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	member, err := client.GetMyMember(args)
	if err != nil {
		logger.Error("trello_my_member.listMyMembers", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, *member)

	return nil, nil
}
