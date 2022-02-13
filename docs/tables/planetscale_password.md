# Table: planetscale_password

List passwords for database branches.

Note:
* `database_name` must be specified in the `where` clause for all queries.

## Examples

### List all passwords for a database

```sql
select
  p.organization_name,
  p.database_name,
  p.branch_name,
  p.name,
  p.created_at
from
  planetscale_password
where
  database_name = 'test'
```

### List all passwords for all databases & branches

```sql
select
  p.organization_name,
  p.database_name,
  p.branch_name,
  p.name,
  p.created_at
from
  planetscale_database as d
join
  planetscale_password as p on d.name = p.database_name
```

### List all passwords more than 90 days old

```sql
select
  p.organization_name,
  p.database_name,
  p.branch_name,
  p.name,
  p.created_at
from
  planetscale_database as d
join
  planetscale_password as p on d.name = p.database_name
where
  age(p.created_at) > interval '90 days'
```
