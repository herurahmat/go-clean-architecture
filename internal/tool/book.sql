CREATE TABLE `books` (
 `id` varchar(255) NOT NULL,
 `title` varchar(255) NOT NULL,
 `author_id` varchar(255) NOT NULL,
 PRIMARY KEY (`id`),
 KEY `id` (`id`),
 KEY `author_id` (`author_id`)
) ENGINE=InnoDB