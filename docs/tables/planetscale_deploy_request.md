---
title: "Steampipe Table: planetscale_deploy_request - Query PlanetScale Deploy Requests using SQL"
description: "Allows users to query PlanetScale Deploy Requests, specifically the status and details of deployment requests, providing insights into deployment progress and potential issues."
---

# Table: planetscale_deploy_request - Query PlanetScale Deploy Requests using SQL

PlanetScale Deploy Request is a feature within PlanetScale that allows you to manage and monitor the deployment of changes to your databases. It provides a centralized way to track and manage deployment requests for various database resources. PlanetScale Deploy Request helps you stay informed about the status and progress of your deployments and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `planetscale_deploy_request` table provides insights into deployment requests within PlanetScale. As a DevOps engineer, explore deployment-specific details through this table, including status, progress, and associated metadata. Utilize it to uncover information about deployments, such as those with pending status, the progress of each deployment, and the verification of deployment details.

## Examples

### List all deploy requests for all databases
Explore which deployment requests have been made for all databases, which can help you keep track of changes and manage resources more effectively.

```sql+postgres
select
  *
from
  planetscale_deploy_request;
```

```sql+sqlite
select
  *
from
  planetscale_deploy_request;
```

### List deploy requests for a specific database
Discover the segments that have requested deployment for a particular database. This can be useful to track and manage deployment requests, ensuring that they are handled appropriately.

```sql+postgres
select
  *
from
  planetscale_deploy_request
where
  database_name = 'test';
```

```sql+sqlite
select
  *
from
  planetscale_deploy_request
where
  database_name = 'test';
```

### List open deploy requests for all databases
Explore which deploy requests are currently open across all databases. This is useful for identifying ongoing deployments and managing resource allocation.

```sql+postgres
select
  *
from
  planetscale_deploy_request
where
  state = 'open';
```

```sql+sqlite
select
  *
from
  planetscale_deploy_request
where
  state = 'open';
```

### List deploy requests with no changes
Explore the instances where deployment requests have been made, but no changes were made. This can help in identifying and reducing unnecessary deployment requests, thereby streamlining the deployment process.

```sql+postgres
select
  *
from
  planetscale_deploy_request
where
  deployment ->> 'state' = 'no_changes';
```

```sql+sqlite
select
  *
from
  planetscale_deploy_request
where
  json_extract(deployment, '$.state') = 'no_changes';
```