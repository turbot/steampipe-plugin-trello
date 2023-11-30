---
title: "Steampipe Table: trello_my_organization - Query Trello Organizations using SQL"
description: "Allows users to query Trello Organizations, specifically providing details about organization's name, display name, description, website, and more."
---

# Table: trello_my_organization - Query Trello Organizations using SQL

Trello is a web-based, list-making application, originally made by Fog Creek Software in 2011. It operates via boards (which correspond to projects), where cards (correspond to tasks) can be created and moved between lists (which usually correspond to stages of a project). An organization in Trello is a way to group boards and people together to create a shared workspace.

## Table Usage Guide

The `trello_my_organization` table provides insights into Organizations within Trello. As a project manager or team lead, explore organization-specific details through this table, including name, display name, description, website, and more. Utilize it to uncover information about organizations, such as their members, associated boards, and overall organization structure.

## Examples

### Basic info
Explore your Trello organization's basic details such as identification, name, description, display name, website URL, etc. This can be useful for gaining a quick overview or for auditing purposes.

```sql
select
  id,
  name,
  description,
  display_name,
  url,
  website
from
  trello_my_organization;
```

### List the members assigned to my organization
Determine the members linked to your specific organization, in order to understand team composition and manage resources effectively.

```sql
select
  m.id as member_id,
  username,
  o.id as organization_id,
  o.name as organization_name
from
  trello_member m,
  jsonb_array_elements_text(m.id_organizations) ido,
  trello_my_organization o
where
  o.id = ido
  and o.id = '123ace0f581f4de8a0dc184c';
```

### List details of the board associated to my organization
Explore the specifics of an organization's board, such as its status and web address, to gain insights into its operational details and online presence. This is particularly useful for understanding the organization's current activities and accessibility.

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
  trello_my_organization as o
where
  b.id_organization = o.id
  and b.id_organization = '1234ce0f581f4de8a0dc184c';
```