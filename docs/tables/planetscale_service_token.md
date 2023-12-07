---
title: "Steampipe Table: planetscale_service_token - Query PlanetScale Service Tokens using SQL"
description: "Allows users to query PlanetScale Service Tokens, specifically the token ID and associated service, providing insights into service access and permissions."
---

# Table: planetscale_service_token - Query PlanetScale Service Tokens using SQL

PlanetScale Service Token is a resource in the PlanetScale Database-as-a-Service platform that allows you to authenticate your services and applications to access your databases. It provides a secure way to manage and control access to your PlanetScale databases, ensuring that only authenticated services can interact with your data. PlanetScale Service Token is essential for maintaining the security and integrity of your databases.

## Table Usage Guide

The `planetscale_service_token` table provides insights into service tokens within PlanetScale. As a database administrator, explore token-specific details through this table, including token ID, associated service, and creation timestamp. Utilize it to uncover information about service tokens, such as those associated with specific services, the lifespan of tokens, and the verification of service access.

## Examples

### List all service tokens
Explore the range of service tokens available in your system. This can be particularly useful when you need to manage or audit your service tokens.

```sql+postgres
select
  *
from
  planetscale_service_token;
```

```sql+sqlite
select
  *
from
  planetscale_service_token;
```