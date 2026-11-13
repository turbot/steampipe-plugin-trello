package trello

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		errMsg := err.Error()
		for _, msg := range notFoundErrors {
			if strings.Contains(errMsg, msg) {
				return true
			}
		}
		return false
	}
}

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		errMsg := err.Error()
		for _, msg := range retryErrors {
			if strings.Contains(errMsg, msg) {
				plugin.Logger(ctx).Error("jira_errors.shouldRetryError", "rate_limit_error", err)
				return true
			}
		}
		return false
	}
}
