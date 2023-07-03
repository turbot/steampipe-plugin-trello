# Table: trello_webhook

A webhook is a feature that allows you to receive real-time notifications and updates about events that occur on your Trello boards. It enables you to integrate Trello with other applications, services, or custom systems by sending HTTP requests to a specific URL of your choice whenever a specified event happens in Trello.

The `trello_webhook` table can be used to query information about **ANY** webhook, and **you must specify which webhook** in the where clause (`where id=`).

## Examples

### Basic info

```sql
select
  id,
  active,
  description,
  callback_url,
  id_model
from
  trello_webhook;
```

### List all active webhooks

```sql
select
  id,
  active,
  description,
  callback_url,
  id_model
from 
  trello_webhook
where
  active;
```