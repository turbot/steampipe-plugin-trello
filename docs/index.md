---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/trello.svg"
brand_color: "#FFE01B"
display_name: "Trello"
short_name: "trello"
description: "Steampipe plugin to query audiences, automation workflows, campaigns, and more from Trello."
og_description: "Query Trello with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/trello-social-graphic.png"
---

# Trello + Steampipe

[Trello](https://trello.com) is a marketing automation and email marketing platform.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List devices which block incoming connections in your Trello tailnet:

```sql
select
  id,
  title,
  content_type,
  create_time,
  emails_sent,
  send_time,
  status,
  type
from
  trello_campaign;
```

```
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
| id         | title                              | content_type | create_time               | emails_sent | send_time | status | type             |
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
| f739729f66 | We're here to help you get started | template     | 2023-06-16T17:51:52+05:30 | <null>      | <null>    | save   | automation-email |
+------------+------------------------------------+--------------+---------------------------+-------------+-----------+--------+------------------+
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
| Credentials | Trello requires an [API key](https://trello.com/developer/marketing/guides/quick-start/#generate-your-api-key/) for all requests.                                                               |
| Permissions | API keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                             |
| Radius      | Each connection represents a single Trello Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/trello.spc`)<br />2. Credentials specified in environment variables, e.g., `TRELLO_API_KEY`.                     |

### Configuration

Installing the latest trello plugin will create a config file (`~/.steampipe/config/trello.spc`) with a single connection named `trello`:

```hcl
connection "trello" {
  plugin = "trello"

  # Trello API key for requests. Required.
  # Generate your API Key as per: https://trello.com/developer/marketing/guides/quick-start/#generate-your-api-key/
  # This can also be set via the `TRELLO_API_KEY` environment variable.
  # trello_api_key = "08355689e3e6f9fd0f5630362b16b1b5-us21"
}
```

Alternatively, you can also use the standard Trello environment variables to obtain credentials **only if other arguments (`trello_api_key`) is not specified** in the connection:

```sh
export TRELLO_API_KEY=q8355689e3e6f9fd0f5630362b16b1b5-us21
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-trello
- Community: [Slack Channel](https://steampipe.io/community/join)
