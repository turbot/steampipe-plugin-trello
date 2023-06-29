package trello

import (
	"context"
	"fmt"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTrelloCard(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "trello_card",
		Description: "Get details of a card.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.AnyColumn([]string{"id_list"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           listCards,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCard,
		},
		Columns: getCardColumns(),
	}
}

func getCardColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "The id of the card.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "name",
			Description: "The full name of the card.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "closed",
			Description: "Whether the card is closed.",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "date_last_activity",
			Description: "The date of the last activity on the card.",
			Type:        proto.ColumnType_TIMESTAMP,
		},
		{
			Name:        "description",
			Description: "The description of the card.",
			Transform:   transform.FromField("Desc"),
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "due",
			Description: "The estimated due time of the card.",
			Type:        proto.ColumnType_TIMESTAMP,
		},
		{
			Name:        "due_complete",
			Description: "Whether the task is complete.",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "email",
			Description: "The email id associated with the card.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "id_attachment_cover",
			Description: "The id of the attachment cover of the card.",
			Transform:   transform.FromField("IDAttachmentCover"),
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "id_board",
			Description: "The id of the board.",
			Transform:   transform.FromField("IDShort"),
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "id_list",
			Description: "The id of the list.",
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
			Description: "The short link of the card.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "short_url",
			Description: "The short URL of the card.",
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
			Description: "Whether the card has been subscribed.",
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
			Name:        "actions",
			Description: "The actions of the card.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "attachments",
			Description: "The attachments of the card.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "badges",
			Description: "The badges of the card.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "board",
			Description: "The board of the card.",
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
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "custom_field_items",
			Description: "The custom field items of the card.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "custom_field_map",
			Description: "The custom field map of the card.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "id_check_lists",
			Description: "The check list ids of the card.",
			Transform:   transform.FromField("IDCheckLists"),
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "id_labels",
			Description: "The label ids of the card.",
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
			Description: "The member ids of the card.",
			Transform:   transform.FromField("IDMembers"),
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "id_members_voted",
			Description: "The member voted ids of the card.",
			Transform:   transform.FromField("IDMembersVoted"),
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "list",
			Description: "The list of the card.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "members",
			Description: "The members of the card.",
			Type:        proto.ColumnType_JSON,
		},

		// Standard Steampipe columns
		{
			Name:        "title",
			Description: "The title of the card.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Name"),
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
	var cards []*trello.Card

	path := fmt.Sprintf("lists/%s/cards", listId)
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

func getCard(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

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

	return card, nil
}
