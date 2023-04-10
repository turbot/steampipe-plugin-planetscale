package planetscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tablePlanetScaleOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_organization",
		Description: "Organizations in the PlanetScale account.",
		List: &plugin.ListConfig{
			Hydrate: listOrganization,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the organization was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the organization was updated."},
		},
	}
}

func listOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_organization.listOrganization", "connection_error", err)
		return nil, err
	}

	// list all organizations for the given organization
	items, err := conn.Organizations.List(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_organization.listOrganization", "query_error", err)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
