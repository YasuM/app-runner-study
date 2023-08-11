
-- +migrate Up
alter table task add column status int not null;
-- +migrate Down
