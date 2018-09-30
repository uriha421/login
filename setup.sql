drop table if exists users cascade;
drop table if exists sessions;

create table users (
    id		serial primary key,
    email	varchar(256) not null unique,
    password	varchar(256) not null,
    created_at	timestamp not null
);

create table sessions (
    id		serial primary key,
    uuid	varchar(64) not null unique,
    user_id	integer references users(id),
    created_at	timestamp not null
);

