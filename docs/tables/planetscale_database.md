# Table: planetscale_database

List databases in your account.

## Examples

### List all databases

```sql
select
  *
from
  planetscale_database
```

### Databases created in the last week

```sql
select
  *
from
  planetscale_database
where
  age(created_at) < interval '7 days'
```

### Databases in US regions

```sql
select
  *
from
  planetscale_database
where
  region like 'us-%'
```
