package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePlanetScalePassword(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_password",
		Description: "Database Branches in the PlanetScale account.",
		List: &plugin.ListConfig{
			ParentHydrate: listDatabaseBranch,
			Hydrate:       listPassword,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "database_name"},
				{Name: "branch_name", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "database_name"},
				{Name: "branch_name"},
				{Name: "name"},
			},
			Hydrate: getPassword,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "Name of the database."},
			{Name: "branch_name", Type: proto.ColumnType_STRING, Description: "Name of the database branch."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Password.Name"), Description: "Name of the password."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Password.PublicID"), Description: "ID of the password."},
			{Name: "role", Type: proto.ColumnType_STRING, Transform: transform.FromField("Password.Role"), Description: "Role for the password."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Password.CreatedAt"), Description: "When the password was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Password.DeletedAt"), Description: "When the password was deleted."},
			// Not available - {Name: "plain_text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Password.PlainText"), Description: "Plain text of the password."},
			{Name: "connection_strings", Type: proto.ColumnType_JSON, Transform: transform.FromField("Password.ConnectionStrings"), Description: "Connection strings for the branch."},
		},
	}
}

type passwordRow struct {
	OrganizationName string
	DatabaseName     string
	BranchName       string
	Password         *planetscale.DatabaseBranchPassword
}

func listPassword(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listPassword", "connection_error", err)
		return nil, err
	}

	branch := h.Item.(databaseBranchRow)

	var orgName, dbName, branchName string
	if h.Item != nil {
		branch = h.Item.(databaseBranchRow)
		orgName = branch.OrganizationName
		dbName = branch.DatabaseName
		branchName = branch.Branch.Name
	} else {
		orgName = organization(ctx, d)
		dbName = d.KeyColumnQuals["database_name"].GetStringValue()
		branchName = d.KeyColumnQuals["branch_name"].GetStringValue()
	}

	// list all databases for the given organization
	opts := &planetscale.ListDatabaseBranchPasswordRequest{Organization: orgName, Database: dbName, Branch: branchName}
	items, err := conn.Passwords.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listPassword", "query_error", err, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, passwordRow{
			OrganizationName: orgName,
			DatabaseName:     dbName,
			BranchName:       branchName,
			Password:         i,
		})
	}
	return nil, nil
}

func getPassword(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getPassword", "connection_error", err)
		return nil, err
	}
	org := organization(ctx, d)
	dbName := d.KeyColumnQuals["database_name"].GetStringValue()
	branchName := d.KeyColumnQuals["branch_name"].GetStringValue()
	name := d.KeyColumnQuals["name"].GetStringValue()
	opts := &planetscale.GetDatabaseBranchPasswordRequest{Organization: org, Database: dbName, Branch: branchName, DisplayName: name}
	item, err := conn.Passwords.Get(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getPassword", "query_error", err, "opts", opts)
		return nil, err
	}
	result := passwordRow{
		OrganizationName: org,
		DatabaseName:     dbName,
		BranchName:       branchName,
		Password:         item,
	}
	return result, nil
}
