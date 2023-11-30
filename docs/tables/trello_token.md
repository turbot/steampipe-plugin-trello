---
title: "Steampipe Table: trello_token - Query Trello Tokens using SQL"
description: "Allows users to query Trello Tokens, specifically providing insights into token details such as ID, name, and associated member ID."
---

# Table: trello_token - Query Trello Tokens using SQL

Trello Tokens are unique identifiers that are generated for each user session. These tokens are used to authenticate Trello API requests and are tied to the user's Trello account, providing access to their boards, lists, and cards. It is crucial to manage and monitor these tokens to ensure the security and integrity of the user's data.

## Table Usage Guide

The `trello_token` table provides insights into Trello Tokens within Trello API. As a developer or security analyst, explore token-specific details through this table, including token ID, name, and associated member ID. Utilize it to uncover information about tokens, such as those associated with specific users, the permissions granted by each token, and the verification of token validity.

## Examples

### Basic info
This query enables you to analyze the creation and expiration dates for a specific member's token in Trello. It can be used to monitor token validity and ensure timely updates, enhancing security and functionality.

```sql
select
  id,
  id_member,
  identifier,
  date_created,
  date_expires
from
  trello_token
where 
  id_member='1234ace54605094aa59b02c4b';
```

### List all active tokens
Explore which tokens are currently active and when they were created. This is useful for managing and monitoring token usage and validity in real-time, particularly for the specified member.

```sql
select
  id,
  id_member,
  identifier,
  date_created,
  date_expires
from
  trello_token
where
  date_expires > now()
  and id_member='1234ace54605094aa59b02c4b';
```

### Get details of a particular token
Analyze the settings to understand the specifics of a particular access token, such as its creation date and expiration date. This can be useful for security auditing or troubleshooting access issues.

```sql
select
  id,
  id_member,
  identifier,
  date_created,
  date_expires
from
  trello_token
where
  id='1234ace54605094aa59b02c4b';
```

### Get all tokens of a particular member
Discover the segments that represent all the tokens associated with a specific member. This can be useful to understand the member's activity history or to analyze token usage patterns over time.

```sql
select
  id,
  id_member,
  identifier,
  date_created,
  date_expires
from
  trello_token
where
  id_member='1234ace54605094aa59b02c4b';
```

### Get all tokens that are expiring in a week
Determine the areas in which Trello tokens are set to expire within the upcoming week for a specific member. This allows for proactive management of token renewals, ensuring continuous access and avoiding potential disruptions.

```sql
select
  id,
  id_member,
  identifier,
  date_created,
  date_expires
from
  trello_token
where
  date_expires < now() + interval '7 days'
  and id_member='1234ace54605094aa59b02c4b';
```

### Get permission details of each token
Explore which permissions are associated with each token to better understand the access level and control of each token. This could be useful in assessing the security and access management of your system.

```sql
select
  id,
  identifier,
  p ->> 'idModel' as id_model,
  p ->> 'modelType' as model_type,
  p ->> 'read' as read,
  p ->> 'write' as write
from
  trello_token,
  jsonb_array_elements(permissions) p
where
  id_member='1234ace54605094aa59b02c4b';
```

### Get all tokens that have write access to a particular board
Identify instances where specific user tokens have write access to a particular board. This is useful to manage and monitor user permissions, ensuring only authorized users can modify the board.

```sql
select
  id,
  id_member,
  identifier,
  date_created,
  date_expires
from
  trello_token,
  jsonb_array_elements(permissions) p
where
  id_member = '1234ace54605094aa59b02c4b'
  and p ->> 'idModel' = 'a1asq1244605094aa59b02c4b'
  and p ->> 'modelType' = 'Board'
  and p ->> 'write' = 'true';
```

### Get all tokens that have write access to all the boards
Explore which tokens have the authority to modify all boards. This is useful for identifying potential security risks, ensuring only authorized users have such access.

```sql
select
  id,
  id_member,
  identifier,
  date_created,
  date_expires
from
  trello_token,
  jsonb_array_elements(permissions) p
where
  id_member = '1234ace54605094aa59b02c4b'
  and p ->> 'idModel' = '*'
  and p ->> 'modelType' = 'Board'
  and p ->> 'write' = 'true';
```