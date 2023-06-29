package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableTrelloBoard(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_board",
		Description: "Get details of a board.",
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
		Columns: getBoardColumns(),
	}
}

//// LIST FUNCTION

func listBoards(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := h.Item.(*trello.Organization).ID

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

func getBoard(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

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
