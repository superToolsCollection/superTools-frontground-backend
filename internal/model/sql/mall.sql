CREATE TABLE `products`
(
    `id`            int(50) unsigned NOT NULL AUTO_INCREMENT,
    `product_name`  varchar(100) DEFAULT '' COMMENT '商品名',
    `product_num`   int COMMENT '商品数量',
    `product_image` VARCHAR(200) DEFAULT '' COMMENT '商品图片',
    `product_url`   varchar(200) DEFAULT '' COMMENT '商品地址',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='商品表';

create table orders
(
    id         int primary key,
    user_id    int               null comment '用户id',
    product_id int               null comment '商品id',
    state      tinyint default 0 null comment '状态 0为待支付 1为下单成功 2为下单失败'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='秒杀系统订单';