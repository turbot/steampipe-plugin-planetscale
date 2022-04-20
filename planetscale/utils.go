package planetscale

import (
	"context"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/planetscale/planetscale-go/planetscale"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*planetscale.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "planetscale"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*planetscale.Client), nil
	}

	// Default to the env var settings
	token := os.Getenv("PLANETSCALE_TOKEN")
	organization := ""

	// Prefer config settings
	planetscaleConfig := GetConfig(d.Connection)
	if &planetscaleConfig != nil {
		if planetscaleConfig.Token != nil {
			token = *planetscaleConfig.Token
		}
		if planetscaleConfig.Organization != nil {
			organization = *planetscaleConfig.Organization
		}
	}

	// Error if the minimum config is not set
	if token == "" {
		return nil, errors.New("token must be configured")
	}
	if organization == "" {
		return nil, errors.New("organization must be configured")
	}

	conn, err := planetscale.NewClient(planetscale.WithAccessToken(token))
	if err != nil {
		return nil, err
	}

	/*
		// Service tokens don't seem to be supported in planetscale yet?
		// See https://github.com/planetscale/planetscale-go/issues/105
		conn, err := planetscale.NewClient(
			planetscale.WithServiceToken("nw-test", token),
		)
		if err != nil {
			return nil, err
		}
	*/

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func organization(_ context.Context, d *plugin.QueryData) string {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "planetscale_organization"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(string)
	}

	// Default to the env var settings
	org := ""

	// Prefer config settings
	planetscaleConfig := GetConfig(d.Connection)
	if &planetscaleConfig != nil {
		if planetscaleConfig.Organization != nil {
			org = *planetscaleConfig.Organization
		}
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, org)

	return org
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Not Found")
}
