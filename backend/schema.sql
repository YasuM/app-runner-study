CREATE TABLE `task` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  `created_at` datetime NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`id`)
);