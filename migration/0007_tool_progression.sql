-- +migrate Up

-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS `tools`
(
    `id`         int          NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) NOT NULL default '',
    `ingame_id`  int          NOT NULL,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS `tool_progress`
(
    `id`           int        NOT NULL AUTO_INCREMENT,
    `tool_id`      int        NOT NULL,
    `user_id`      int        NOT NULL,
    `kills`        int        NOT NULL,
    `max_distance` int        NOT NULL,
    `is_official`  tinyint(1) NOT NULL,
    `created_at`   timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
-- +migrate StatementEnd