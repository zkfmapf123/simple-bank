# Simple Bank

<details>
<summary>DB Schema</summary>
<div markdown=1>

## Schema

<img src="./public/schema.png">

## Script

```sql
Table accounts as A {
  id int [pk, increment]
  owner varchar [not null]
  balance bigint [not null]
  currency varchar [not null]
  created_at bigint [not null,default: `now()`]
  updated_at bigint [not null,default: `now()`]
  timezone varchar

  Indexes {
    owner
  }
}

Table entries as E {
  id int [pk, increment]
  account_id bigint [ref: > A.id]
  amount bigint [not null, note: 'can be nagative']
  created_at bigint [not null,default: `now()`]

  Indexes {
    account_id
  }
}

Table transfers {
  id int [pk, increment]
  from_account_id bigint [ref: > A.id]
  to_account_id bigint [ref: > A.id]
  amount bigint [not null, note: 'must be positive']
  created_at bigint [not null,default: `now()`]

  Indexes {
    from_account_id
    to_account_id
    (from_account_id, to_account_id)
  }
}
```

</details>
