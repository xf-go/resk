-- 红包商品表的简化版本，为了方便测试评估，只保留了部分字段
-- 红包商品表
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`
(
  `id`              bigint(20)     NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `envelope_no`     varchar(32)    NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `remain_amount`   decimal(30, 6) NOT NULL DEFAULT '0.000000' COMMENT '红包剩余金额',
  `remain_quantity` int(10)        NOT NULL COMMENT '红包剩余数量 ',
  `created_at`      datetime(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at`      datetime(3)    NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `envelope_no_idx` (`envelope_no`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8
  ROW_FORMAT = DYNAMIC;

-- 红包商品表的无符号类型字段版本
DROP TABLE IF EXISTS `goods_unsigned`;
CREATE TABLE `goods_unsigned`
(
  `id`              bigint(20)              NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `envelope_no`     varchar(32)             NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `remain_amount`   decimal(30, 6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '红包剩余金额',
  `remain_quantity` int(10) unsigned        NOT NULL COMMENT '红包剩余数量 ',
  `created_at`      datetime(3)             NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at`      datetime(3)             NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `envelope_no_idx` (`envelope_no`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8
  ROW_FORMAT = DYNAMIC;