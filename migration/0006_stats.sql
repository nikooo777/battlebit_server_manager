-- +migrate Up

-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS `stats`
(
    `id`           int       NOT NULL AUTO_INCREMENT,
    `player_count` int       NOT NULL,
    `created_at`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
-- +migrate StatementEnd