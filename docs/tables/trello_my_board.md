---
title: "Steampipe Table: trello_my_board - Query Trello Boards using SQL"
description: "Allows users to query Trello Boards, specifically for personal boards, providing insights into board details and activity."
---

# Table: trello_my_board - Query Trello Boards using SQL

Trello is a web-based, Kanban-style, list-making application which is a subsidiary of Atlassian. Originally made by Fog Creek Software in 2011, it was spun out to form the basis of a separate company in 2014 and later sold to Atlassian in January 2017. The service operates via cloud computing, with the system being operated using JavaScript, HTML and CSS.

## Table Usage Guide

The `trello_my_board` table provides insights into personal boards within Trello. As a project manager, explore board-specific details through this table, including board name, description, privacy settings, and associated metadata. Utilize it to uncover information about boards, such as their visibility settings, the members associated with each board, and the activity on each board.

## Examples

### Basic info
Explore the Trello boards you're a part of, including their names, descriptions, and associated organizations. This can help you manage your work more efficiently by giving you a comprehensive view of your tasks and projects.

```sql+postgres
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from 
  trello_my_board;
```

```sql+sqlite
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from 
  trello_my_board;
``` 

### List my open boards
Explore your active Trello boards to gain insights into their details and assess their organization affiliations. This allows for efficient board management by focusing only on boards that are currently open and active.

```sql+postgres
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  not closed;
```

```sql+sqlite
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  closed is not 1;
```

### List my pinned boards
Discover the details of your preferred boards on Trello. This query allows you to quickly identify and access your favorite boards, enhancing your ability to manage and prioritize tasks.

```sql+postgres
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  pinned;
```

```sql+sqlite
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  pinned = 1;
```

### List my starred boards
Explore your favorite Trello boards, gaining insights into their names, descriptions, associated organizations, and URLs. This is useful for quickly accessing and managing your preferred projects.

```sql+postgres
select
  id,
  name,
  description,
  id_organization,
  closed,
  url 
from
  trello_my_board
where
  starred;
```

```sql+sqlite
select
  id,
  name,
  description,
  id_organization,
  closed,
  url 
from
  trello_my_board
where
  starred = 1;
```

### List my boards in a specific organization
Explore the boards that you have in a particular organization, helping you to manage and track your projects effectively. This can be particularly useful when you are part of multiple organizations and need to quickly identify all your boards within a specific one.

```sql+postgres
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  id_organization = '1234ce0f581f4de8a0dc184c';
```

```sql+sqlite
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  id_organization = '1234ce0f581f4de8a0dc184c';
```

### List my boards with a specific label
Explore which of your Trello boards have been tagged with a 'Blocked' label. This can help you quickly identify and prioritize tasks that are facing obstacles.

```sql+postgres
select
  tmb.id,
  tmb.name,
  id_organization,
  l.key as label_key,
  l.value as label_value
from
  trello_my_board as tmb,
  jsonb_each_text(label_names) l
where
  l.value = 'Blocked';
```

```sql+sqlite
select
  tmb.id,
  tmb.name,
  id_organization,
  l.key as label_key,
  l.value as label_value
from
  trello_my_board as tmb,
  json_each(label_names) as l
where
  l.value = 'Blocked';
```

### List preferences for each board
Discover the individual preferences for each board you have on Trello. This can be useful to understand and optimize your board settings for better project management.

```sql+postgres
select
  tmb.id,
  tmb.name,
  id_organization,
  p.key as pref_key,
  p.value as pref_value
from
  trello_my_board as tmb,
  jsonb_each_text(prefs) p;
```

```sql+sqlite
select
  tmb.id,
  tmb.name,
  id_organization,
  p.key as pref_key,
  p.value as pref_value
from
  trello_my_board as tmb,
  json_each(prefs) p;
```

### List custom fields for each board
Explore the customization options of each board by identifying the unique fields associated with them. This analysis can assist in understanding the various customization settings and configurations applied to each board, providing insights into their unique characteristics and usage.

```sql+postgres
select
  tmb.id,
  tmb.name,
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
  trello_my_board as tmb,
  jsonb_array_elements(custom_fields) c;
```

```sql+sqlite
select
  tmb.id,
  tmb.name,
  json_extract(c.value, '$.ID') as custom_field_id,
  json_extract(c.value, '$.Name') as custom_field_name,
  json_extract(c.value, '$.IDModel') as custom_field_model_id,
  json_extract(c.value, '$.IDModelType') as custom_field_model_type_id,
  json_extract(c.value, '$.FieldGroup') as custom_field_group,
  json_extract(c.value, '$.Pos') as custom_field_pos,
  json_extract(c.value, '$.Type') as custom_field_type,
  json_extract(c.value, '$.Display') as custom_field_display,
  json_extract(c.value, '$.Options') as custom_field_options
from
  trello_my_board as tmb,
  json_each(custom_fields) as c;
```