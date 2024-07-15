package trello

import (
	"context"

	"github.com/adlio/trello"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Member ID would be same per connection.
// API key and API Token is specific to a member
// A member can have multiple organizations, workspaces, boards, etc...
func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "member_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique identifier of the member.",
			Hydrate:     getMemberId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getMemberIdMemoized = plugin.HydrateFunc(getMemberIdUncached).Memoize(memoize.WithCacheKeyFunction(getMemberIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getMemberId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getMemberIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getMemberIdCacheKey.
func getMemberIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getMemberId"
	return key, nil
}

func getMemberIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	// Create client
	client, err := connectTrello(ctx, d)
	if err != nil {
		logger.Error("getMemberIdUncached", "connection_error", err)
		return nil, err
	}

	args := trello.Arguments{}

	member, err := client.GetMyMember(args)
	if err != nil {
		logger.Error("getMemberIdUncached", "api_error", err)
		return nil, err
	}

	// Member ID would be same per connection.
	// API key and API Token is specific to a member
	// A member can have multiple organizations, workspaces, boards, etc...
	return member.ID, nil
}
