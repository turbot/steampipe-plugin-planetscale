---
title: "Steampipe Table: planetscale_audit_log - Query PlanetScale Audit Logs using SQL"
description: "Allows users to query Audit Logs in PlanetScale, specifically the logs of actions performed in the database, providing insights into database activities and potential anomalies."
---

# Table: planetscale_audit_log - Query PlanetScale Audit Logs using SQL

PlanetScale's Audit Log is a feature that records all actions performed in your database. It provides a comprehensive view of the activities and changes within your database, including who made the changes, when they were made, and what exactly was changed. This feature is crucial for maintaining the security and integrity of your database.

## Table Usage Guide

The `planetscale_audit_log` table provides insights into actions performed within your PlanetScale database. As a database administrator, explore details through this table, including the actor, the action performed, and the timestamp of the action. Utilize it to uncover information about database activities, such as who made changes, when they were made, and what exactly was changed.

## Examples

### Most recent actions for the account
Analyze the most recent activities in your account to understand and track changes made by users. This can be useful for auditing purposes, to ensure accountability and transparency in account operations.

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
Explore the historical timeline of branch creation events within your database. This can help you understand project progression, user activity, and potential areas of focus.

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
Explore events specifically aimed at a particular database to gain insights into actions performed, who performed them, and when they occurred. This can be particularly useful for auditing purposes, allowing you to track any changes or activities related to your database.

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
Determine the most frequent actions in your system's audit log to understand user behavior patterns and system usage. This can help in optimizing resources and identifying potential security issues.

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
Explore which actors are most frequently logged in the audit log to understand system usage patterns and identify potential security concerns.

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