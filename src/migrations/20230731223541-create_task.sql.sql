
-- +migrate Up
create table task(
    id bigint not null primary key auto_increment,
    name varchar(256) not null,
    created_at datetime not null
);
-- +migrate Down
