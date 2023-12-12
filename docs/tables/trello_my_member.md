---
title: "Steampipe Table: trello_my_member - Query Trello Members using SQL"
description: "Allows users to query Trello Members, specifically providing details about the user profile, including full name, username, email, and avatar hash."
---

# Table: trello_my_member - Query Trello Members using SQL

Trello is a web-based, list-making application, originally made by Fog Creek Software in 2011. It operates via boards (which correspond to projects), lists (which correspond to task lists), and cards (which correspond to tasks). A Trello member refers to a user who has an account with Trello and can create and manage boards, lists, and cards.

## Table Usage Guide

The `trello_my_member` table provides insights into Trello members. As a project manager or team leader, explore member-specific details through this table, including full name, username, email, and avatar hash. Utilize it to uncover information about members, such as their activity on different boards and their interaction with various tasks.

## Examples

### Basic info
Explore which member details are available in your Trello account, such as user ID, username, full name, email, and initials. This can be useful in identifying and managing users, especially in a large team.

```sql+postgres
select
  id,
  username,
  full_name,
  email,
  initials
from
  trello_my_member;
```

```sql+sqlite
select
  id,
  username,
  full_name,
  email,
  initials
from
  trello_my_member;
```

### Get boards I have access to
Explore the various boards you have access to within Trello, allowing you to quickly identify and manage your permissions across different projects. This is particularly useful for understanding your involvement and access level across various team boards.

```sql+postgres
select
  m.id as member_id,
  username,
  b.id as board_id,
  b.name as board_name
from
  trello_my_member m,
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
  trello_my_member m,
  json_each(m.id_boards) as idb,
  trello_board b
where
  b.id = idb.value;
```

### Get organizations I have been assigned to
Explore which organizations you've been assigned to, providing a clear view of your involvement across different teams. This can help manage your tasks and collaborations more effectively.

```sql+postgres
select
  m.id as member_id,
  username,
  o.id as organization_id,
  o.name as organization_name
from
  trello_my_member m,
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
  trello_my_member m,
  json_each(m.id_organizations) as ido,
  trello_organization o
where
  o.id = ido.value;
```