package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tablePlanetScaleDeployRequest(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_deploy_request",
		Description: "Deploy Requests in the PlanetScale account.",
		List: &plugin.ListConfig{
			ParentHydrate: listDatabase,
			Hydrate:       listDeployRequest,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "database_name", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "database_name"},
				{Name: "number"},
			},
			Hydrate: getDeployRequest,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "Name of the database."},
			{Name: "number", Type: proto.ColumnType_INT, Transform: transform.FromField("DeployRequest.Number"), Description: "Number for this deploy request."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeployRequest.ID"), Description: "Unique ID for the deplloy request."},
			{Name: "branch", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeployRequest.Branch"), Description: "Deploy request branch."},
			{Name: "into_branch", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeployRequest.IntoBranch"), Description: "Deploy request into branch."},
			{Name: "state", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeployRequest.State"), Description: "State of the deploy request."},
			{Name: "approved", Type: proto.ColumnType_BOOL, Transform: transform.FromField("DeployRequest.Approved"), Description: "True if the deploy request is approved."},
			{Name: "notes", Type: proto.ColumnType_STRING, Transform: transform.FromField("DeployRequest.Notes"), Description: "Notes for the deploy request."},
			{Name: "deployment", Type: proto.ColumnType_JSON, Transform: transform.FromField("DeployRequest.Deployment"), Description: "Details of the deployment."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DeployRequest.CreatedAt"), Description: "When the deploy request was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DeployRequest.UpdatedAt"), Description: "When the deploy request was updated."},
			{Name: "closed_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("DeployRequest.ClosedAt"), Description: "When the deploy request was closed."},
		},
	}
}

type deployRequestRow struct {
	OrganizationName string
	DatabaseName     string
	DeployRequest    *planetscale.DeployRequest
}

func listDeployRequest(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_deploy_request.listDeployRequest", "connection_error", err)
		return nil, err
	}

	org := organization(ctx, d)

	var dbName string
	if h.Item != nil {
		dbName = h.Item.(*planetscale.Database).Name
	} else if d.KeyColumnQuals["database_name"] != nil {
		dbName = d.KeyColumnQuals["database_name"].GetStringValue()
	}

	// list all databases for the given organization
	opts := &planetscale.ListDeployRequestsRequest{Organization: org, Database: dbName}
	items, err := conn.DeployRequests.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_deploy_request.listDeployRequest", "query_error", err, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, deployRequestRow{
			OrganizationName: org,
			DatabaseName:     dbName,
			DeployRequest:    i,
		})
	}
	return nil, nil
}

func getDeployRequest(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_deploy_request.getDeployRequest", "connection_error", err)
		return nil, err
	}
	org := organization(ctx, d)
	dbName := d.KeyColumnQuals["database_name"].GetStringValue()
	num := d.KeyColumnQuals["number"].GetInt64Value()
	opts := &planetscale.GetDeployRequestRequest{Organization: org, Database: dbName, Number: uint64(num)}
	item, err := conn.DeployRequests.Get(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_deploy_request.getDeployRequest", "query_error", err, "opts", opts)
		return nil, err
	}
	result := deployRequestRow{
		OrganizationName: org,
		DatabaseName:     dbName,
		DeployRequest:    item,
	}
	return result, nil
}
