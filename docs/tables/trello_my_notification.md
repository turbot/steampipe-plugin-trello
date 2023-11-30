---
title: "Steampipe Table: trello_my_notification - Query Trello Notifications using SQL"
description: "Allows users to query Trello Notifications, specifically providing details on notification type, date, member creator, and associated data."
---

# Table: trello_my_notification - Query Trello Notifications using SQL

Trello Notifications are a feature within the Trello platform that provides updates to users about activities related to their boards, cards, and other Trello entities. These notifications can be triggered by various actions such as card movements, comments, due dates, and more. Notifications play a critical role in keeping users informed about the status and updates of their tasks and projects in Trello.

## Table Usage Guide

The `trello_my_notification` table provides insights into Trello Notifications within the Trello platform. As a project manager or team member, explore notification-specific details through this table, including the type of notification, date it was created, member who created it, and associated data. Utilize it to track activities, stay updated on project progress, and manage your tasks more effectively.

## Examples

### Basic info
Discover the segments that contain unread notifications and their associated dates and types in your Trello account. This can help you pinpoint specific areas that require your attention, enabling you to manage your workload more effectively.

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
Discover the segments that contain all your unread notifications. This can help you manage your workload by focusing on the tasks that require your immediate attention.

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
Determine the notifications initiated by a specific member. This can help understand the member's activity and engagement level on the platform.

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
Discover the segments that have received notifications in the past week. This can help you keep track of recent activities and understand the areas where action might be needed.

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