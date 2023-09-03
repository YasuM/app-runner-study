
-- +migrate Up
create table user(
    id bigint not null primary key auto_increment,
    name varchar(256) not null,
    email varchar(256) not null,
    password varchar(256) not null,
    created_at datetime not null
);
-- +migrate Down
