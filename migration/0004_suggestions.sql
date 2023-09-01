-- +migrate Up

-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS `suggestions`
(
    `id`         int                                                   NOT NULL AUTO_INCREMENT,
    `player_id`  int                                                            DEFAULT NULL,
    `feedback`   text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` timestamp                                             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp                                             NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (player_id) REFERENCES player (id) on update cascade on delete set null
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
-- +migrate StatementEnd