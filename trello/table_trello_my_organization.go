package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloMyOrganization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_my_organization",
		Description: "Get details of my organizations.",
		List: &plugin.ListConfig{
			Hydrate: listMyOrganizations,
		},
		Columns: getOrganizationColumns(),
	}
}

func getOrganizationColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "The id of the organization.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "name",
			Description: "The name of the organization.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "description",
			Description: "The description of the organization.",
			Transform:   transform.FromField("Desc"),
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "display_name",
			Description: "The display name of the organization.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "member_id",
			Description: "The id of the member.",
			Transform:   transform.FromQual("member_id"),
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "url",
			Description: "The URL of the organization.",
			Transform:   transform.FromField("URL"),
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "website",
			Description: "Link to the organization's website.",
			Type:        proto.ColumnType_STRING,
		},

		// JSON fields
		{
			Name:        "products",
			Description: "The products of the organization.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "power_ups",
			Description: "The power ups that are a part of the organization.",
			Type:        proto.ColumnType_JSON,
		},

		// Standard Steampipe columns
		{
			Name:        "title",
			Description: "The title of the organization.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Name"),
		},
	}
}

//// LIST FUNCTION

func listMyOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_my_organization.listMyOrganizations", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var organizations []trello.Organization

	error := client.Get("members/me/organizations", args, &organizations)
	if error != nil {
		logger.Error("trello_my_organization.listMyOrganizations", "api_error", err)
		return nil, err
	}

	for _, organization := range organizations {
		d.StreamListItem(ctx, organization)
	}

	return nil, nil
}
