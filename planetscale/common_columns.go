package planetscale

import (
	"context"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "organization_name",
			Description: "The name of the organization.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getOrganizationName,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getOrganizationNameMemoized = plugin.HydrateFunc(getOrganizationNameUncached).Memoize(memoize.WithCacheKeyFunction(getOrganizationNameCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getOrganizationName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getOrganizationNameMemoized(ctx, d, h)
}

// Build a cache key for the call to getOrganizationNameCacheKey.
func getOrganizationNameCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getOrganizationName"
	return key, nil
}

func getOrganizationNameUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	organization := os.Getenv("PLANETSCALE_ORGANIZATION")

	// Prefer config settings
	planetscaleConfig := GetConfig(d.Connection)
	if planetscaleConfig.Organization != nil {
		organization = *planetscaleConfig.Organization
	}

	return organization, nil
}
