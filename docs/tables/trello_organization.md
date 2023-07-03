# Table: trello_organization

An organization is a higher-level entity that helps users manage multiple boards and collaborate with larger teams or groups. Organizations in Trello provide a way to group boards and members together, making it easier to organize and share work among team members.

The `trello_organization` table can be used to query information about **ANY** organization, and **you must specify which organization** in the where clause (`where id=`).

## Examples

### Basic info

```sql
select
  id,
  name,
  description,
  display_name,
  url,
  website
from 
  trello_organization
where
  id = '123ace0f581f4de8a0dc184c';
```