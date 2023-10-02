## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#17](https://github.com/turbot/steampipe-plugin-planetscale/pull/17))
- Recompiled plugin with Go version `1.21`. ([#17](https://github.com/turbot/steampipe-plugin-planetscale/pull/17))

## v0.3.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#12](https://github.com/turbot/steampipe-plugin-planetscale/pull/12))

## v0.2.0 [2022-09-26]

_Enhancements_

- Added support for `PLANETSCALE_ORGANIZATION` env variable.
- Added pagination support in `planetscale_audit_log` table.

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements.
- Recompiled plugin with Go version `1.19`.

## v0.1.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#5](https://github.com/turbot/steampipe-plugin-planetscale/pull/5))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`.([#4](https://github.com/turbot/steampipe-plugin-planetscale/pull/4))

## v0.0.1 [2022-02-16]

_What's new?_

- New tables added
  - [planetscale_audit_log](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_audit_log)
  - [planetscale_backup](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_backup)
  - [planetscale_certificate](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_certificate)
  - [planetscale_database](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_database)
  - [planetscale_database_branch](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_database_branch)
  - [planetscale_deploy_request](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_deploy_request)
  - [planetscale_organization](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_organization)
  - [planetscale_password](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_password)
  - [planetscale_region](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_region)
  - [planetscale_service_token](https://hub.steampipe.io/plugins/turbot/planetscale/tables/planetscale_service_token)
