# Table: planetscale_database_branch

List database branches in your account.

## Examples

### List all branches for all databases

```sql
select
  *
from
  planetscale_database_branch
```

### List branches for a specific database

```sql
select
  *
from
  planetscale_database_branch
where
  database_name = 'test'
```

### Branches derived from the main branch

```sql
select
  *
from
  planetscale_database_branch
where
  database_name = 'test'
  and parent_branch = 'main'
```
