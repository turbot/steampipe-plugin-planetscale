package planetscale

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type planetscaleConfig struct {
	Token        *string `hcl:"token"`
	Organization *string `hcl:"organization"`
}

func ConfigInstance() interface{} {
	return &planetscaleConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) planetscaleConfig {
	if connection == nil || connection.Config == nil {
		return planetscaleConfig{}
	}
	config, _ := connection.Config.(planetscaleConfig)
	return config
}
