package planetscale

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
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
	if connection == nil || connection.GetConfig() == nil {
		return planetscaleConfig{}
	}
	config, _ := connection.GetConfig().(planetscaleConfig)
	return config
}
