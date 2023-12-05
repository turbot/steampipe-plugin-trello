---
title: "Steampipe Table: trello_member - Query Trello Members using SQL"
description: "Allows users to query Trello Members, providing insights into member details, including their ID, username, full name, and avatar hash."
---

# Table: trello_member - Query Trello Members using SQL

Trello is a web-based, Kanban-style, list-making application which is a subsidiary of Atlassian. Members in Trello are users who have access to and can perform certain actions within a board. They can be part of one or more boards and can have different permissions and roles depending on the board they are part of.

## Table Usage Guide

The `trello_member` table provides insights into the members within Trello. As a project manager or team leader, you can explore member-specific details through this table, including their ID, username, full name, and avatar hash. Utilize it to uncover information about members, such as their roles in different boards, their activity, and their contribution to various tasks.

## Examples

### Basic info
Explore which Trello members are registered in your system, including their usernames and contact details. This can help manage user access and maintain up-to-date records.

```sql+postgres
select
  id,
  username,
  full_name,
  email,
  initials
from
  trello_member;
```

```sql+sqlite
select
  id,
  username,
  full_name,
  email,
  initials
from
  trello_member;
```

### List members registered under a given email address
Determine the areas in which members are registered under a specific email address. This can be particularly useful in identifying duplicate accounts or tracking user activity across different platforms.

```sql+postgres
select
  id,
  username,
  full_name,
  email,
  initials
from
  trello_member
where
  email = 'abc@gmal.com';
```

```sql+sqlite
select
  id,
  username,
  full_name,
  email,
  initials
from
  trello_member
where
  email = 'abc@gmal.com';
```

### List the boards that each member has access to
Explore which Trello boards each member has access to, providing a comprehensive view of individual access rights across different projects. This can help in analyzing the distribution of work and ensuring appropriate access control.

```sql+postgres
select
  m.id as member_id,
  username,
  b.id as board_id,
  b.name as board_name
from
  trello_member m,
  jsonb_array_elements_text(m.id_boards) idb,
  trello_board b
where
  b.id = idb;
```

```sql+sqlite
select
  m.id as member_id,
  username,
  b.id as board_id,
  b.name as board_name
from
  trello_member m,
  json_each(m.id_boards) as idb,
  trello_board b
where
  b.id = idb.value;
```

### List the organizations that each member have been assigned to
Explore which organizations each member is affiliated with, to gain insights into team structures and collaborations. This can be useful in identifying potential overlaps or gaps in team assignments.

```sql+postgres
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
  o.id = ido;
```

```sql+sqlite
select
  m.id as member_id,
  username,
  o.id as organization_id,
  o.name as organization_name
from
  trello_member m,
  json_each(m.id_organizations) as ido,
  trello_organization o
where
  o.id = ido.value;
```