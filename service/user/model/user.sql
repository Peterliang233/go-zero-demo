CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `username` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户名称',
    `number` varchar(255) NOT NULL DEFAULT '' COMMENT '学号',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender` char(5)  NOT NULL DEFAULT '' COMMENT '男｜女｜未公开',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `number_unique` (`number`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;