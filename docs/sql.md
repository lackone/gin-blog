### 标签表

```sql
CREATE TABLE `blog_tags`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`       varchar(32) NOT NULL DEFAULT '' COMMENT '标签名',
    `created`    int unsigned NOT NULL COMMENT '创建时间',
    `created_by` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated`    int unsigned NOT NULL COMMENT '更新时间',
    `updated_by` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
    `deleted`    int unsigned NOT NULL COMMENT '删除时间',
    `status`     tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态(0禁用，1启用)',
    `is_del`     tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除(0未删除，1已删除)',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='标签表';
```

### 权限表

```sql
CREATE TABLE `blog_auths`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `app_key`    varchar(32) NOT NULL DEFAULT '' COMMENT 'key',
    `app_secret` varchar(64) NOT NULL DEFAULT '' COMMENT 'secret',
    `created`    int unsigned NOT NULL COMMENT '创建时间',
    `created_by` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated`    int unsigned NOT NULL COMMENT '更新时间',
    `updated_by` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
    `deleted`    int unsigned NOT NULL COMMENT '删除时间',
    `is_del`     tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除(0未删除，1已删除)',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限表';
```

### 文章标签表

```sql
CREATE TABLE `blog_article_tags`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `article_id` int unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
    `tag_id`     int unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
    `created`    int unsigned NOT NULL COMMENT '创建时间',
    `created_by` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated`    int unsigned NOT NULL COMMENT '更新时间',
    `updated_by` varchar(32) NOT NULL DEFAULT '' COMMENT '更新人',
    `deleted`    int unsigned NOT NULL COMMENT '删除时间',
    `is_del`     tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除(0未删除，1已删除)',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章标签表';
```

### 文章表

```sql
CREATE TABLE `blog_articles`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `title`      varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
    `desc`       varchar(255) NOT NULL DEFAULT '' COMMENT '简述',
    `cover`      varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片',
    `content`    longtext     NOT NULL COMMENT '内容',
    `created`    int unsigned NOT NULL COMMENT '创建时间',
    `created_by` varchar(32)  NOT NULL DEFAULT '' COMMENT '创建人',
    `updated`    int unsigned NOT NULL COMMENT '更新时间',
    `updated_by` varchar(32)  NOT NULL DEFAULT '' COMMENT '更新人',
    `deleted`    int unsigned NOT NULL COMMENT '删除时间',
    `is_del`     tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除(0未删除，1已删除)',
    `status`     tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态(0禁用，1启用)',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章表';
```