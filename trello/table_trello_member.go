package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableTrelloMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_member",
		Description: "Get details of a member.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AnyColumn([]string{"email", "full_name"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
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

func listMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	var email, fullName string
	email = d.EqualsQualString("email")
	fullName = d.EqualsQualString("full_name")
	var query string

	if email != "" {
		query += "email:" + email + " "
	}
	if fullName != "" {
		query += "fullName:" + fullName
	}

	// Return nil if the query is empty
	if query == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_member.listMembers", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	members, err := client.SearchMembers(query, args)
	if err != nil {
		logger.Error("trello_member.listMembers", "api_error", err)
		return nil, err
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
