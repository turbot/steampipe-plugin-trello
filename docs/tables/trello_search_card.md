---
title: "Steampipe Table: trello_search_card - Query Trello Cards using SQL"
description: "Allows users to query Trello Cards, specifically by searching for specific card data, providing insights into card details and activity."
---

# Table: trello_search_card - Query Trello Cards using SQL

Trello is a web-based, Kanban-style, list-making application, and a subsidiary of Atlassian. Users can create their task boards with several columns and move the tasks between them. Typically columns include task statuses: To Do, In Progress, Done. The tool can be used for personal and business purposes including real estate management, software project management, school bulletin boards, lesson planning, and law office case management.

## Table Usage Guide

The `trello_search_card` table provides insights into Trello Cards within the Trello platform. As a project manager or team member, explore card-specific details through this table, including card names, descriptions, due dates, and associated lists. Utilize it to uncover information about cards, such as their current list, members assigned, and the status of their completion.

**Important Notes**
- You must always include at least one search term in the where or join clause using the `query` column. You can narrow the results using the search qualifiers in any combination. See [Trello search](https://trello.com/search) for details on the Trello query syntax.

## Examples

### Get all cards assigned to a specific member
Identify all task cards assigned to a specific team member. This is useful for tracking individual workloads and progress in project management scenarios.

```sql+postgres
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
  trello_search_card
where
  query='@username';
```

```sql+sqlite
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
  trello_search_card
where
  query='@username';
```

### Get cards without any members assigned
Identify instances where task cards have no members assigned. This could be useful to pinpoint areas that may need additional resources or oversight in a project management context.

```sql+postgres
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
  trello_search_card
where
  query='-has:members';
```

```sql+sqlite
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
  trello_search_card
where
  query='-has:members';
```

### Get all cards from a specific board
Determine the areas in which specific board cards are used, including their status and due dates. This is useful for tracking task progress and managing project timelines effectively.

```sql+postgres
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
  trello_search_card
where
  query='board:123ace54605094aa59b02c4b';
```

```sql+sqlite
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
  trello_search_card
where
  query='board:123ace54605094aa59b02c4b';
```

### Get all cards created within a month
Discover the segments that have been recently added within the past month. This can be useful to track the progress of tasks or projects in real-time and maintain an updated workflow.  

```sql+postgres
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
  trello_search_card
where
  query='created:month';
```

```sql+sqlite
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
  trello_search_card
where
  query='created:month';
```

### Get all cards created within 14 days
Explore which tasks have been initiated within the last two weeks. This is useful for keeping track of recent project developments and understanding current workload.

```sql+postgres
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
  trello_search_card
where
  query='created:14';
```

```sql+sqlite
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
  trello_search_card
where
  query='created:14';
```

### Get all cards due within a week
Discover the tasks that are due within the upcoming week. This can help you prioritize your work and manage your time effectively.

```sql+postgres
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
  trello_search_card
where
  query='due:week';
```

```sql+sqlite
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
  trello_search_card
where
  query='due:week';
```

### Get all cards due within a day
Discover the segments that have tasks due within a day. This is beneficial for managing time-sensitive projects and ensuring all tasks are completed on schedule.

```sql+postgres
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
  trello_search_card
where
  query='due:day';
```

```sql+sqlite
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
  trello_search_card
where
  query='due:day';
```

### Get all cards with attachments
Discover the segments that have attachments within a project management tool. This is particularly useful for identifying tasks that contain additional information or resources, aiding in effective project tracking and management.  

```sql+postgres
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
  trello_search_card
where
  query='has:attachments';
```

```sql+sqlite
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
  trello_search_card
where
  query='has:attachments';
```

### Get all cards with a specific label
Discover the segments that include all cards marked with a specific label. This allows for a focused overview of tasks that are blocked, facilitating effective project management and troubleshooting.

```sql+postgres
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
  trello_search_card
where
  query='label:Blocked';
```

```sql+sqlite
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
  trello_search_card
where
  query='label:Blocked';
```

### Get all cards with a specific label color
Explore which tasks have been tagged with a specific label color in a project management tool. This can help in quickly identifying and categorizing tasks based on their labels, enhancing project organization and management.

```sql+postgres
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
  trello_search_card 
where
  query='label:blue';
```

```sql+sqlite
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
  trello_search_card 
where
  query='label:blue';
```

### Get all cards from a specific list in a board
Explore which tasks are currently listed under the 'To Do' category for a specific project board. This is useful for project management, allowing you to quickly assess the status and details of all pending tasks in one place.

```sql+postgres
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
  trello_search_card 
where
  query='list:"To Do" board:123ace54605094aa59b02c4b';
```

```sql+sqlite
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
  trello_search_card 
where
  query='list:"To Do" board:123ace54605094aa59b02c4b';
```

### Get all cards that are overdue
Discover the segments that have overdue tasks in Trello. This query is useful in identifying and managing tasks that have passed their due date, assisting in project management and task prioritization.

```sql+postgres
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
  trello_search_card 
where
  query='due:overdue';
```

```sql+sqlite
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
  trello_search_card 
where
  query='due:overdue';
```

### Get all cards that are overdue and no members assigned
Explore which tasks are overdue and currently unassigned, to efficiently allocate resources and prioritize workload. This can help in identifying bottlenecks and improving project management.

```sql+postgres
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
  trello_search_card 
where
  query='due:overdue -has:members';
```

```sql+sqlite
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
  trello_search_card 
where
  query='due:overdue -has:members';
```