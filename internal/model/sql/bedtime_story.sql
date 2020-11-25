create table stories
(
    id          int auto_increment primary key,
    author      varchar(50)                       null comment '作者',
    story       longtext                          null comment '睡前故事',
    is_del      tinyint      default 0            null comment '是否删除',
    create_on   date         default '2020-09-16' null comment '新建时间',
    modified_on date         default '2020-09-16' null comment '修改时间',
    delete_on   date                              null comment '删除时间',
    state       tinyint      default 1            null comment '状态 0为禁用 1为启用',
    created_by  varchar(100) default 'super'      null comment '创建人',
    modified_by varchar(100)                      null comment '修改人'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='睡前故事';


CREATE TABLE `story_tag`
(
    `id`    int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`  VARCHAR(100)        DEFAULT '' COMMENT '标签名称',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用 1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='睡前故事标签';

CREATE TABLE `story_tag_map`
(
    `id`        int(10) unsigned                  NOT NULL AUTO_INCREMENT,
    `story_id`  int                               NOT NULL COMMENT '睡前故事ID',
    `tag_id`    int(10) unsigned                  NOT NULL DEFAULT 0 COMMENT '标签ID',
    create_on   date         default '2020-09-16' null comment '新建时间',
    modified_on date         default '2020-09-16' null comment '修改时间',
    delete_on   date                              null comment '删除时间',
    created_by  varchar(100) default 'super'      null comment '创建人',
    modified_by varchar(100)                      null comment '修改人',
    is_del      tinyint      default 0            null comment '是否删除 0为未删除 1为已删除',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='睡前故事标签关联';

CREATE TABLE `story_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)        DEFAULT '' COMMENT '标签名称',
    `created_on`  int(10) unsigned    DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`  int(10) unsigned    DEFAULT '0' COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='标签管理';

CREATE TABLE `story_auth`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key`     varchar(20)         DEFAULT '' COMMENT 'Key',
    `app_secret`  VARCHAR(50)         DEFAULT '' COMMENT 'Secret',
    `created_on`  VARCHAR(20)         DEFAULT '0' COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` VARCHAR(20)         DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `deleted_on`  VARCHAR(20)         DEFAULT '0' COMMENT '删除时间',
    `is_del`      tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='认证管理';