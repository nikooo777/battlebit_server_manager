-- +migrate Up

-- +migrate StatementBegin
alter table player_progress
    add is_official tinyint not null comment 'when 1 it means the stats are from the official game, when 0 it''s from our own servers' after player_id;
-- +migrate StatementEnd

-- +migrate StatementBegin
alter table player_progress
    drop foreign key player_progress_ibfk_1;
-- +migrate StatementEnd

-- +migrate StatementBegin
alter table player_progress
    drop key player_progress_player_uniq;
-- +migrate StatementEnd

-- +migrate StatementBegin
alter table player_progress
    add constraint player_progress_player_uniq
        unique (player_id, is_official);
-- +migrate StatementEnd

-- +migrate StatementBegin
alter table player_progress
    add constraint player_progress_ibfk_1
        foreign key (player_id) references player (id)
            on update cascade on delete cascade;
-- +migrate StatementEnd

-- +migrate StatementBegin
update player_progress set is_official = 1;
-- +migrate StatementEnd