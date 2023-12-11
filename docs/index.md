---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/planetscale.svg"
brand_color: "#8467F3"
display_name: "PlanetScale"
short_name: "planetscale"
description: "Steampipe plugin to query databases, logs and more from PlanetScale."
og_description: "Query PlanetScale with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/planetscale-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# PlanetScale + Steampipe

[PlanetScale](https://planetscale.com) is a MySQL-compatible serverless database platform.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List databases in your PlanetScale account:

```sql
select
  name,
  region_slug,
  created_at
from
  planetscale_database;
```

```
+------+-------------+---------------------------+
| name | region_slug | created_at                |
+------+-------------+---------------------------+
| test | us-east     | 2021-11-16T22:31:03-05:00 |
| prod | us-west     | 2022-02-11T14:03:24-05:00 |
+------+-------------+---------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/planetscale/tables)**

## Get started

### Install

Download and install the latest PlanetScale plugin:

```bash
steampipe plugin install planetscale
```

### Configuration

Installing the latest planetscale plugin will create a config file (`~/.steampipe/config/planetscale.spc`) with a single connection named `planetscale`:

```hcl
connection "planetscale" {
  plugin       = "planetscale"

  organization = "my_org"

  # Required: Set your access token
  # To get this token:
  # 1. Install the pscale CLI
  # 2. Login to the CLI
  # 3. cat ~/.config/planetscale/access-token
  token = "pscale_oauth_FWdKCeYK6sYQeJhNPTHRf3Ew_EXAMPLE"
}
```

- `organization` - Organization to scope all queries to.
- `token` - Access token (note: NOT a service token) from PlanetScale.

Environment variables are also available as an alternate configuration method:
* `PLANETSCALE_ORGANIZATION`
* `PLANETSCALE_TOKEN`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-planetscale
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
