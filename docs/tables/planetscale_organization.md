---
title: "Steampipe Table: planetscale_organization - Query PlanetScale Organizations using SQL"
description: "Allows users to query PlanetScale Organizations, providing insights into the configuration, status, and metadata associated with each organization."
---

# Table: planetscale_organization - Query PlanetScale Organizations using SQL

PlanetScale is a database platform that simplifies the process of scaling databases. It allows for the creation and management of organizations, which serve as a way to group related databases and database branches. Each organization in PlanetScale has its own configuration, status, and metadata.

## Table Usage Guide

The `planetscale_organization` table provides insights into organizations within PlanetScale. As a database administrator or developer, you can explore organization-specific details through this table, including its configuration, status, and associated metadata. Utilize it to manage and monitor your organizations, to ensure optimal organization structure and database performance.

## Examples

### List all organizations
Explore all the organizations within your database, helping you to understand the breadth and diversity of data you're dealing with. This can be particularly useful for large-scale projects or when managing databases across multiple departments or teams.

```sql
select
  *
from
  planetscale_organization
```