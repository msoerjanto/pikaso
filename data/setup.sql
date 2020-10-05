drop table pieces;
drop table artists;
drop table sessions;
drop table users;


create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table artists (
  id          serial primary key,
  uuid        varchar(64) not null unique,
  first_name  text,
  last_name   text,
  description text,
  user_id     integer references users(id),
  created_at  timestamp not null       
);

create table pieces (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  artist_id  integer references artists(id),
  created_at timestamp not null  
);