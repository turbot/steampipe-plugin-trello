---
title: "Steampipe Table: trello_list - Query Trello Lists using SQL"
description: "Allows users to query Trello Lists, specifically enabling the extraction of detailed information about lists in a Trello board."
---

# Table: trello_list - Query Trello Lists using SQL

Trello is a collaboration tool that organizes your projects into boards. A Trello board is a list of lists, filled with cards, used by you and your team. The Trello List resource in this context represents a list within a specific Trello board.

## Table Usage Guide

The `trello_list` table provides insights into lists within Trello boards. As a project manager or team leader, you can use this table to explore list-specific details, including card count, list status (open or closed), and associated metadata. This table is useful for tracking the progress of tasks and managing workflow across different projects or teams.

## Examples

### Basic info
Explore which Trello lists are active or closed within a specific board. This can help manage workflow and track project progress.

```sql
select
  id,
  name,
  id_board,
  closed,
  pos
from
  trello_list
where
  id_board = '12330ad5e3b81053d7d5315b';
```

### Get lists in a specific board which are closed
Explore which lists within a specific project board have been closed. This is particularly useful to track project progress and identify any tasks that are no longer active.

```sql
select
  id,
  name,
  id_board,
  closed,
  pos
from
  trello_list
where
  id_board = '12330ad5e3b81053d7d5315b'
  and closed;
```

### Get lists in a board which have been subscribed to
Discover the segments that have been subscribed to within a specific board. This is useful in identifying the areas of interest or focus for a particular team or project.

```sql
select
  id,
  name,
  id_board,
  closed,
  pos
from
  trello_list
where
  id_board = '12330ad5e3b81053d7d5315b'
  and subscribed;
```

### Get total lists in each board
Explore which boards have the most lists to better manage your project workflows and resource allocation. This can help in identifying where most activity is concentrated and aid in balancing the workload.

```sql
select
  id_board,
  count(*) as total_lists
from
  trello_list l,
  trello_board b
where
  l.id_board = b.id
group by
  id_board;
```