
-- +migrate Up
alter table task add column user_id bigint not null after name;
-- +migrate Down
