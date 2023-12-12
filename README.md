![image](https://hub.steampipe.io/images/plugins/turbot/trello-social-graphic.png)

# Trello Plugin for Steampipe

Use SQL to query boards, cards, lists, and more from Trello.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/trello)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/trello/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-trello/issues)

## Quick start

### Install

Download and install the latest Trello plugin:

```bash
steampipe plugin install trello
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/trello#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/trello#configuration).

Configure your account details in `~/.steampipe/config/trello.spc`:

```hcl
connection "trello" {
  plugin = "trello"

  # Authentication information
  api_key = "a25ad2e37570117c0bad72d0a711ba5af"
  token = "ATTAb179ea3c211722b0ebb2d223e1922b5e1ab1d28a3caac8d3722a83e9f91f25b973FDCC07"
}
```

Or through environment variables:

```sh
export TRELLO_API_KEY=a25ad2e37570117c0bad72d0a711ba5af
export TRELLO_TOKEN=ATTAb179ea3c211722b0ebb2d223e1922b5e1ab1d28a3caac8d3722a83e9f91f25b973FDCC07
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-trello.git
cd steampipe-plugin-trello
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/trello.spc
```

Try it!

```
steampipe query
> .inspect trello
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Trello Plugin](https://github.com/turbot/steampipe-plugin-trello/labels/help%20wanted)
