package trello

import (
	"context"
	"strconv"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloSearchCard(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_search_card",
		Description: "Find cards you have access to that match search keywords.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("query"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           searchCards,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique identifier for the card.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the card.",
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
				Description: "Indicates whether the card is closed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "date_last_activity",
				Description: "The timestamp of the last activity on the card.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "The description or summary of the card.",
				Hydrate:     getCard,
				Transform:   transform.FromField("Desc"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "due",
				Description: "The due date of the card, if set.",
				Hydrate:     getCard,
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "due_complete",
				Description: "Indicates whether the due date of the card is complete.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "email",
				Description: "The email id associated with the card.",
				Hydrate:     getCard,
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_attachment_cover",
				Description: "The id of the attachment used as the card cover.",
				Hydrate:     getCard,
				Transform:   transform.FromField("IDAttachmentCover"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_board",
				Description: "The id of the board the card belongs to.",
				Transform:   transform.FromField("IDShort"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_list",
				Description: "The id of the list the card belongs to.",
				Transform:   transform.FromField("IDList"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_short",
				Description: "The short id of the card.",
				Transform:   transform.FromField("IDShort"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "manual_cover_attachment",
				Description: "The manual cover attachment of the card.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "pos",
				Description: "The position of the card.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "short_link",
				Description: "The shortened link of the card.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "short_url",
				Description: "The shortened URL of the card.",
				Transform:   transform.FromField("ShortURL"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "start",
				Description: "The start time of the card.",
				Hydrate:     getCard,
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "subscribed",
				Description: "Indicates whether the card has been subscribed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "url",
				Description: "The URL of the card.",
				Transform:   transform.FromField("URL"),
				Type:        proto.ColumnType_STRING,
			},

			// JSON fields
			{
				Name:        "attachments",
				Description: "The attachments of the card.",
				Hydrate:     getCardAttachments,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "badges",
				Description: "The badges of the card.",
				Hydrate:     getCard,
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "check_item_states",
				Description: "The check item states of the card.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "checklists",
				Description: "The checklists of the card.",
				Hydrate:     getCardChecklists,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "custom_field_items",
				Description: "The custom field items of the card.",
				Hydrate:     getCardCustomFieldItems,
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "custom_field_map",
				Description: "The custom field map of the card.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id_check_lists",
				Description: "The ids of checklists attached to the card.",
				Hydrate:     getCard,
				Transform:   transform.FromField("IDCheckLists"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id_labels",
				Description: "The ids of labels attached to the card.",
				Transform:   transform.FromField("IDLabels"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "labels",
				Description: "The labels of the card.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id_members",
				Description: "The ids of members attached to the card.",
				Hydrate:     getCard,
				Transform:   transform.FromField("IDMembers"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id_members_voted",
				Description: "The ids of members who voted on the card.",
				Hydrate:     getCard,
				Transform:   transform.FromField("IDMembersVoted"),
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			{
				Name:        "title",
				Description: "The title of the card.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func searchCards(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	query := d.EqualsQualString("query")

	// Return nil if query is empty
	if query == "" {
		return nil, nil
	}

	// Limiting the results
	maxLimit := int32(1000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_search_card.searchCards", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{
		"cards_limit": strconv.Itoa(int(maxLimit)),
		"cards_page":  "0",
	}

	for {
		cards, err := client.SearchCards(query, args)
		if err != nil {
			logger.Error("trello_search_card.searchCards", "api_error", err)
			return nil, err
		}

		for _, card := range cards {
			d.StreamListItem(ctx, *card)
		}

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}

		if len(cards) < int(maxLimit) {
			return nil, nil
		} else {
			page, _ := strconv.Atoi(args["cards_page"])
			nextPage := page + 1
			args["cards_page"] = strconv.Itoa(int(nextPage))
		}
	}
}
