# Table: trello_organization

An organization is a higher-level entity that helps users manage multiple boards and collaborate with larger teams or groups. Organizations in Trello provide a way to group boards and members together, making it easier to organize and share work among team members.

The `trello_organization` table can be used to query information about **ANY** organization, and **you must specify which organization** in the where clause (`where id=`).

## Examples

### Basic info

```sql
select
  id,
  name,
  description,
  display_name,
  url,
  website
from
  trello_organization
where
  id = '123ace0f581f4de8a0dc184c';
```

### List the members assigned to a particular organization

```sql
select
  m.id as member_id,
  username,
  o.id as organization_id,
  o.name as organization_name
from
  trello_member m,
  jsonb_array_elements_text(m.id_organizations) ido,
  trello_organization o
where
  o.id = ido
  and o.id = '123ace0f581f4de8a0dc184c';
```

### List details of the board associated to a particular organization

```sql
select
  b.id,
  b.name,
  b.description,
  b.id_organization,
  b.closed,
  b.url
from
  trello_board as b,
  trello_organization as o
where
  b.id_organization = o.id
  and b.id_organization = '1234ce0f581f4de8a0dc184c';
```