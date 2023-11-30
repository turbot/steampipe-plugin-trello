---
title: "Steampipe Table: trello_organization - Query Trello Organizations using SQL"
description: "Allows users to query Trello Organizations, specifically providing details on the organization's name, description, and associated boards."
---

# Table: trello_organization - Query Trello Organizations using SQL

Trello Organizations represent a team or group within Trello. They are used to group together boards and users to create a shared workspace. They provide a way to manage permissions across several boards at once.

## Table Usage Guide

The `trello_organization` table provides insights into Trello Organizations. As a project manager or team leader, explore organization-specific details through this table, including names, descriptions, and associated boards. Utilize it to uncover information about organizations, such as their membership, associated boards, and the level of visibility of each board within the organization.

## Examples

### Basic info
Explore basic information about a specific Trello organization, such as its name, description, and website. This can be useful for understanding the organization's online presence and identity.

```sql
select
  id,
  name,
  description,
  display_name,
  url,
  website
from
  trello_organization
where
  id = '123ace0f581f4de8a0dc184c';
```

### List the members assigned to a particular organization
Explore which users are linked to a specific organization. This can be useful for assessing membership or participation within certain organizational structures.

```sql
select
  m.id as member_id,
  username,
  o.id as organization_id,
  o.name as organization_name
from
  trello_member m,
  jsonb_array_elements_text(m.id_organizations) ido,
  trello_organization o
where
  o.id = ido
  and o.id = '123ace0f581f4de8a0dc184c';
```

### List details of the board associated to a particular organization
Explore the characteristics of a specific organization's board. This query is useful for gaining insights into the board's status, description, and accessibility, which can aid in organizational management and planning.

```sql
select
  b.id,
  b.name,
  b.description,
  b.id_organization,
  b.closed,
  b.url
from
  trello_board as b,
  trello_organization as o
where
  b.id_organization = o.id
  and b.id_organization = '1234ce0f581f4de8a0dc184c';
```