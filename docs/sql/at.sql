CREATE TABLE `at_merchant` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '密码',
  `ctime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'create time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;