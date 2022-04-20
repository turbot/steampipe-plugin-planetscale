package main

import (
	"github.com/turbot/steampipe-plugin-planetscale/planetscale"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: planetscale.Plugin})
}
