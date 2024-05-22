CREATE DATABASE IF NOT EXISTS mydatabase;
USE mydatabase;

CREATE TABLE IF NOT EXISTS `user_character` (
 `user_character_id` CHAR(36) NOT NULL PRIMARY KEY UNIQUE,
 `character_id` CHAR(36),
 `name` VARCHAR(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `user_character` (`user_character_id`, `character_id`, `name`) VALUES (UUID(), '1', 'Test');