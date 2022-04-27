package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tablePlanetScaleDatabaseBranch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_database_branch",
		Description: "Database Branches in the PlanetScale account.",
		List: &plugin.ListConfig{
			ParentHydrate: listDatabase,
			Hydrate:       listDatabaseBranch,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "database_name", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "database_name"},
				{Name: "name"},
			},
			Hydrate: getDatabaseBranch,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "Name of the database."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Branch.Name"), Description: "Name of the branch."},
			{Name: "parent_branch", Type: proto.ColumnType_STRING, Transform: transform.FromField("Branch.Name"), Description: "Parent of this branch."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Branch.Region.Slug"), Description: "Region where the database is located."},
			{Name: "ready", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Branch.Ready"), Description: "True if the branch is ready."},
			{Name: "production", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Branch.Production"), Description: "True if this branch is in production."},
			{Name: "access_host_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Branch.AccessHostURL"), Description: "Host name to access the database."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Branch.CreatedAt"), Description: "When the branch was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Branch.UpdatedAt"), Description: "When the branch was updated."},
		},
	}
}

type databaseBranchRow struct {
	OrganizationName string
	DatabaseName     string
	Branch           *planetscale.DatabaseBranch
}

func listDatabaseBranch(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listDatabaseBranch", "connection_error", err)
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
	opts := &planetscale.ListDatabaseBranchesRequest{Organization: org, Database: dbName}
	items, err := conn.DatabaseBranches.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listDatabaseBranch", "query_error", err, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, databaseBranchRow{
			OrganizationName: org,
			DatabaseName:     dbName,
			Branch:           i,
		})
	}
	return nil, nil
}

func getDatabaseBranch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getDatabaseBranch", "connection_error", err)
		return nil, err
	}
	org := organization(ctx, d)
	dbName := d.KeyColumnQuals["database_name"].GetStringValue()
	name := d.KeyColumnQuals["name"].GetStringValue()
	opts := &planetscale.GetDatabaseBranchRequest{Organization: org, Database: dbName, Branch: name}
	item, err := conn.DatabaseBranches.Get(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getDatabaseBranch", "query_error", err, "opts", opts)
		return nil, err
	}
	result := databaseBranchRow{
		OrganizationName: org,
		DatabaseName:     dbName,
		Branch:           item,
	}
	return result, nil
}
