package planetscale

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-planetscale",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"planetscale_audit_log":       tablePlanetScaleAuditLog(ctx),
			"planetscale_backup":          tablePlanetScaleBackup(ctx),
			"planetscale_certificate":     tablePlanetScaleCertificate(ctx),
			"planetscale_database":        tablePlanetScaleDatabase(ctx),
			"planetscale_database_branch": tablePlanetScaleDatabaseBranch(ctx),
			"planetscale_deploy_request":  tablePlanetScaleDeployRequest(ctx),
			"planetscale_organization":    tablePlanetScaleOrganization(ctx),
			"planetscale_password":        tablePlanetScalePassword(ctx),
			"planetscale_region":          tablePlanetScaleRegion(ctx),
			"planetscale_service_token":   tablePlanetScaleServiceToken(ctx),
		},
	}
	return p
}
