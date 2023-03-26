CREATE TABLE `license_user`
(
    `auto_id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '授权操作用户ID，主键',
    `username`          varchar(32) NOT NULL COMMENT '用户名称，不为空',
    `password`          varchar(64) NOT NULL COMMENT '登录密码，不为空',
    `account`           varchar(64) NOT NULL COMMENT '账号，不为空',
    `access_key_id`     varchar(255)         DEFAULT NULL COMMENT '访问密钥 ID',
    `secret_access_key` varchar(255)         DEFAULT NULL COMMENT '私密访问密钥',
    `phone`             varchar(20) NOT NULL COMMENT '联系方式，不为空',
    `is_deleted`        tinyint(1) NOT NULL COMMENT '软删除，不为空',
    `created_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，不为空，自动更新',
    `updated_at`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间，不为空，自动更新',
    `deleted_at`        datetime    NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '删除时间，不为空，自动更新',
    PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='授权操作用户';

CREATE TABLE `license_admin`
(
    `auto_id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT '软件服务提供商ID，主键',
    `username`   varchar(32) NOT NULL COMMENT '用户名称，不为空',
    `password`   varchar(64) NOT NULL COMMENT '登录密码，不为空',
    `account`    varchar(64) NOT NULL COMMENT '账号，不为空',
    `phone`      varchar(20) NOT NULL COMMENT '联系方式，不为空',
    `is_deleted` tinyint(1) NOT NULL COMMENT '软删除，不为空',
    `created_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，不为空，自动更新',
    `updated_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间，不为空，自动更新',
    `deleted_at` datetime    NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '删除时间，不为空，自动更新',
    PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='软件服务提供商';

CREATE TABLE `role`
(
    `auto_id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID，主键',
    `role_name`  varchar(32) NOT NULL COMMENT '角色名称，不为空',
    `is_deleted` tinyint(1) NOT NULL COMMENT '软删除，不为空',
    `created_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，不为空，自动更新',
    `updated_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间，不为空，自动更新',
    `deleted_at` datetime    NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '删除时间，不为空，自动更新',
    PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色';

CREATE TABLE `permission`
(
    `auto_id`         BIGINT       NOT NULL AUTO_INCREMENT COMMENT '权限ID',
    `permission_name` VARCHAR(32)  NOT NULL COMMENT '权限名称',
    `permission_url`  VARCHAR(255) NOT NULL COMMENT '访问路径',
    `is_deleted`      TINYINT(1) NOT NULL COMMENT '软删除',
    `created_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`      DATETIME              DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`auto_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='权限';

CREATE TABLE `license_user_role`
(
    `auto_id`         BIGINT   NOT NULL AUTO_INCREMENT COMMENT '用户角色ID',
    `license_user_id` BIGINT   NOT NULL COMMENT '授权操作用户ID',
    `role_id`         BIGINT   NOT NULL COMMENT '角色ID',
    `is_deleted`      TINYINT(1) NOT NULL COMMENT '软删除',
    `created_at`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`      DATETIME          DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`auto_id`),
    KEY               `idx_license_user_role_license_user_id` (`license_user_id`),
    KEY               `idx_license_user_role_role_id` (`role_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='授权操作用户-角色';

CREATE TABLE `license_admin_role`
(
    `auto_id`          BIGINT   NOT NULL AUTO_INCREMENT COMMENT '用户角色ID',
    `license_admin_id` BIGINT   NOT NULL COMMENT '软件服务提供商ID',
    `role_id`          BIGINT   NOT NULL COMMENT '角色ID',
    `is_deleted`       TINYINT(1) NOT NULL COMMENT '软删除',
    `created_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`       DATETIME          DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`auto_id`),
    KEY                `idx_license_admin_role_license_admin_id` (`license_admin_id`),
    KEY                `idx_license_admin_role_role_id` (`role_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='软件服务提供商-角色';

CREATE TABLE `role_permission`
(
    `auto_id`       BIGINT   NOT NULL COMMENT '用户角色ID 主键',
    `permission_id` BIGINT   NOT NULL COMMENT '权限ID 外键，权限（ID）',
    `role_id`       BIGINT   NOT NULL COMMENT '角色ID 外键，角色（ID）',
    `is_deleted`    TINYINT(1) NOT NULL COMMENT '软删除 不为空',
    `created_at`    DATETIME NOT NULL COMMENT '创建时间 不为空，自动更新',
    `updated_at`    DATETIME NOT NULL COMMENT '更新时间 不为空，自动更新',
    `deleted_at`    DATETIME NOT NULL COMMENT '删除时间 不为空，自动更新',
    PRIMARY KEY (`auto_id`)
) COMMENT='角色-权限';
