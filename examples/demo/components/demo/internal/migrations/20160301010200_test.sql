-- +migrate Up
CREATE TABLE IF NOT EXISTS `test` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `data` TEXT
) ENGINE=InnoDB CHARSET=UTF8;

-- +migrate Down
DROP TABLE test;