package planetscale

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tablePlanetScaleAuditLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "planetscale_audit_log",
		Description: "AuditLogs in the PlanetScale account.",
		List: &plugin.ListConfig{
			Hydrate: listAuditLog,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of the log entry."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of log entry, e.g. AuditLogEvent."},
			{Name: "actor_id", Type: proto.ColumnType_STRING, Description: "Unique ID of the actor."},
			{Name: "actor_type", Type: proto.ColumnType_STRING, Description: "Type of the actor, e.g. User."},
			{Name: "actor_display_name", Type: proto.ColumnType_STRING, Description: "Display name of the actor."},
			{Name: "auditable_id", Type: proto.ColumnType_STRING, Description: "Unique ID for the resource type of the audit entry."},
			{Name: "auditable_type", Type: proto.ColumnType_STRING, Description: "Resource type the audit entry is for, e.g. Branch."},
			{Name: "auditable_display_name", Type: proto.ColumnType_STRING, Description: "Display name of the resource for this audit entry, e.g. test_branch."},
			{Name: "audit_action", Type: proto.ColumnType_STRING, Description: "Full action for this audit record, e.g. deploy_request.created."},
			{Name: "action", Type: proto.ColumnType_STRING, Description: "Short action for this audit record, e.g. created."},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "Geographic location the action was requested from."},
			{Name: "remote_ip", Type: proto.ColumnType_STRING, Description: "IP address the action was requested from."},
			{Name: "target_id", Type: proto.ColumnType_STRING, Description: "ID of the resource type for this audit record."},
			{Name: "target_type", Type: proto.ColumnType_STRING, Description: "Resource type for this audit record, e.g. Database."},
			{Name: "target_display_name", Type: proto.ColumnType_STRING, Description: "Display name for the target resoruce, e.g. test_db."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Metadata for the audit record."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the audit record was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the audit record was updated."},
		},
	}
}

func listAuditLog(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("planetscale_audit_log.listAuditLog", "connection_error", err)
		return nil, err
	}

	// list all audit logs for the given organization
	opts := &planetscale.ListAuditLogsRequest{Organization: organization(ctx, d)}

	startingAfter := ""
	for {
		// Use a limit of 1000 - seems to work, but I can't find docs for max
		resp, err := conn.AuditLogs.List(ctx, opts, planetscale.WithStartingAfter(startingAfter), planetscale.WithLimit(1000))
		if err != nil {
			plugin.Logger(ctx).Error("planetscale_audit_log.listAuditLog", "query_error", err, "opts", opts, "startingAfter", startingAfter)
			return nil, err
		}
		for _, i := range resp.Data {
			d.StreamListItem(ctx, i)
		}
		if resp.HasNext {
			startingAfter = *resp.CursorEnd
		} else {
			break
		}
	}

	return nil, nil
}
