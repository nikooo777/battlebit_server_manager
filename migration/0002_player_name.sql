-- +migrate Up

-- +migrate StatementBegin
alter table player add name varchar(255) not null after steam_id;
-- +migrate StatementEnd