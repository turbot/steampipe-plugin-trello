# Table: trello_my_board

A board represents a project or a high-level category. Within a board, you can create lists and cards to organize your work. This table is a view of all boards that you are a member of.

## Examples

### Basic info

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from 
  trello_my_board;
``` 

### List my open boards

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  not closed;
```

### List my pinned boards

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  pinned;
```

### List my starred boards

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url 
from
  trello_my_board
where
  starred;
```

### List my boards in a specific organization

```sql
select
  id,
  name,
  description,
  id_organization,
  closed,
  url
from
  trello_my_board
where
  id_organization = '1234ce0f581f4de8a0dc184c';
```

### List my boards with a specific label

```sql
select
  id,
  name,
  id_organization,
  l.key as label_key,
  l.value as label_value
from
  trello_my_board,
  jsonb_each_text(label_names) l
where
  l.value = 'Blocked';
```

### List preferences for each board

```sql
select
  id,
  name,
  id_organization,
  p.key as pref_key,
  p.value as pref_value
from
  trello_my_board,
  jsonb_each_text(prefs) p;
```

### List custom fields for each board

```sql
select
  id,
  name,
  c ->> 'ID' as custom_field_id,
  c ->> 'Name' as custom_field_name,
  c ->> 'IDModel' as custom_field_model_id,
  c ->> 'IDModelType' as custom_field_model_type_id,
  c ->> 'FieldGroup' as custom_field_group,
  c ->> 'Pos' as custom_field_pos,
  c ->> 'Type' as custom_field_type,
  c ->> 'Display' as custom_field_display,
  c ->> 'Options' as custom_field_options
from
  trello_my_board,
  jsonb_array_elements(custom_fields) c;
```