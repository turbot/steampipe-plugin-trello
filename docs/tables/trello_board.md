---
title: "Steampipe Table: trello_board - Query Trello Boards using SQL"
description: "Allows users to query Trello Boards, specifically the detailed information about each board such as its name, description, privacy settings, and more."
---

# Table: trello_board - Query Trello Boards using SQL

Trello is a web-based, Kanban-style, list-making application which is a subsidiary of Atlassian. It allows users to create, manage, and organize boards, which represent projects, and within boards, cards, which represent tasks. Trello Boards are the basic units of a Trello workflow and can be used to represent a project, a team, or even a company.

## Table Usage Guide

The `trello_board` table provides insights into Trello Boards within the Trello application. As a project manager or team leader, explore board-specific details through this table, including names, descriptions, privacy settings, and more. Utilize it to uncover information about boards, such as those that are public, the ones that are private, and the verification of privacy settings.

## Examples

### Basic info
Explore the basic details of your Trello boards, including determining if they are closed or open, to better manage your organization's workflow and tasks. This can help you assess each board's status and understand its role within your organization.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_board;
```

### List all boards that are open
Discover the segments that are still active within your organization. This query is particularly useful for tracking ongoing projects and tasks, offering a clear view of all open boards.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_board
where
  not closed;
```

### List all boards that are pinned
Explore which boards have been marked as important or frequently used by identifying those that have been pinned. This can aid in quickly accessing key information and prioritizing tasks.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_board
where
  pinned;
```

### List all boards that are starred
Discover the segments that are marked as important or highlighted, allowing you to focus on priority areas or tasks within your project management tool.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_board
where
  starred;
```

### List all boards in a specific organization
Discover the segments that belong to a specific organization by examining all the boards associated with it. This can aid in understanding the structure and operations of the organization.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_board
where
  id_organization = '1234ce0f581f4de8a0dc184c';
```

### List the details of a specific board
Explore a specific board's details, such as its name, status, and associated organization, to gain insights into its usage and relevance. This can be beneficial in understanding the board's context and role within your Trello organization.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_board
where
  id = '1234ce0f581f4de8a0dc184c';
```

### List all boards with a specific label
Explore which Trello boards have been marked with a 'Blocked' label to understand areas of your projects that might be facing obstacles. This can help identify potential bottlenecks and improve project management efficiency.

```sql
select
  id,
  name,
  id_organization,
  l.key as label_key,
  l.value as label_value
from
  trello_board,
  jsonb_each_text(label_names) l
where
  l.value = 'Blocked';
```

### List preferences of each board
Explore the various preferences set for each board in an organization. This can help in understanding the customization and settings applied to different boards, aiding in effective management and control.

```sql
select
  id,
  name,
  id_organization,
  p.key as pref_key,
  p.value as pref_value
from
  trello_board,
  jsonb_each_text(prefs) p;
```

### List all subscribed boards
Explore which Trello boards you're subscribed to, helping you manage and stay updated with the boards that matter to you. This is beneficial in scenarios where you need to quickly identify your active subscriptions amongst numerous boards.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_board
where
  subscribed;
```

### List the custom fields of each board
Explore the customization of different boards by analyzing the unique fields associated with each. This can be useful in understanding the organization and categorization of information within each board.

```sql
select
  id,
  name,
  c ->> 'ID' as custom_field_id,
  c ->> 'Name' as custom_field_name,
  c ->> 'IDModel' as custom_field_model_id,
  c ->> 'IDModelType' as custom_field_model_type_id,
  c ->> 'FieldGroup' as custom_field_group,
  c ->> 'Pos' as custom_field_pos,
  c ->> 'Type' as custom_field_type,
  c ->> 'Display' as custom_field_display,
  c ->> 'Options' as custom_field_options
from
  trello_board,
  jsonb_array_elements(custom_fields) c;
```