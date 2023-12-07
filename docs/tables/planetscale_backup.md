---
title: "Steampipe Table: planetscale_backup - Query PlanetScale Backup using SQL"
description: "Allows users to query PlanetScale Backups, specifically the data related to the backup of databases, providing insights into backup status, size, and related details."
---

# Table: planetscale_backup - Query PlanetScale Backup using SQL

PlanetScale Backup is a feature within PlanetScale that allows you to create, manage, and restore backups of your databases. It provides a reliable way to ensure the safety of your data, allowing you to recover your databases in the event of data loss. PlanetScale Backup helps you maintain the integrity and availability of your data by providing automated and manual backup options.

## Table Usage Guide

The `planetscale_backup` table provides insights into the backup details within PlanetScale. As a database administrator, explore backup-specific details through this table, including backup status, size, and related metadata. Utilize it to uncover information about backups, such as backup creation time, duration, and the associated database details.

**Important Notes**
- You must specify the `database_name` in the `where` clause to query this table.

## Examples

### List all backups for a database
Explore all backups related to a specific database to maintain data integrity and ensure business continuity in case of data loss or corruption. This is particularly useful in disaster recovery scenarios or when performing routine data audits.

```sql+postgres
select
  *
from
  planetscale_backup
where
  database_name = 'test';
```

```sql+sqlite
select
  *
from
  planetscale_backup
where
  database_name = 'test';
```

### Get a specific backup
Discover the details of a specific backup within your main branch of the 'test' database, allowing you to gain insights into that particular backup's status and details. This is useful for assessing the health and success of your backup operations.

```sql+postgres
select
  *
from
  planetscale_backup
where
  database_name = 'test'
  and branch_name = 'main'
  and name = '2022.02.12 04:11:03';
```

```sql+sqlite
select
  *
from
  planetscale_backup
where
  database_name = 'test'
  and branch_name = 'main'
  and name = '2022.02.12 04:11:03';
```

### List all backups for all branches
Discover the segments that consist of all the backups for each branch in your database system. This is useful for maintaining data integrity and ensuring you have access to all necessary data in case of system failure or data loss.

```sql+postgres
select
  b.*
from
  planetscale_database as d
join
  planetscale_backup as b on d.name = b.database_name;
```

```sql+sqlite
select
  b.*
from
  planetscale_database as d
join
  planetscale_backup as b on d.name = b.database_name;
```

### Backups expiring in the next week
Discover the segments that are set to expire within the next week. This is useful in proactive planning to prevent any unexpected data loss.

```sql+postgres
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
  b.expires_at;
```

```sql+sqlite
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
  b.expires_at < datetime('now', '+7 days')
order by
  b.expires_at;
```