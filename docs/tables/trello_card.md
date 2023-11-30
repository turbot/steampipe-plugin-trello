---
title: "Steampipe Table: trello_card - Query Trello Cards using SQL"
description: "Allows users to query Cards in Trello, specifically the details of each card, providing insights into task details, labels, and member assignments."
---

# Table: trello_card - Query Trello Cards using SQL

A Trello Card is a fundamental unit in Trello, a popular project management tool. Cards represent tasks and contain information such as descriptions, comments, attached files, and checklists. They can be assigned to specific team members, labeled for quick identification, and moved across different lists representing various stages of a project.

## Table Usage Guide

The `trello_card` table provides insights into the tasks within Trello. As a project manager or team member, explore card-specific details through this table, including descriptions, comments, attached files, and checklists. Utilize it to uncover information about tasks, such as their current stage, assigned members, and associated labels.

## Examples

### Basic info
Determine the areas in which specific tasks are assigned by identifying the details of Trello cards within a particular list. This helps in understanding the task distribution and planning future tasks effectively.

```sql
select
  id,
  name,
  description,
  id_board,
  id_list,
  start,
  due,
  closed
from
  trello_card
where
  id_list='123ace54605094aa59b02c4b';
```

### Get all cards in a list that are past due
Explore which tasks in a specific project are overdue and still open. This is useful for tracking project progress and identifying areas that may require additional resources or attention.

```sql
select
  id,
  name,
  description,
  id_board,
  id_list,
  start,
  due,
  closed
from
  trello_card
where
  id_list='123ace54605094aa59b02c4b'
  and due < now()
  and not closed;
```

### Get all cards in a list that are due in a week
Determine the tasks in a specific project that are due in the upcoming week, to help prioritize work and manage deadlines effectively. This query is particularly useful for project managers keeping track of task deadlines and ensuring project progress.

```sql
select
  id,
  name,
  description,
  id_board,
  id_list,
  start,
  due,
  closed
from
  trello_card
where
  id_list='123ace54605094aa59b02c4b'
  and due < now() + interval '7 days'
  and not closed;
```

### Get all cards in a list in order of their due date, with the most recent due first
Identify tasks in a specific project that are still ongoing, arranged based on their deadlines starting from the most urgent. This helps in prioritizing work and ensuring timely completion of tasks.

```sql
select
  id,
  name,
  description,
  id_board,
  id_list,
  start,
  due,
  closed
from
  trello_card
where
  id_list='123ace54605094aa59b02c4b'
  and not closed
order by
  due desc;
```

### List all cards attachments
This query allows you to identify all attachments linked to a specific task in a project management tool. It provides valuable insights into the resources associated with a task, which can aid in task comprehension and completion.

```sql
select
  id,
  name,
  a ->> 'id' as attachment_id,
  a ->> 'name' as attachment_name,
  a ->> 'url' as attachment_url,
  a ->> 'date' as attachment_date,
  a ->> 'edgeColor' as attachment_edge_color,
  a ->> 'idMember' as attachment_id_member,
  a ->> 'mimeType' as attachment_mime_type,
  a ->> 'pos' as attachment_pos
from 
  trello_card c,
  jsonb_array_elements(c.attachments) as a
where
  id_list='123ace54605094aa59b02c4b';
```

### List all badges for a card
Discover the segments that have all the badges assigned to a specific card. This is beneficial in understanding the attributes or achievements associated with that card, providing insights into its importance or role within the project.

```sql
select
  id,
  name,
  b.key as badge_key,
  b.value as badge_value
from
  trello_card c,
  jsonb_each_text(badges) b
where
  id_list='123ace54605094aa59b02c4b';
```

### List all cards with a specific label
Discover the segments that have been marked as 'Blocked' within a specific Trello list. This can help in identifying bottlenecks or issues that are hindering progress in a project.

```sql
select
  id,
  name,
  l ->> 'id' as label_id,
  l ->> 'name' as label_name,
  l ->> 'color' as label_color
from
  trello_card c,
  jsonb_array_elements(labels) l
where
  id_list='123ace54605094aa59b02c4b'
  and l ->> 'name' = 'Blocked';
```

### List all the member details assigned to a card
Explore which members are associated with a specific card in your Trello board. This can help in understanding task delegation and tracking the progress of individual team members.

```sql
select
  c.id,
  c.name,
  m.id as member_id,
  m.username as member_username,
  m.full_name as member_full_name,
  m.initials as member_initials
from
  trello_card c,
  trello_member m,
  jsonb_array_elements_text(c.id_members) mid
where
  id_list='123ace54605094aa59b02c4b'
  and m.id = mid;
```

### List all open cards in a list
Explore which tasks are currently active on a specific Trello list. This is useful for monitoring project progress and identifying outstanding tasks.

```sql
select
  id,
  name,
  description,
  id_board,
  id_list,
  start,
  due,
  closed
from
  trello_card
where
  id_list='123ace54605094aa59b02c4b'
  and not closed;
```

### List all cards in a list that have been subscribed to
Explore which tasks within a specific task list you are actively following. This is useful for keeping track of important tasks and ensuring you don't miss any updates.

```sql
select
  id,
  name,
  description,
  id_board,
  id_list,
  start,
  due,
  closed
from
  trello_card
where
  id_list='123ace54605094aa59b02c4b'
  and subscribed;
```