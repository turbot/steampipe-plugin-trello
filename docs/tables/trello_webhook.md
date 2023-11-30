---
title: "Steampipe Table: trello_webhook - Query Trello Webhooks using SQL"
description: "Allows users to query Trello Webhooks, specifically the event notifications that occur in a Trello Board. This provides insights into board activities and potential tracking of tasks."
---

# Table: trello_webhook - Query Trello Webhooks using SQL

Trello Webhooks are used to receive immediate updates for events that occur in a Trello Board. They are triggered by events, such as creating a card, updating a list, and other board activities. Webhooks provide a way for applications to get real-time information, making them essential for interactive applications.

## Table Usage Guide

The `trello_webhook` table provides insights into Webhooks within Trello. As a project manager or team lead, explore webhook-specific details through this table, including associated board, list, and card activities. Utilize it to uncover information about events, such as card creation, list updates, and other board activities, helping you track task progress and team activities in real-time.

## Examples

### Basic info
Explore active webhooks in Trello by identifying their unique IDs and descriptions. This is particularly useful for understanding the context and purpose of each webhook, especially when managing multiple webhooks tied to a specific token.

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
Explore which webhooks are currently active in your Trello account. This can help in managing and monitoring the operations triggered by these webhooks.

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
Analyze the settings of a specific webhook to understand its active status, purpose, and associated model. This is particularly useful when you need to troubleshoot or verify the functionality of a particular webhook.

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