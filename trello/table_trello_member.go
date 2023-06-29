package trello

import (
	"context"
	"fmt"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableTrelloMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_member",
		Description: "Get details of a member.",
		List: &plugin.ListConfig{
			// KeyColumns:        plugin.OptionalColumns([]string{"id_organizations"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			ParentHydrate:     listMyOrganizations,
			Hydrate:           listMembers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "username"}),
			Hydrate:    getMember,
		},
		Columns: getMemberColumns(),
	}
}

//// LIST FUNCTION

func listMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	organizationId := h.Item.(*trello.Organization).ID

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_member.listMembers", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var members []*trello.Member

	path := fmt.Sprintf("organizations/%s/members", organizationId)
	error := client.Get(path, args, members)
	if error != nil {
		logger.Error("trello_member.listMembers", "api_error", error)
		return nil, error
	}

	for _, member := range members {
		d.StreamListItem(ctx, member)
	}

	return nil, nil
}

func getMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")
	if id == "" {
		id = d.EqualsQualString("username")
	}

	// Return if the id is empty
	if id == "" {
		return nil, nil
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

	d.StreamListItem(ctx, member)

	return nil, nil
}
