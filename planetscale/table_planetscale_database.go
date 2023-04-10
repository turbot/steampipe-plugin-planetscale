package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePlanetScaleDatabase(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_database",
		Description: "Databases in the PlanetScale account.",
		List: &plugin.ListConfig{
			Hydrate: listDatabase,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getDatabase,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the database."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Slug"), Description: "Region where the database is located."},
			{Name: "notes", Type: proto.ColumnType_STRING, Description: "Notes for the database."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the database was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the database was updated."},
		},
	}
}

func listDatabase(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database.listDatabase", "connection_error", err)
		return nil, err
	}

	// list all databases for the given organization
	opts := &planetscale.ListDatabasesRequest{Organization: organization(ctx, d)}
	items, err := conn.Databases.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database.listDatabase", "query_error", err, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getDatabase(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database.getDatabase", "connection_error", err)
		return nil, err
	}
	name := d.EqualsQuals["name"].GetStringValue()
	opts := &planetscale.GetDatabaseRequest{Organization: organization(ctx, d), Database: name}
	item, err := conn.Databases.Get(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database.getDatabase", "query_error", err, "opts", opts)
		return nil, err
	}
	return item, err
}
