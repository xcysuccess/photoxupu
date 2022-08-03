CREATE TABLE `users` (
                         `user_id` int NOT NULL AUTO_INCREMENT,
                         `user_name` varchar(60) DEFAULT NULL,
                         `password` varchar(32) DEFAULT NULL,
                         `sex` int DEFAULT NULL,
                         `birthday` timestamp NULL DEFAULT NULL,
                         `last_login` timestamp NULL DEFAULT NULL,
                         `bind_phone` varchar(20) DEFAULT NULL,
                         `reg_time` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb3