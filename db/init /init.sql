CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`username`, `password`, `email`) VALUES
('dummyuser1', 'dummypassword1', 'dummy1@example.com'),
('dummyuser2', 'dummypassword2', 'dummy2@example.com'),
('dummyuser3', 'dummypassword3', 'dummy3@example.com');