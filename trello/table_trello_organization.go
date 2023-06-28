package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableTrelloOrganization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_organization",
		Description: "Get details of the organization.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AnyColumn([]string{"id"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           listOrganizations,
		},
		Columns: getOrganizationColumns(),
	}
}

//// LIST FUNCTION

func listOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// Return if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_organization.listOrganizations", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	organization, err := client.GetOrganization(id, args)
	if err != nil {
		logger.Error("trello_organization.listOrganizations", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, organization)

	return nil, nil
}
