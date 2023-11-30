---
title: "Steampipe Table: planetscale_region - Query PlanetScale Regions using SQL"
description: "Allows users to query PlanetScale Regions, providing details such as slug, display name, and enabled status."
---

# Table: planetscale_region - Query PlanetScale Regions using SQL

PlanetScale Regions are the geographical locations available for deploying databases in the PlanetScale service. Each region is identified by a unique slug and has a display name. The enabled status indicates whether the region is currently available for use.

## Table Usage Guide

The `planetscale_region` table provides insights into the regions available within PlanetScale. As a database administrator or developer, explore region-specific details through this table, including the unique slug, display name, and enabled status. Utilize it to understand the geographical distribution of your databases and plan for future deployments.

## Examples

### List all regions
Explore all the geographical regions in which PlanetScale operates, to better understand its global distribution and reach. This knowledge can be useful for planning resource allocation and optimizing service performance.

```sql
select
  *
from
  planetscale_region
```