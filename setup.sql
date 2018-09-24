drop table if exists users cascade;
drop table if exists sessions;

create table users (
  id         serial primary key,
  uuid varchar(64) not null unique,
  name varchar(255) not null,
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  admin smallint not null,
  created_at timestamp not null
);

create table sessions (
    id	    serial primary key,
    uuid    varchar(64) not null unique,
    user_id integer references users(id),
    admin smallint not null,
    created_at timestamp not null
);

