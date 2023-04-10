package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tablePlanetScaleBackup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_backup",
		Description: "Database Branches in the PlanetScale account.",
		List: &plugin.ListConfig{
			ParentHydrate: listDatabaseBranch,
			Hydrate:       listBackup,
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
			Hydrate: getBackup,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "Name of the database."},
			{Name: "branch_name", Type: proto.ColumnType_STRING, Description: "Name of the database branch."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Backup.Name"), Description: "Name of the backup."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Backup.PublicID"), Description: "ID of the backup."},
			{Name: "state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Backup.State"), Description: "State of the backup."},
			{Name: "size", Type: proto.ColumnType_INT, Transform: transform.FromField("Backup.Size"), Description: "Size of the backup."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Backup.CreatedAt"), Description: "When the backup was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Backup.UpdatedAt"), Description: "When the backup was updated."},
			{Name: "started_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Backup.StartedAt"), Description: "When the backup was started."},
			{Name: "completed_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Backup.CompletedAt"), Description: "When the backup was completed."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Backup.ExpiresAt"), Description: "When the backup expires."},
		},
	}
}

type backupRow struct {
	OrganizationName string
	DatabaseName     string
	BranchName       string
	Backup           *planetscale.Backup
}

func listBackup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listBackup", "connection_error", err)
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
		dbName = d.EqualsQuals["database_name"].GetStringValue()
		branchName = d.EqualsQuals["branch_name"].GetStringValue()
	}

	// list all databases for the given organization
	opts := &planetscale.ListBackupsRequest{Organization: orgName, Database: dbName, Branch: branchName}
	items, err := conn.Backups.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listBackup", "query_error", err, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, backupRow{
			OrganizationName: orgName,
			DatabaseName:     dbName,
			BranchName:       branchName,
			Backup:           i,
		})
	}
	return nil, nil
}

func getBackup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getBackup", "connection_error", err)
		return nil, err
	}
	org := organization(ctx, d)
	dbName := d.EqualsQuals["database_name"].GetStringValue()
	branchName := d.EqualsQuals["branch_name"].GetStringValue()
	name := d.EqualsQuals["name"].GetStringValue()
	opts := &planetscale.GetBackupRequest{Organization: org, Database: dbName, Branch: branchName, Backup: name}
	item, err := conn.Backups.Get(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getBackup", "query_error", err, "opts", opts)
		return nil, err
	}
	result := backupRow{
		OrganizationName: org,
		DatabaseName:     dbName,
		BranchName:       branchName,
		Backup:           item,
	}
	return result, nil
}
