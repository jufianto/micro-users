# My Analyst for Microservices Users 
it will use for almost application system with user managing

## Global Interface Function

### `generate_uid(timestamp)`
uid generate from timestamp in unixtimestamp with algorithm SHA1
with format `28-4` character

## Table Design
### user
* username (varchar 100) NOT NULL
* password (text) NOT NULL
* remember_token (text) NULL
* reset_token (text) NULL
* created_at (timestamp)
  
### detail_user
* uid varchar(32)
* username varchar(100) NOT NULL
* uid_identity varchar(32)
* no_identity varchar(100)
* full_name text NOT NULL
* gender varchar(25) (L/P/NULL)
* created_at timestamp

### identity_name
* uid varchar(32)
* name_identity varchar(100)
* information text
* crated_at (timestamp)

### user_log
* uid
* created_at (timestamp)
* action test NOT NULL
* status varchar(100) NULL

### role
* uid
* role_name
* information

### user_role
* uid
* uid_user
* uid_role
* created_at
* updated_at

## Action User 
* login 
* logout
* send_telegram   
   Send user action to telegram channels / user admin  
* add_user
* update_user
* delete_user

