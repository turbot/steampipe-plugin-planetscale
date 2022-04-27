package planetscale

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type planetscaleConfig struct {
	Token        *string `cty:"token"`
	Organization *string `cty:"organization"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"organization": {
		Type: schema.TypeString,
	},
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
