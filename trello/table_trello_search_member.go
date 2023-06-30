package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloSearchMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_search_member",
		Description: "Get details of a member.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("query"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			ParentHydrate:     listMyOrganizations,
			Hydrate:           searchMembers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the member.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "query",
				Description: "The query provided for the search.",
				Transform:   transform.FromQual("query"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "username",
				Description: "The username of the member.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "avatar_hash",
				Description: "The hash of the avatar of the member.",
				Hydrate:     getMember,
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The email address of the member.",
				Hydrate:     getMember,
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
				Hydrate:     getMember,
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "id_boards",
				Description: "An array of board IDs that the member is on.",
				Hydrate:     getMember,
				Transform:   transform.FromField("IDBoards"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id_organizations",
				Description: "An array of organization IDs that the member belongs to.",
				Hydrate:     getMember,
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

func searchMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	query := d.EqualsQualString("query")

	// Return nil if query is empty
	if query == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_search_member.searchMembers", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	members, err := client.SearchMembers(query, args)
	if err != nil {
		logger.Error("trello_search_member.searchMembers", "api_error", err)
		return nil, err
	}

	for _, member := range members {
		d.StreamListItem(ctx, member)
	}

	return nil, nil
}
