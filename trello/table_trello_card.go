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

func tableTrelloCard(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_card",
		Description: "Get details of all cards in a list.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AnyColumn([]string{"id_list"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           listCards,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCard,
		},
		Columns: []*plugin.Column{
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
				Transform:   transform.FromField("Desc"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "due",
				Description: "The due date of the card, if set.",
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
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id_attachment_cover",
				Description: "The id of the attachment used as the card cover.",
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
				Transform:   transform.FromField("IDMembers"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "id_members_voted",
				Description: "The ids of members who voted on the card.",
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
		},
	}
}

//// LIST FUNCTION

func listCards(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	listId := d.EqualsQualString("id_list")

	// Return nil if the id is empty
	if listId == "" {
		return nil, nil
	}

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_card.listCards", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var cards []trello.Card

	path := path.Join("lists", listId, "cards")
	error := client.Get(path, args, &cards)
	if error != nil {
		logger.Error("trello_card.listCards", "api_error", error)
		return nil, error
	}

	for _, card := range cards {
		d.StreamListItem(ctx, card)
	}

	return nil, nil
}

func getCard(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	var id string
	if h.Item != nil {
		id = h.Item.(trello.Card).ID
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
		logger.Error("trello_card.listCards", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	card, error := client.GetCard(id, args)
	if error != nil {
		logger.Error("trello_card.listCards", "api_error", err)
		return nil, error
	}

	return *card, nil
}

func getCardAttachments(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(trello.Card).ID

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_card.getCardAttachments", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var attachments []trello.Attachment

	path := path.Join("cards", id, "attachments")
	error := client.Get(path, args, &attachments)
	if error != nil {
		logger.Error("trello_card.getCardAttachments", "api_error", error)
		return nil, error
	}

	return attachments, nil
}

func getCardChecklists(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(trello.Card).ID

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_board.getCardChecklists", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var checklists []trello.Checklist

	path := path.Join("cards", id, "checklists")
	error := client.Get(path, args, &checklists)
	if error != nil {
		logger.Error("trello_board.getCardChecklists", "api_error", error)
		return nil, error
	}

	return checklists, nil
}

func getCardCustomFieldItems(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := h.Item.(trello.Card).ID

	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("trello_board.getCardCustomFieldItems", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}
	var customFields []trello.CustomFieldItem

	path := path.Join("cards", id, "customFieldItems")
	error := client.Get(path, args, &customFields)
	if error != nil {
		logger.Error("trello_board.getCardCustomFieldItems", "api_error", error)
		return nil, error
	}

	return customFields, nil
}
