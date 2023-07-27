![image](https://hub.steampipe.io/images/plugins/turbot/planetscale-social-graphic.png)

# PlanetScale Plugin for Steampipe

Use SQL to query databases, logs and more from PlanetScale.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/planetscale)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/planetscale/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-planetscale/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install planetscale
```

Configure the server address in `~/.steampipe/config/planetscale.spc`:

```hcl
connection "planetscale" {
  plugin       = "planetscale"
  organization = "my_org"
  token        = "pscale_oauth_FWdKCeYK6sYQeJhNPTHRf3Ew_EXAMPLE"
}
```

Run steampipe:

```shell
steampipe query
```

Query your databases:

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-planetscale.git
cd steampipe-plugin-planetscale
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/planetscale.spc
```

Try it!

```
steampipe query
> .inspect planetscale
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-planetscale/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [PlanetScale Plugin](https://github.com/turbot/steampipe-plugin-planetscale/labels/help%20wanted)
