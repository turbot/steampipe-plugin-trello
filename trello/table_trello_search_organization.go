package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloSearchOrganization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_search_organization",
		Description: "Get details of a organization.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("query"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			ParentHydrate:     listMyOrganizations,
			Hydrate:           searchOrganizations,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The full name of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "query",
				Description: "The query provided for the search.",
				Transform:   transform.FromQual("query"),
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
			{
				Name:        "tags",
				Description: "The organization tags.",
				Hydrate:     getOrganizationTags,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func searchOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	query := d.EqualsQualString("query")

	// Return nil if query is empty
	if query == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_search_organization.searchOrganizations", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{
		"query":      query,
		"modelTypes": "organizations",
	}

	var organizations []trello.Organization

	error := client.Get("search", args, organizations)
	if error != nil {
		logger.Error("trello_search_organization.searchOrganizations", "api_error", error)
		return nil, err
	}

	for _, organization := range organizations {
		d.StreamListItem(ctx, organization)
	}

	return nil, nil
}
