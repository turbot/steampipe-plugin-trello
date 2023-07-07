# Table: trello_my_notification

When a member is added to a board, they can receive notifications about board activity, such as updates to cards, mentions, or changes made by other members. Notifications help keep everyone informed and engaged in the project.

## Examples

### Basic info

```sql
select
  id,
  date,
  unread,
  type,
  data,
  id_member_creator
from
  trello_my_notification;
```

### List all unread notifications

```sql
select
  id,
  date,
  unread,
  type,
  data,
  id_member_creator
from
  trello_my_notification
where
  unread;
```

### List all notifications created by a particular member

```sql
select
  id,
  date,
  unread,
  type,
  data,
  id_member_creator
from
  trello_my_notification
where
  id_member_creator = '34458739284892367890';
```

### List all notifications from last week

```sql
select
  id,
  date,
  unread,
  type,
  data,
  id_member_creator
from
  trello_my_notification
where
  date > now() - interval '1 week';
```