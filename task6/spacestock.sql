/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MariaDB
 Source Server Version : 100136
 Source Host           : localhost:3306
 Source Schema         : spacestock

 Target Server Type    : MariaDB
 Target Server Version : 100136
 File Encoding         : 65001

 Date: 09/09/2019 09:24:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for apartment
-- ----------------------------
DROP TABLE IF EXISTS `apartment`;
CREATE TABLE `apartment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of apartment
-- ----------------------------
BEGIN;
INSERT INTO `apartment` VALUES (2, 'asdasda', '123131');
INSERT INTO `apartment` VALUES (3, 'asdasda', '123131ww');
INSERT INTO `apartment` VALUES (4, 'asdaswwdss2222a', '123131ww');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
