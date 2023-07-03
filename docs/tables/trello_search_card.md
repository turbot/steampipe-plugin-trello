# Table: trello_search_card

Cards are individual items or tasks within a list. They can represent specific tasks, ideas, or notes. Cards can be moved between lists to indicate progress or changes in status.

**You must always include at least one search term when searching** in the where or join clause using the `query` column. You can narrow the results using the search qualifiers in any combination. See [Trello search](https://trello.com/search) for details on the Trello query syntax.

## Examples

### Get all cards assigned to a specific member

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
  trello_search_card
where
  query='@username';
```

### Get cards without any members assigned

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
  trello_search_card
where
  query='-has:members';
```

### Get all cards from a specific board

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
  trello_search_card
where
  query='board:123ace54605094aa59b02c4b';
```

### Get all cards created within a month
  
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
  trello_search_card
where
  query='created:month';
```

### Get all cards created within 14 days

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
  trello_search_card
where
  query='created:14';
```

### Get all cards due within a week

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
  trello_search_card
where
  query='due:week';
```

### Get all cards due within a day

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
  trello_search_card
where
  query='due:day';
```

### Get all cards with attachments
  
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
  trello_search_card
where
  query='has:attachments';
```

### Get all cards with a specific label

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
  trello_search_card
where
  query='label:Blocked';
```

### Get all cards with a specific label color

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
  trello_search_card 
where
  query='label:blue';
```

### Get all cards from a specific list in a board

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
  trello_search_card 
where
  query='list:"To Do" board:123ace54605094aa59b02c4b';
```

### Get all cards that are overdue

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
  trello_search_card 
where
  query='due:overdue';
```

### Get all cards that are overdue and no members assigned

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
  trello_search_card 
where
  query='due:overdue -has:members';
```