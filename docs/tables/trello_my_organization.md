# Table: trello_my_organization

An organization is a higher-level entity that helps users manage multiple boards and collaborate with larger teams or groups. Organizations in Trello provide a way to group boards and members together, making it easier to organize and share work among team members. This table is a view of all organizations that you are a member of.

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
  trello_my_organization;
```