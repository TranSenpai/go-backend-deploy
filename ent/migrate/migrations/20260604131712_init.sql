-- Create "users" table
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL,
  `email` varchar(255) NOT NULL,
  `full_name` varchar(255) NULL,
  `avatar` varchar(255) NULL,
  `password` varchar(255) NULL,
  `totp_secret` varchar(255) NULL,
  `google_id` varchar(255) NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email` (`email`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "articles" table
CREATE TABLE `articles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL,
  `title` varchar(255) NOT NULL,
  `content` varchar(255) NULL,
  `image_url` varchar(255) NULL,
  `like_count` bigint NOT NULL DEFAULT 0,
  `views` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `articles_users_Articles` (`user_id`),
  CONSTRAINT `articles_users_Articles` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "foods" table
CREATE TABLE `foods` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "orders" table
CREATE TABLE `orders` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `food_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `orders_foods_Orders` (`food_id`),
  INDEX `orders_users_Orders` (`user_id`),
  CONSTRAINT `orders_foods_Orders` FOREIGN KEY (`food_id`) REFERENCES `foods` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `orders_users_Orders` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
