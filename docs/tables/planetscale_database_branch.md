---
title: "Steampipe Table: planetscale_database_branch - Query PlanetScale Database Branches using SQL"
description: "Allows users to query PlanetScale Database Branches, specifically providing detailed insights into each database branch's configuration and status."
---

# Table: planetscale_database_branch - Query PlanetScale Database Branches using SQL

PlanetScale is a database platform that allows you to create, scale, and manage databases for your applications. A PlanetScale Database Branch is a version of your database that you can modify without affecting the main database. It is used for testing changes, debugging issues, and developing new features without risking the stability of your production database.

## Table Usage Guide

The `planetscale_database_branch` table provides insights into the branches of databases within PlanetScale. As a database administrator or developer, explore branch-specific details through this table, including the branch's configuration, status, and associated metadata. Utilize it to uncover information about branches, such as their creation time, update time, and the readiness status of the branch.

## Examples

### List all branches for all databases
Explore the various branches across all your databases to manage and keep track of different versions of your data effectively. This aids in making informed decisions for version control and data integrity.

```sql+postgres
select
  *
from
  planetscale_database_branch;
```

```sql+sqlite
select
  *
from
  planetscale_database_branch;
```

### List branches for a specific database
Explore the different branches within a specific database to understand its various versions or instances. This is useful for managing changes and updates to the database.

```sql+postgres
select
  *
from
  planetscale_database_branch
where
  database_name = 'test';
```

```sql+sqlite
select
  *
from
  planetscale_database_branch
where
  database_name = 'test';
```

### Branches derived from the main branch
Discover the segments that have been created from the main branch in a specific database, which can be useful when tracking the evolution of your data and managing your database structure.

```sql+postgres
select
  *
from
  planetscale_database_branch
where
  database_name = 'test'
  and parent_branch = 'main';
```

```sql+sqlite
select
  *
from
  planetscale_database_branch
where
  database_name = 'test'
  and parent_branch = 'main';
```