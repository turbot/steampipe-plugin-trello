package trello

import (
	"context"
	"path"
	"strings"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableTrelloBoard(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_board",
		Description: "Get details of all the boards in your organization.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.OptionalColumns([]string{"id_organization"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			ParentHydrate:     listMyOrganizations,
			Hydrate:           listBoards,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getBoard,
		},
		Columns: commonColumns(getBoardColumns()),
	}
}

//// LIST FUNCTION

func listBoards(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := h.Item.(trello.Organization).ID

	if d.EqualsQuals["id_organization"] != nil && d.EqualsQualString("id_organization") != id {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_board.listBoards", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	boards, err := client.GetBoardsInOrganization(id, args)
	if err != nil {
		logger.Error("trello_board.listBoards", "api_error", err)
		return nil, err
	}

	for _, board := range boards {
		d.StreamListItem(ctx, board)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getBoard(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	var id string
	if h.Item != nil {
		id = h.Item.(*trello.Board).ID
	} else {
		id = d.EqualsQualString("id")
	}

	// Return nil if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_board.getBoard", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	board, err := client.GetBoard(id, args)
	if err != nil {
		logger.Error("trello_board.getBoard", "api_error", err)
		return nil, err
	}

	return board, nil
}

func getBoardCustomFields(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(*trello.Board).ID

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_board.getBoardCustomFields", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var customFields []trello.CustomField

	path := path.Join("boards", id, "customFields")
	error := client.Get(path, args, &customFields)
	if error != nil {
		if strings.Contains(error.Error(), "non-pointer") {
			return nil, nil
		}
		logger.Error("trello_board.getBoardCustomFields", "api_error", error)
		return nil, error
	}

	return customFields, nil
}
