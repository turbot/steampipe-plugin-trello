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

func tableTrelloMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_member",
		Description: "Get details of all members in an organization.",
		List: &plugin.ListConfig{
			ParentHydrate:     listMyOrganizations,
			Hydrate:           listMembers,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "username"}),
			Hydrate:    getMember,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier of the member.",
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

func listMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	data := h.Item.(trello.Organization)
	organizationId := data.ID

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_member.listMembers", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var members []trello.Member

	path := path.Join("organizations", organizationId, "members")
	error := client.Get(path, args, &members)
	if error != nil {
		logger.Error("trello_member.listMembers", "api_error", error)
		return nil, error
	}

	for _, member := range members {
		d.StreamListItem(ctx, member)
	}

	return nil, nil
}

func getMember(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	var id string
	if h.Item != nil {
		id = h.Item.(trello.Member).ID
	} else {
		id = d.EqualsQualString("id")
		if id == "" {
			id = d.EqualsQualString("username")
		}

		// Return if the id is empty
		if id == "" {
			return nil, nil
		}
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_member.getMember", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	member, err := client.GetMember(id, args)
	if err != nil {
		logger.Error("trello_member.getMember", "api_error", err)
		return nil, err
	}

	return member, nil
}
