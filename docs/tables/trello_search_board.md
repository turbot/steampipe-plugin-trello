# Table: trello_search_board

A board represents a project or a high-level category. Within a board, you can create lists and cards to organize your work.

**You must always include at least one search term** in the where or join clause using the `query` column. You can narrow the results using the search qualifiers in any combination. See [Trello search](https://trello.com/search) for details on the Trello query syntax.

## Examples

### List all boards with the word "test" in the name

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

