# Table: planetscale_backup

List backups in your account.

Note:
* `database_name` must be specified in the `where` clause for all queries.

## Examples

### List all backups for a database

```sql
select
  *
from
  planetscale_backup
where
  database_name = 'test'
```

### Get a specific backup

```sql
select
  *
from
  planetscale_backup
where
  database_name = 'test'
  and branch_name = 'main'
  and name = '2022.02.12 04:11:03'
```

### List all backups for all branches

```sql
select
  b.*
from
  planetscale_database as d
join
  planetscale_backup as b on d.name = b.database_name
```

### Backups expiring in the next week

```sql
select
  b.database_name,
  b.branch_name,
  b.name,
  b.expires_at
from
  planetscale_database as d
join
  planetscale_backup as b on d.name = b.database_name
where
  b.expires_at < current_timestamp + interval '7 days'
order by
  b.expires_at
```
