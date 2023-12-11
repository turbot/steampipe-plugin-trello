---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/trello.svg"
brand_color: "#217EF9"
display_name: "Trello"
short_name: "trello"
description: "Steampipe plugin to query boards, cards, lists, and more from Trello."
og_description: "Query Trello with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/trello-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Trello + Steampipe

[Trello](https://trello.com) is a web-based, kanban-style, list-making application.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List details of the boards associated with your Trello account:

```sql
select
  id,
  name,
  id_organization,
  closed,
  url
from 
  trello_board;
```

```
+--------------------------+------------------------------------+--------------------------+--------+------------------------------------------------------------------+
| id                       | name                               | id_organization          | closed | url                                                              |
+--------------------------+------------------------------------+--------------------------+--------+------------------------------------------------------------------+
| 123ace54605094aa59b02c42 | Trello Agile Sprint Board Template | 649ace0f581f4de8a0dc184c | true   | https://trello.com/b/21wGVYiR/trello-agile-sprint-board-template |
+--------------------------+------------------------------------+--------------------------+--------+------------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/trello/tables)**

## Quick start

### Install

Download and install the latest Trello plugin:

```bash
steampipe plugin install trello
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Trello requires an [API key](https://developer.atlassian.com/cloud/trello/guides/rest-api/authorization/) and a [Token](https://trello.com/1/token) for all requests.                                                               |
| Permissions | API keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                             |
| Radius      | Each connection represents a single Trello Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/trello.spc`)<br />2. Credentials specified in environment variables, e.g., `TRELLO_API_KEY` and `TRELLO_TOKEN`.                     |

### Configuration

Installing the latest trello plugin will create a config file (`~/.steampipe/config/trello.spc`) with a single connection named `trello`:

Configure your account details in `~/.steampipe/config/trello.spc`:

```hcl
connection "trello" {
  plugin = "trello"

  # Trello API key for requests. Required.
  # See instructions at https://developer.atlassian.com/cloud/trello/guides/rest-api/authorization/
  # This can also be set via the `TRELLO_API_KEY` environment variable.
  # api_key = "a25ad2e37570117c0bad72d0a711ba5af"

  # Trello token for requests. Required.
  # This can also be set via the `TRELLO_TOKEN` environment variable.
  # token = "ATTAb179ea3c211722b0ebb2d223e1922b5e1ab1d28a3caac8d3722a83e9f91f25b973FDCC07"  
}
```

Alternatively, you can also use the standard Trello environment variables to obtain credentials **only if other arguments (`api_key` and `token`) are not specified** in the connection:

```sh
export TRELLO_API_KEY=a25ad2e37570117c0bad72d0a711ba5af
export TRELLO_TOKEN=ATTAb179ea3c211722b0ebb2d223e1922b5e1ab1d28a3caac8d3722a83e9f91f25b973FDCC07
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-trello
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
