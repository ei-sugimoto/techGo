CREATE DATABASE IF NOT EXISTS mydatabase;
USE mydatabase;

CREATE TABLE IF NOT EXISTS `user_character` (
    `user_character_id` CHAR(36) NOT NULL PRIMARY KEY UNIQUE,
    `name` VARCHAR(255) NOT NULL,
    `user_id` KEY (`user_id`) REFERENCES `user` (`user_id`),
    `character_id` KEY (`character_id`) REFERENCES `character` (`character_id`)
) ;

CREATE TABLE IF NOT EXISTS `user` (
    `user_id` CHAR(36) NOT NULL PRIMARY KEY UNIQUE,
    `name` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `character` (
    `character_id` CHAR(36) NOT NULL PRIMARY KEY UNIQUE,
    `name` VARCHAR(255) NOT NULL
    `probability` INT NOT NULL
);

INSERT INTO `user` (`user_id`, `name`) VALUES (UUID(), 'Test');
INSERT INTO `character` (`character_id`, `name`, `probability`) VALUES (UUID(), 'Test', 50);
INSERT INTO `user_character` (`user_character_id`, `name`, `user_id`, `character_id`) VALUES (UUID(), 'Test', (SELECT `user_id` FROM `user` WHERE `name` = 'Test'), (SELECT `character_id` FROM `character` WHERE `name` = 'Test'));