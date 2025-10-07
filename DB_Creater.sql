-- =================================================================
-- 1. 数据库清理 (可选)
-- =================================================================
DROP TABLE IF EXISTS `comments`;
DROP TABLE IF EXISTS `posts`;
DROP TABLE IF EXISTS `users`;

-- =================================================================
-- 2. users 表
-- 存储用户信息
-- =================================================================
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `email` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  UNIQUE KEY `uni_users_username` (`username`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- =================================================================
-- 3. posts 表
-- 存储博客文章信息
-- =================================================================
CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext NOT NULL,
  `content` longtext NOT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_posts_deleted_at` (`deleted_at`),
  KEY `fk_posts_user` (`user_id`),
  CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- =================================================================
-- 4. comments 表
-- 存储评论信息
-- =================================================================
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `content` longtext NOT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `post_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comments_deleted_at` (`deleted_at`),
  KEY `fk_comments_user` (`user_id`),
  KEY `fk_comments_post` (`post_id`),
  CONSTRAINT `fk_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- =================================================================
-- 5. 插入测试数据 (已增加更多数据)
-- =================================================================

-- 插入用户
-- ID 1: henry_blog (原作者)
-- ID 2: alice_reader (读者)
-- ID 3: bob_developer (新增的开发者用户)
INSERT INTO users (id, created_at, updated_at, username, password, email) VALUES
(1, NOW(), NOW(), 'alice_writer', 'hashed_password_1', 'alice@example.com'),
(2, NOW(), NOW(), 'bob_reader', 'hashed_password_2', 'bob@example.com'),
(3, NOW(), NOW(), 'charlie_dev', 'hashed_password_3', 'charlie@example.com');

-- 插入文章 (作者 ID 1: henry_blog, ID 3: bob_developer)
INSERT INTO posts (id, created_at, updated_at, title, content, user_id) VALUES
(101, NOW(), NOW(), 'Go语言并发模型初探', 'Go语言通过 Goroutine 和 Channel 实现了轻量级并发，这是其最大的特色之一。', 1),
(102, NOW(), NOW(), '为什么使用 Gin 框架', 'Gin 是一个用 Go 编写的 HTTP Web 框架，具有高性能、路由分组等特性。', 1),
(103, NOW(), NOW(), '数据库事务隔离级别', '介绍 READ UNCOMMITTED, READ COMMITTED, REPEATABLE READ, SERIALIZABLE 四种隔离级别及其问题。', 3),
(104, NOW(), NOW(), '我的旅行日记：京都之旅', '在京都看到了美丽的红叶和古老的寺庙，分享一些照片。', 1);
-- 插入评论 (共 8 条评论)
INSERT INTO comments (id, created_at, updated_at, content, user_id, post_id) VALUES
-- 评论 文章 101
(1001, NOW(), NOW(), '很棒的文章，对 Goroutine 的理解更深入了！', 2, 101),
(1002, NOW(), NOW(), '请问 Channel 的缓冲和非缓冲有什么区别呢？', 3, 101),

-- 评论 文章 102
(1003, NOW(), NOW(), '除了 Gin，有没有推荐其他 Go 框架？', 2, 102),

-- 评论 文章 103
(1004, NOW(), NOW(), '关于幻读的解释非常到位，谢谢分享！', 1, 103),

-- 评论 文章 104
(1005, NOW(), NOW(), '照片太美了！下次去我也要去那里。', 2, 104);