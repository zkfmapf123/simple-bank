CREATE TABLE `accounts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `owner` varchar(255) NOT NULL unique,
  `balance` int NOT NULL,
  `currency` varchar(255) NOT NULL,
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  `timezone` varchar(255)
);

CREATE TABLE `entries` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `account_id` int,
  `amount` int NOT NULL COMMENT 'can be nagative',
  `created_at` int NOT NULL,
  FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE `transfers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` int,
  `to_account_id` int,
  `amount` int NOT NULL COMMENT 'must be positive',
  `created_at` int NOT NULL,
  FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX `accounts_index_0` ON `accounts` (`owner`);

CREATE INDEX `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_4` ON `transfers` (`from_account_id`, `to_account_id`);