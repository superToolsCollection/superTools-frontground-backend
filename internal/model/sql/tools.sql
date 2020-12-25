create table tools
(
    id           varchar(100)      not null primary key,
    name         varchar(50)       null comment '工具名称',
    api          varchar(100)      null comment '工具api链接',
    api_describe longtext          null comment '工具描述',
    created_on   varchar(20)       null comment '新建时间',
    modified_on  varchar(20)       null comment '修改时间',
    deleted_on   varchar(20)       null comment '删除时间',
    is_del       tinyint default 0 null comment '是否删除 0为未删除 1为已删除',
    state        tinyint default 0 null comment '状态 0为未上线 1为上线',
    created_by   varchar(100)      null comment '创建人',
    modified_by  varchar(100)      null comment '修改人',
    constraint api
        unique (api)
)ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='工具列表';