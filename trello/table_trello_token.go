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

func tableTrelloToken(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_token",
		Description: "Get details of the tokens.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("id_member"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           listTokens,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getToken,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the token.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "id_member",
				Description: "The id of the member of the token.",
				Transform:   transform.FromField("IDMember"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The timestamp of when the token was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "date_expires",
				Description: "The timestamp of when the token expires.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "identifier",
				Description: "The identifier of the token.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "permissions",
				Description: "The permissions of the token.",
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the token.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Identifier"),
			},
		},
	}
}

//// LIST FUNCTION

func listTokens(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id_member")

	// Return if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_token.listTokens", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var tokens []trello.Token

	path := path.Join("members", id, "tokens")
	error := client.Get(path, args, &tokens)
	if error != nil {
		logger.Error("trello_token.listTokens", "api_error", error)
		return nil, err
	}

	for _, token := range tokens {
		d.StreamListItem(ctx, token)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// Return nil if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_token.getToken", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	token, error := client.GetToken(id, args)
	if error != nil {
		logger.Error("trello_token.getToken", "api_error", error)
		return nil, err
	}

	return token, nil
}
