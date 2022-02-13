package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePlanetScaleCertificate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_certificate",
		Description: "Database Branches in the PlanetScale account.",
		List: &plugin.ListConfig{
			ParentHydrate: listDatabaseBranch,
			Hydrate:       listCertificate,
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
			Hydrate: getCertificate,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "Name of the database."},
			{Name: "branch_name", Type: proto.ColumnType_STRING, Description: "Name of the database branch."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Certificate.Name"), Description: "Name of the certificate."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Certificate.PublicID"), Description: "ID of the certificate."},
			{Name: "role", Type: proto.ColumnType_STRING, Transform: transform.FromField("Certificate.Role"), Description: "Role for the certificate."},
			{Name: "certificate", Type: proto.ColumnType_STRING, Transform: transform.FromField("Certificate.Certificate"), Description: "Certificate string."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Certificate.CreatedAt"), Description: "When the certificate was created."},
			{Name: "deleted_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Certificate.DeletedAt"), Description: "When the certificate was deleted."},
		},
	}
}

type certificateRow struct {
	OrganizationName string
	DatabaseName     string
	BranchName       string
	Certificate      *planetscale.DatabaseBranchCertificate
}

func listCertificate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listCertificate", "connection_error", err)
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
	opts := &planetscale.ListDatabaseBranchCertificateRequest{Organization: orgName, Database: dbName, Branch: branchName}
	items, err := conn.Certificates.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.listCertificate", "query_error", err, "opts", opts)
		return nil, err
	}

	for _, i := range items {
		d.StreamListItem(ctx, certificateRow{
			OrganizationName: orgName,
			DatabaseName:     dbName,
			BranchName:       branchName,
			Certificate:      i,
		})
	}
	return nil, nil
}

func getCertificate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getCertificate", "connection_error", err)
		return nil, err
	}
	org := organization(ctx, d)
	dbName := d.KeyColumnQuals["database_name"].GetStringValue()
	branchName := d.KeyColumnQuals["branch_name"].GetStringValue()
	name := d.KeyColumnQuals["name"].GetStringValue()
	opts := &planetscale.GetDatabaseBranchCertificateRequest{Organization: org, Database: dbName, Branch: branchName, DisplayName: name}
	item, err := conn.Certificates.Get(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_database_branch.getCertificate", "query_error", err, "opts", opts)
		return nil, err
	}
	result := certificateRow{
		OrganizationName: org,
		DatabaseName:     dbName,
		BranchName:       branchName,
		Certificate:      item,
	}
	return result, nil
}
