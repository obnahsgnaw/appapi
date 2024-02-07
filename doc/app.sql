CREATE TABLE `app_applications`
(
    `id`          int(10) unsigned    NOT NULL AUTO_INCREMENT,
    `app_id`      char(24) unique     NOT NULL COMMENT 'app id',
    `name`        varchar(100) unique NOT NULL COMMENT '名称',
    `description` varchar(255)        NOT NULL DEFAULT '' COMMENT '描述',
    `scope`       varchar(255)        NOT NULL COMMENT '作用域范围',
    `type`        tinyint(2) unsigned NOT NULL COMMENT '类型：0=管理端,1=用户端,2=混合端',
    `public_key`  text                NOT NULL COMMENT '公钥',
    `private_key` text                NOT NULL COMMENT '私钥',
    `package`     varchar(100) unicode NOT NULL COMMENT '包名',
    `project`     varchar(100)        NOT NULL COMMENT '项目标识',
    `manage`      tinyint(1)          NOT NULL COMMENT '是否管理型应用 ',
    `created_at`  datetime(3)         NOT NULL COMMENT '创建时间',
    `updated_at`  datetime(3)         NOT NULL COMMENT '更新时间',
    `deleted_at`  datetime(3)         NULL COMMENT '删除时间',
    `enabled`      tinyint(2)          NOT NULL COMMENT '状态：-1-未启用=默认，0-禁用，1-启用',
    PRIMARY KEY (`id`),
    KEY `idx_appid` (`app_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='应用';


CREATE TABLE `app_application_attrs`
(
    `app_id` int(10) unsigned unique NOT NULL COMMENT 'app.id',
    `name`   varchar(100)            NOT NULL COMMENT '属性',
    `val`    varchar(255)            NOT NULL COMMENT '值',
    PRIMARY KEY (`app_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='应用的属性';


CREATE TABLE `app_application_versions`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_id`        int(10) unsigned NOT NULL COMMENT 'app.id',
    `title`         varchar(100)     NOT NULL COMMENT '标题',
    `version`       varchar(20)      NOT NULL COMMENT '版本号',
    `version_num`   int(10) unsigned NOT NULL COMMENT '版本号数值',
    `force`         tinyint(1)       NOT NULL COMMENT '是否强制更新',
    `url`           varchar(255)     NOT NULL COMMENT '文件路径',
    `description`   varchar(255)     NOT NULL COMMENT '更新内容',
    `created_at`    datetime(3)      NOT NULL COMMENT '添加时间',
    `updated_at`    datetime(3)      NOT NULL COMMENT '更新时间',
    `deleted_at`    datetime(3)      NULL COMMENT '删除时间',
    `published_at`  datetime(3)      NULL COMMENT '发布时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='应用的版本'
