# Table: planetscale_audit_log

Query audit log records for API operations on your account.

## Examples

### Most recent actions for the account

```sql
select
  created_at,
  audit_action,
  actor_display_name,
  auditable_display_name,
  target_display_name
from
  planetscale_audit_log
order by
  created_at desc
limit
  10
```

### Find all branch creation events

```sql
select
  created_at,
  actor_display_name,
  auditable_display_name,
  target_display_name
from
  planetscale_audit_log
where
  audit_action = 'branch.created'
order by
  created_at desc
```

### Find events targeting a specific database

```sql
select
  created_at,
  audit_action,
  actor_display_name,
  auditable_display_name,
  target_display_name
from
  planetscale_audit_log
where
  target_type = 'Database'
  and target_display_name = 'my_db'
order by
  created_at desc
```

### Most common actions

```sql
select
  audit_action,
  count(*)
from
  planetscale_audit_log
group by
  audit_action
order by
  count desc
```

### Most common actors

```sql
select
  actor_id,
  actor_display_name,
  count(*)
from
  planetscale_audit_log
group by
  actor_id,
  actor_display_name
order by
  count desc
```
