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

func tableTrelloList(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_list",
		Description: "Get details of a list.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AnyColumn([]string{"id_board"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           listLists,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getList,
		},
		Columns: getListColumns(),
	}
}

func getListColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "The id of the list.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "name",
			Description: "The full name of the list.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "closed",
			Description: "Whether the list is closed.",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "id_board",
			Description: "The id of the board the list belongs to.",
			Transform:   transform.FromField("IDBoard"),
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "pos",
			Description: "The position of the list.",
			Type:        proto.ColumnType_DOUBLE,
		},
		{
			Name:        "subscribed",
			Description: "Whether the list has been subscribed.",
			Type:        proto.ColumnType_BOOL,
		},

		// JSON fields
		{
			Name:        "board",
			Description: "The board the list belongs to.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "cards",
			Description: "The cards in the list.",
			Type:        proto.ColumnType_JSON,
		},

		// Standard Steampipe columns
		{
			Name:        "title",
			Description: "The title of the list.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Name"),
		},
	}
}

//// LIST FUNCTION

func listLists(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	boardId := d.EqualsQualString("id_board")

	// Return nil if the id is empty
	if boardId == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_list.listLists", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var lists []trello.List

	path := path.Join("boards", boardId, "lists")
	error := client.Get(path, args, &lists)
	if error != nil {
		logger.Error("trello_list.listLists", "api_error", error)
		return nil, error
	}

	for _, list := range lists {
		d.StreamListItem(ctx, list)
	}

	return nil, nil
}

func getList(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// Return nil if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_list.listLists", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	list, err := client.GetList(id, args)
	if err != nil {
		logger.Error("trello_list.listLists", "api_error", err)
		return nil, err
	}

	return list, nil
}
