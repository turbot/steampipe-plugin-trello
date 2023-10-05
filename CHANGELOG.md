## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#8](https://github.com/turbot/steampipe-plugin-trello/pull/8))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#6](https://github.com/turbot/steampipe-plugin-trello/pull/6))
- Recompiled plugin with Go version `1.21`. ([#6](https://github.com/turbot/steampipe-plugin-trello/pull/6))

## v0.0.1 [2023-07-12]

_What's new?_

- New tables added
  - [trello_board](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_board)
  - [trello_card](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_card)
  - [trello_list](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_list)
  - [trello_member](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_member)
  - [trello_my_board](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_my_board)
  - [trello_my_member](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_my_member)
  - [trello_my_notification](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_my_notification)
  - [trello_my_organization](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_my_organization)
  - [trello_organization](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_organization)
  - [trello_search_board](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_search_board)
  - [trello_search_card](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_search_card)
  - [trello_webhook](https://hub.steampipe.io/plugins/turbot/trello/tables/trello_webhook)
