-- Create "chat_groups" table
CREATE TABLE `chat_groups` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL,
  `name` varchar(255) NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `chat_groups_users_ChatGroups` (`user_id`),
  CONSTRAINT `chat_groups_users_ChatGroups` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "chat_group_members" table
CREATE TABLE `chat_group_members` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `chat_group_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `chat_group_members_chat_groups_ChatGroupMembers` (`chat_group_id`),
  INDEX `chat_group_members_users_ChatGroupMembers` (`user_id`),
  CONSTRAINT `chat_group_members_chat_groups_ChatGroupMembers` FOREIGN KEY (`chat_group_id`) REFERENCES `chat_groups` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `chat_group_members_users_ChatGroupMembers` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "chat_messages" table
CREATE TABLE `chat_messages` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL,
  `message_text` longtext NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `chat_group_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `chat_messages_chat_groups_ChatMessages` (`chat_group_id`),
  INDEX `chat_messages_users_ChatMessages` (`user_id`),
  CONSTRAINT `chat_messages_chat_groups_ChatMessages` FOREIGN KEY (`chat_group_id`) REFERENCES `chat_groups` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `chat_messages_users_ChatMessages` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
