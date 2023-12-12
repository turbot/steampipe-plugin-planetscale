---
title: "Steampipe Table: planetscale_certificate - Query PlanetScale Certificates using SQL"
description: "Allows users to query PlanetScale Certificates, specifically the details of SSL certificates used in PlanetScale database service."
---

# Table: planetscale_certificate - Query PlanetScale Certificates using SQL

PlanetScale Certificates are SSL certificates used within the PlanetScale database service for encrypting connections between the client and the server. This ensures the secure transmission of sensitive data over networks. The certificates are automatically managed and renewed by PlanetScale, providing a seamless and secure database experience.

## Table Usage Guide

The `planetscale_certificate` table provides insights into SSL certificates used within the PlanetScale database service. As a Database Administrator, explore certificate-specific details through this table, including certificate state, creation and expiration dates, and associated metadata. Utilize it to uncover information about certificates, such as their current state, the duration of validity, and the verification of expiration dates.

**Important Notes**
- You must specify the `database_name` in the `where` clause to query this table.

## Examples

### List all certificates for a database
Explore all certificates related to a particular database to ensure proper security management and compliance. This can be particularly useful in maintaining data integrity and safeguarding sensitive information.

```sql+postgres
select
  *
from
  planetscale_certificate
where
  database_name = 'test';
```

```sql+sqlite
select
  *
from
  planetscale_certificate
where
  database_name = 'test';
```

### List all certificates for all branches
Explore the certificates associated with each branch of your database. This can help in ensuring that all branches are secure and properly certified, aiding in maintaining the overall security posture of your database.

```sql+postgres
select
  c.*
from
  planetscale_database as d
join
  planetscale_certificate as c on d.name = c.database_name;
```

```sql+sqlite
select
  c.*
from
  planetscale_database as d
join
  planetscale_certificate as c on d.name = c.database_name;
```