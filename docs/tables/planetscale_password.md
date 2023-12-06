---
title: "Steampipe Table: planetscale_password - Query PlanetScale Passwords using SQL"
description: "Allows users to query PlanetScale Passwords. The table provides details about the passwords associated with PlanetScale databases, allowing for a comprehensive understanding of their configurations."
---

# Table: planetscale_password - Query PlanetScale Passwords using SQL

PlanetScale is a database platform that allows developers to build applications on MySQL-compatible databases. It provides a scalable, resilient, and secure database service suitable for mission-critical applications. The PlanetScale Password resource represents the passwords associated with PlanetScale databases.

## Table Usage Guide

The `planetscale_password` table allows users to query and analyze the passwords associated with PlanetScale databases. As a Database Administrator or Security Analyst, you can use this table to gain insights into the configuration and status of these passwords. This can assist in identifying potential security vulnerabilities, ensuring compliance with password policies, and maintaining overall database security.

**Important Notes**
- You must specify the `database_name` in the `where` clause to query this table.

## Examples

### List all passwords for a database
Identify all passwords associated with a specific database to enhance security monitoring and ensure proper access control. This is particularly useful in managing user permissions and maintaining database integrity.

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
Explore which passwords are associated with specific organizations, databases, and branches. This can be beneficial for managing and reviewing access control in a real-world scenario.

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
Explore which passwords in your organization's database are more than 90 days old. This can be crucial for maintaining security standards, as it allows you to identify and update potentially vulnerable or outdated passwords.

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