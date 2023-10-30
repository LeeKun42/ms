/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 127.0.0.1:3309
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 30/10/2023 16:25:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_user_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_roles`;
CREATE TABLE `admin_user_roles`  (
                                     `admin_user_id` int(0) UNSIGNED NOT NULL,
                                     `role_id` int(0) NOT NULL,
                                     PRIMARY KEY (`admin_user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user_roles
-- ----------------------------
INSERT INTO `admin_user_roles` VALUES (1, 1);
INSERT INTO `admin_user_roles` VALUES (2, 1);
INSERT INTO `admin_user_roles` VALUES (4, 8);

-- ----------------------------
-- Table structure for admin_user_socials
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_socials`;
CREATE TABLE `admin_user_socials`  (
                                       `admin_user_id` int(0) UNSIGNED NOT NULL COMMENT '管理员账号id',
                                       `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '社会化登录类型:wechat、dingding、enterprisewechat',
                                       `unionid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                       `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                                       `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
                                       `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '最后更新时间',
                                       PRIMARY KEY (`admin_user_id`, `type`) USING BTREE,
                                       UNIQUE INDEX `type`(`type`, `unionid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '管理账号关联社会化登录信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user_socials
-- ----------------------------
INSERT INTO `admin_user_socials` VALUES (1, 'dingtalk', 'SXBx8nnk6jqsnknK24wrSAiEiE', 'ROjghrTrEWkfYrK3zZvReQiEiE', '2023-10-26 17:09:45', '2023-10-26 17:09:45');

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users`  (
                                `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号码',
                                `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
                                `passwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录密码',
                                `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
                                `avatar` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像地址',
                                `status` tinyint(0) NOT NULL COMMENT '状态： 1：正常  0：禁用',
                                `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
                                `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '最后更新时间',
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `account`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_users
-- ----------------------------
INSERT INTO `admin_users` VALUES (1, '13802990402', 'jun.l@sumian.com', '$2a$10$pw5GjwnI2cvVPg9qxYfvruTZfabpXKmON08VBxzTz9wGRN.afV74i', '李俊', 'https://static-legacy.dingtalk.com/media/lADObl_2lM0BSs0BvA_444_330.jpg', 1, '2023-10-19 15:53:27', '2023-10-26 14:56:05');
INSERT INTO `admin_users` VALUES (2, '13802990401', '123@qq.com', '$2a$10$I75iQffYxBjsa0YasrIj.O.i5/6Gb4.lc8qSKibMgyv7KoouYrEqm', '李俊2', '', 1, '2023-10-23 14:26:15', '2023-10-26 11:07:24');
INSERT INTO `admin_users` VALUES (4, '13802990405', '321@qq.com', '$2a$10$Dl.h/AI8bHaRQX2BqKiNHOAOBJ2ephAWlzZiUUhg68KwSISolBS2i', '李俊3', '', 1, '2023-10-26 11:07:51', '2023-10-26 11:07:51');

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions`  (
                                `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                `parent_id` int(0) UNSIGNED NOT NULL COMMENT '父级权限id：操作权限的父级id为所属页面权限id',
                                `flag` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限标识',
                                `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '权限名称',
                                `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限说明',
                                `type` tinyint(0) NOT NULL COMMENT '权限类型：10：页面权限  20：操作权限',
                                `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
                                `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '最后更新时间',
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `flag`(`flag`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permissions
-- ----------------------------
INSERT INTO `permissions` VALUES (1, 0, 'system-user-view', '管理员账号管理', '菜单权限', 10, '2023-10-23 15:08:40', '2023-10-23 16:08:37');
INSERT INTO `permissions` VALUES (3, 1, 'system-user-create', '管理员账号创建', '操作权限', 20, '2023-10-23 15:11:04', '2023-10-25 14:26:56');
INSERT INTO `permissions` VALUES (4, 0, 'system-role-view', '角色管理', '菜单权限', 10, '2023-10-23 16:07:05', '2023-10-23 16:08:37');
INSERT INTO `permissions` VALUES (6, 0, 'system-permission-view', '权限管理', '菜单权限', 10, '2023-10-23 16:08:03', '2023-10-23 16:08:39');

-- ----------------------------
-- Table structure for role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions`  (
                                     `role_id` int(0) UNSIGNED NOT NULL,
                                     `permission_id` int(0) UNSIGNED NOT NULL,
                                     PRIMARY KEY (`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_permissions
-- ----------------------------
INSERT INTO `role_permissions` VALUES (1, 1);
INSERT INTO `role_permissions` VALUES (1, 3);
INSERT INTO `role_permissions` VALUES (1, 4);
INSERT INTO `role_permissions` VALUES (1, 6);
INSERT INTO `role_permissions` VALUES (8, 1);
INSERT INTO `role_permissions` VALUES (8, 3);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
                          `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                          `flag` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色flag',
                          `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
                          `is_system` tinyint(0) UNSIGNED NOT NULL COMMENT '是否是系统内置角色 1：是 0：否',
                          `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
                          `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '最后更新时间',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `flag`(`flag`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, 'super_admin', '超级管理员', 1, '2023-10-23 11:04:07', '2023-10-23 11:04:07');
INSERT INTO `roles` VALUES (8, 'system-admin', '系统管理员', 0, '2023-10-25 15:48:21', '2023-10-25 15:48:21');

SET FOREIGN_KEY_CHECKS = 1;
