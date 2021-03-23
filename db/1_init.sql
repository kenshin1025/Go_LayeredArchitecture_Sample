CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT(20) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(256) NOT NULL,
  `email` VARCHAR(256) NOT NULL UNIQUE
)