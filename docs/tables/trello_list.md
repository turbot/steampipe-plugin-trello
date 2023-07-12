# Table: trello_list

Lists are columns on a board that represent different stages of a project or workflow.

The `trello_list` table can be used to query information about **ANY** list, and **you must specify the board ID** in the where or join clause (`where id_board=`, `join trello_board b on b.id=`).

## Examples

### Basic info

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
