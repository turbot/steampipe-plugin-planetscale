# Table: planetscale_certificate

List certificates in your account.

Note:
* `database_name` must be specified in the `where` clause for all queries.

## Examples

### List all certificates for a database

```sql
select
  *
from
  planetscale_certificate
where
  database_name = 'test'
```

### List all certificates for all branches

```sql
select
  c.*
from
  planetscale_database as d
join
  planetscale_certificate as c on d.name = c.database_name
```
