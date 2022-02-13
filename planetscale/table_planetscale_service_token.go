package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tablePlanetScaleServiceToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_service_token",
		Description: "Service tokens in the PlanetScale account.",
		List: &plugin.ListConfig{
			Hydrate: listServiceToken,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the service token."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the token."},
			// Always null? {Name: "token", Type: proto.ColumnType_STRING, Description: "The service token."},
		},
	}
}

func listServiceToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_service_token.listServiceToken", "connection_error", err)
		return nil, err
	}

	// list all audit logs for the given organization
	opts := &planetscale.ListServiceTokensRequest{Organization: organization(ctx, d)}
	items, err := conn.ServiceTokens.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_service_token.listServiceToken", "query_error", err, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
