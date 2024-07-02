CREATE TABLE `app_applications`
(
    `id`          int          NOT NULL AUTO_INCREMENT,
    `app_id`      char(24)     NOT NULL COMMENT 'app id',
    `name`        varchar(100) NOT NULL COMMENT '名称',
    `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
    `scope`       varchar(255) NOT NULL COMMENT '作用域范围',
    `type`        tinyint(2)   NOT NULL COMMENT '类型：0=管理端,1=用户端',
    `public_key`  text         NOT NULL COMMENT '公钥',
    `private_key` text         NOT NULL COMMENT '私钥',
    `package`     varchar(100) NOT NULL COMMENT '包名',
    `project`     varchar(100) NOT NULL COMMENT '项目标识',
    `manage`      tinyint(1)   NOT NULL COMMENT '是否管理型应用 ',
    `created_at`  datetime(3)  NOT NULL COMMENT '创建时间',
    `updated_at`  datetime(3)  NOT NULL COMMENT '更新时间',
    `deleted_at`  datetime(3)  NULL COMMENT '删除时间',
    `enabled`     tinyint(2)   NOT NULL COMMENT '状态：-1-未启用=默认，0-禁用，1-启用',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `app_udx_appid` (`app_id`),
    UNIQUE INDEX `app_udx_name` (`name`),
    UNIQUE INDEX `app_udx_package` (`package`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='应用';


CREATE TABLE `app_application_attrs`
(
    `app_id` int          NOT NULL COMMENT 'app.id',
    `name`   varchar(100) NOT NULL COMMENT '属性',
    `val`    varchar(255) NOT NULL COMMENT '值',
    KEY `app_attr_idx_app_id` (`app_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='应用的属性';


CREATE TABLE `app_application_versions`
(
    `id`           int          NOT NULL AUTO_INCREMENT,
    `app_id`       int          NOT NULL COMMENT 'app.id',
    `title`        varchar(100) NOT NULL COMMENT '标题',
    `version`      varchar(20)  NOT NULL COMMENT '版本号',
    `version_num`  int          NOT NULL COMMENT '版本号数值',
    `force`        tinyint(1)   NOT NULL COMMENT '是否强制更新',
    `url`          varchar(255) NOT NULL COMMENT '文件路径',
    `description`  varchar(255) NOT NULL COMMENT '更新内容',
    `created_at`   datetime(3)  NOT NULL COMMENT '添加时间',
    `updated_at`   datetime(3)  NOT NULL COMMENT '更新时间',
    `deleted_at`   datetime(3)  NULL COMMENT '删除时间',
    `published_at` datetime(3)  NULL COMMENT '发布时间',
    PRIMARY KEY (`id`),
    KEY `app_version_idx_app_id` (`app_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='应用的版本'
