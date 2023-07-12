# Table: trello_webhook

A webhook is a feature that allows you to receive real-time notifications and updates about events that occur on your Trello boards. It enables you to integrate Trello with other applications, services, or custom systems by sending HTTP requests to a specific URL of your choice whenever a specified event happens in Trello.

The `trello_webhook` table can be used to query information about **ANY** webhook, and **you must specify the token ID** in the where clause (`where id_token=`).

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
  trello_webhook
where 
  id_token='1234ace54605094aa59b02c4b';
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
  active
  and id_token='1234ace54605094aa59b02c4b';
```

### Get details of a particular webhook

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
  id='1234ace54605094aa59b02c4b';
```