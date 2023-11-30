---
title: "Steampipe Table: trello_search_board - Query Trello Boards using SQL"
description: "Allows users to query Trello Boards, specifically those that match a search query, providing insights into board details and facilitating board management."
---

# Table: trello_search_board - Query Trello Boards using SQL

Trello is a web-based, Kanban-style, list-making application. It's a subsidiary of Atlassian, used for project management and task organization. Trello Boards are the primary components of the Trello system where cards (tasks) are created, organized, and managed.

## Table Usage Guide

The `trello_search_board` table provides insights into Trello Boards that match a specific search query. As a project manager or team lead, explore board-specific details through this table, including board names, descriptions, and associated metadata. Utilize it to uncover information about boards, such as those related to specific projects, teams, or tasks, thereby facilitating efficient board management and task organization.

## Examples

### List all boards with the word "test" in the name
Discover the segments that include the term "test" in their names, allowing for targeted analysis or management of these specific areas within your organization. This is useful for quickly identifying and focusing on testing-related tasks or projects.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_search_board
where
  query = 'name:test';
```

### List all boards that are starred
Explore which of your Trello boards have been marked as important by using the 'starred' feature. This helps prioritize tasks and focus on key projects.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_search_board
where
  query = 'is:starred';
```

### List all boards that are closed
Discover the segments that include all boards which are no longer active. This could be beneficial for understanding the lifecycle of boards and identifying patterns or reasons for closure.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_search_board
where
  query = '-is:open';
```

### List all boards that are open and have the word "test" in the name
Discover the open boards that include 'test' in their name, useful for identifying specific project or task boards in a larger Trello organization.

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_search_board
where
  query = 'is:open name:test';
```