package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloSearchBoard(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_search_board",
		Description: "Get details of a board.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("query"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			ParentHydrate:     listMyOrganizations,
			Hydrate:           searchBoards,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the board.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The full name of the board.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "query",
				Description: "The query provided for the search.",
				Transform:   transform.FromQual("query"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "closed",
				Description: "Whether the board is closed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "description",
				Description: "The description of the board.",
				Hydrate:     getBoard,
				Transform:   transform.FromField("Desc"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_organization",
				Description: "The ID of the organization that the board belongs to.",
				Transform:   transform.FromField("IDOrganization"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "pinned",
				Description: "Whether the board is pinned.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "short_url",
				Description: "The short URL of the board.",
				Hydrate:     getBoard,
				Transform:   transform.FromField("ShortURL"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "starred",
				Description: "Whether the board is starred.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "subscribed",
				Description: "Whether the board has been subscribed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "url",
				Description: "The URL of the board.",
				Hydrate:     getBoard,
				Transform:   transform.FromField("URL"),
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "label_names",
				Description: "The label names of the board.",
				Hydrate:     getBoard,
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "prefs",
				Description: "The preferences of the board.",
				Hydrate:     getBoard,
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the board.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func searchBoards(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	query := d.EqualsQualString("query")

	// Return nil if query is empty
	if query == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_search_board.searchBoards", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	boards, err := client.SearchBoards(query, args)
	if err != nil {
		logger.Error("trello_search_board.searchBoards", "api_error", err)
		return nil, err
	}

	for _, board := range boards {
		d.StreamListItem(ctx, board)
	}

	return nil, nil
}
