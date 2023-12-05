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

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-trello/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-trello/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Trello Plugin](https://github.com/turbot/steampipe-plugin-trello/labels/help%20wanted)
