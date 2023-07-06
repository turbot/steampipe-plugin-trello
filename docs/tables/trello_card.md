# Table: trello_card

Cards are individual items or tasks within a list. They can represent specific tasks, ideas, or notes. Cards can be moved between lists to indicate progress or changes in status.

The `trello_card` table can be used to query information about **ANY** card, and **you must specify which list** in the where or join clause (`where id_list=`, `join trello_list l on l.id=`).

## Examples

### Basic info

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