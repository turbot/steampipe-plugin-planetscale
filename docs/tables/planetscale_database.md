---
title: "Steampipe Table: planetscale_database - Query PlanetScale Databases using SQL"
description: "Allows users to query PlanetScale Databases, specifically their configurations, branches, and connection strings, providing insights into the database's operational status and deployment details."
---

# Table: planetscale_database - Query PlanetScale Databases using SQL

PlanetScale is a database as a service that provides a scalable, transactional, and serverless relational database built on MySQL and Kubernetes. It allows for easy database management and operations, including creating, deploying, and scaling databases. It is designed to handle heavy traffic, large datasets, and to provide high availability and strong consistency.

## Table Usage Guide

The `planetscale_database` table provides insights into databases within PlanetScale. As a database administrator or a DevOps engineer, explore database-specific details through this table, including configurations, branches, and connection strings. Utilize it to uncover information about databases, such as their operational status, deployment details, and to manage database versions and configurations.

## Examples

### List all databases
Explore all the databases available in your system to understand their structure and organization. This could help in managing data more efficiently and making informed decisions for data manipulation or migration.

```sql+postgres
select
  *
from
  planetscale_database;
```

```sql+sqlite
select
  *
from
  planetscale_database;
```

### Databases created in the last week
Discover the databases that have been established within the past week. This can be beneficial for monitoring recent activity and assessing the growth of your data infrastructure.

```sql+postgres
select
  *
from
  planetscale_database
where
  age(created_at) < interval '7 days';
```

```sql+sqlite
select
  *
from
  planetscale_database
where
  julianday('now') - julianday(created_at) < 7;
```

### Databases in US regions
Explore which databases are located in US regions to optimize data management and improve latency. This can be beneficial in enhancing user experience and ensuring efficient data access.

```sql+postgres
select
  *
from
  planetscale_database
where
  region like 'us-%';
```

```sql+sqlite
select
  *
from
  planetscale_database
where
  region like 'us-%';
```