# Table: trello_my_member

A member is an individual user who has access to a Trello board. Members can be added to boards to collaborate on projects, tasks, or any other work being managed within Trello. This table is a view of the your user details.

## Examples

### Basic info

```sql
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

```sql
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

### Get organizations I have been assigned to

```sql
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