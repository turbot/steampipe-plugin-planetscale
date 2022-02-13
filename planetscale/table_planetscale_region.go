package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tablePlanetScaleRegion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_region",
		Description: "Regions in the PlanetScale account.",
		List: &plugin.ListConfig{
			Hydrate: listRegion,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "Slug of the region."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Display name of the region."},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "Location for the region."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: "True if the region is enabled."},
		},
	}
}

func listRegion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_region.listRegion", "connection_error", err)
		return nil, err
	}
	opts := &planetscale.ListRegionsRequest{}
	items, err := conn.Regions.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_region.listRegion", "query_error", err, "opts", opts)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
