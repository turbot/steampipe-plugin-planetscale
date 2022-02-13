# Table: planetscale_deploy_request

List deploy requests in your account.

## Examples

### List all deploy requests for all databases

```sql
select
  *
from
  planetscale_deploy_request
```

### List deploy requests for a specific database

```sql
select
  *
from
  planetscale_deploy_request
where
  database_name = 'test'
```

### List open deploy requests for all databases

```sql
select
  *
from
  planetscale_deploy_request
where
  state = 'open'
```

### List deploy requests with no changes

```sql
select
  *
from
  planetscale_deploy_request
where
  deployment ->> 'state' = 'no_changes'
```
