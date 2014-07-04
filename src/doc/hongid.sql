SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


CREATE TABLE IF NOT EXISTS `hd_admin` (
  `id` mediumint(6) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `roleid` smallint(5) DEFAULT '0' COMMENT '角色',
  `lastloginip` varchar(15) DEFAULT '0.0.0.0' COMMENT '最后登陆地址PI',
  `lastlogintime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登陆时间',
  `email` varchar(40) DEFAULT NULL COMMENT '邮箱',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `lang` varchar(6) NOT NULL DEFAULT 'zh-cn' COMMENT '语言',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1:允许登陆 0:禁止登陆 ',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `username` (`username`,`roleid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='管理员表' AUTO_INCREMENT=2 ;

--创始人
INSERT INTO `hd_admin` (`id`, `username`, `password`, `roleid`, `lastloginip`, `lastlogintime`, `email`, `realname`, `lang`, `status`, `createtime`) VALUES
(1, 'admin', '$2a$10$yNj7fzAZ5J6EmEW17q7R7OaE7bRF1a3FvpXgr3l/QGTLTYFm2Apq2', 1, '127.0.0.1', '2014-06-28 15:00:20', 'zi__chen@163.com', 'admin', 'zh-cn', 1, '2014-06-28 15:00:20');

--
-- 表的结构 `logs`
--

CREATE TABLE IF NOT EXISTS `hd_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` int(10) unsigned NOT NULL COMMENT 'uid',
  `module` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '模型',
  `url` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作对应的url',
  `action` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作对应的action',
  `ip` varchar(15) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0.0.0.0' COMMENT '操作者所在IP',
  `desc` text COLLATE utf8_unicode_ci NOT NULL COMMENT '操作说明',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`,`module`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='后台操作日志表' AUTO_INCREMENT=1 ;


--
-- 表的结构 `menu`
--

CREATE TABLE IF NOT EXISTS `hd_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `pid` int(11) NOT NULL DEFAULT '0',
  `name` char(40) NOT NULL DEFAULT '' COMMENT '名称',
  `enname` char(40) NOT NULL DEFAULT '' COMMENT '英文名称',
  `url` char(100) NOT NULL DEFAULT '' COMMENT '功能地址',
  `data` char(100) DEFAULT '' COMMENT '附加参数',
  `order` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `display` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示，1:显示 0:不显示',
  PRIMARY KEY (`id`),
  KEY `listorder` (`order`),
  KEY `parentid` (`pid`),
  KEY `module` (`url`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='后台菜单表' AUTO_INCREMENT=34 ;

--
-- 转存表中的数据 `menu`
--

INSERT INTO `hd_menu` (`id`, `pid`, `name`, `enname`, `url`, `data`, `order`, `display`) VALUES
(1, 0, '我的面板', 'Panel', 'Panel', '', 10000, 1),
(2, 0, '设置', 'Settings', 'Setting', '', 20000, 1),
(3, 0, '模块', 'Modules', 'Module', '', 30000, 1),
(4, 0, '内容', 'Content', 'Content', '', 40000, 1),
(5, 0, '用户', 'Users', 'User', '', 50000, 1),
(6, 0, '扩展', 'Extensions', 'Extend', '', 60000, 1),
(7, 0, '界面', 'Templates', 'Style', '', 70000, 1),
(8, 0, '应用', 'Plugin', 'Plugin', '', 80000, 1),
(9, 2, '菜单设置', 'Menu Settings', 'javascript:;', '', 20100, 1),
(10, 9, '菜单管理', 'Menu management', 'Menu', '', 20101, 1),
(11, 1, '个人设置', 'Personal Settings', 'javascript:;', '', 10100, 1),
(12, 11, '个人信息', 'Personal information', 'EditInfo', '', 10101, 1),
(13, 11, '修改密码', 'Change password', 'EditPwd', '', 10102, 1),
(14, 2, '管理员管理', 'Administrator manager', 'javascript:;', '', 20200, 1),
(15, 14, '管理员管理', 'Administrator manager', 'Admin', '', 20201, 1),
(16, 14, '角色管理', 'Role management', 'Role', '', 20202, 1),
(17, 2, '日志管理', 'Log management', 'javascript:;', '', 20300, 1),
(18, 17, '日志管理', 'Log management', 'Logs', '', 20301, 1),
(19, 1, '快捷面板', 'Shortcut panel', 'javascript:;', '', 10200, 1),
(20, 4, '内容管理', 'Content management', 'javascript:;', '', 40100, 1),
(21, 4, '相关设置', 'Related settings', 'javascript:;', '', 40200, 1),
(22, 20, '栏目管理', 'Manage column', 'Category', '', 40101, 1),
(23, 20, '内容管理', 'Manage content', 'Content', '', 40102, 1),
(24, 3, '模块管理', 'Manage module', 'javascript:;', '', 30100, 1),
(25, 24, '公告', 'Announcement', 'Announce', '', 30101, 1),
(26, 6, '扩展', 'Extensions', 'javascript:;', '', 60100, 1),
(27, 26, '来源管理', 'Source management', 'Copyfrom', '', 60101, 1),
(28, 5, '会员管理', 'Manage user', 'javascript:;', '', 50100, 1),
(29, 5, '会员组管理', 'Manage user group', 'javascript:;', '', 50200, 1),
(30, 28, '会员管理', 'Manage user', 'User', '', 50101, 1),
(31, 29, '管理会员组', 'Manage user group', 'Group', '', 50201, 1),
(32, 7, '模板管理', 'Manage template', 'javascript:;', '', 70100, 1),
(33, 32, '模板风格', 'Style template', 'Style', '', 70101, 1);

--
-- 表的结构 `admin_panel`
--

CREATE TABLE IF NOT EXISTS `hd_admin_panel` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `name` varchar(40) DEFAULT '' COMMENT '菜单名称',
  `url` varchar(255) DEFAULT '' COMMENT '菜单url',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  UNIQUE KEY `uid` (`id`,`aid`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='快捷面板' AUTO_INCREMENT=5 ;


--
-- 表的结构 `role`
--

CREATE TABLE IF NOT EXISTS `hd_role` (
  `id` int(3) unsigned NOT NULL AUTO_INCREMENT,
  `rolename` varchar(50) NOT NULL COMMENT '角色名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '角色说明',
  `data` text NOT NULL COMMENT '菜单列表',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否启用',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `roleid` (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表' AUTO_INCREMENT=3 ;

--
-- 转存表中的数据 `role`
--

INSERT INTO `hd_role` (`id`, `rolename`, `desc`, `data`, `status`, `createtime`) VALUES
(1, '超级管理员', '超级管理员', '1,11,12,13,2,9,10,14,15,16,17,18,3,4,5,6,7,8', 1, '2014-06-28 20:20:09');

--
-- 表的结构 `menu`
--

CREATE TABLE IF NOT EXISTS `menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `pid` int(11) NOT NULL DEFAULT '0',
  `name` char(40) NOT NULL DEFAULT '' COMMENT '名称',
  `enname` char(40) NOT NULL DEFAULT '' COMMENT '英文名称',
  `url` char(100) NOT NULL DEFAULT '' COMMENT '功能地址',
  `data` char(100) DEFAULT '' COMMENT '附加参数',
  `order` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `display` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示，1:显示 0:不显示',
  PRIMARY KEY (`id`),
  KEY `listorder` (`order`),
  KEY `parentid` (`pid`),
  KEY `module` (`url`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='后台菜单表' AUTO_INCREMENT=34 ;

INSERT INTO `menu` (`id`, `pid`, `name`, `enname`, `url`, `data`, `order`, `display`) VALUES
(1, 0, '我的面板', 'Panel', 'Panel', '', 10000, 1),
(2, 0, '设置', 'Settings', 'Setting', '', 20000, 1),
(3, 0, '模块', 'Modules', 'Module', '', 30000, 1),
(4, 0, '内容', 'Content', 'Content', '', 40000, 1),
(5, 0, '用户', 'Users', 'User', '', 50000, 1),
(6, 0, '扩展', 'Extensions', 'Extend', '', 60000, 1),
(7, 0, '界面', 'Templates', 'Style', '', 70000, 1),
(8, 0, '应用', 'Plugin', 'Plugin', '', 80000, 1),
(9, 2, '菜单设置', 'Menu Settings', 'javascript:;', '', 20100, 1),
(10, 9, '菜单管理', 'Menu management', 'Menu', '', 20101, 1),
(11, 1, '个人设置', 'Personal Settings', 'javascript:;', '', 10100, 1),
(12, 11, '个人信息', 'Personal information', 'EditInfo', '', 10101, 1),
(13, 11, '修改密码', 'Change password', 'EditPwd', '', 10102, 1),
(14, 2, '管理员管理', 'Administrator manager', 'javascript:;', '', 20200, 1),
(15, 14, '管理员管理', 'Administrator manager', 'Admin', '', 20201, 1),
(16, 14, '角色管理', 'Role management', 'Role', '', 20202, 1),
(17, 2, '日志管理', 'Log management', 'javascript:;', '', 20300, 1),
(18, 17, '日志管理', 'Log management', 'Logs', '', 20301, 1),
(19, 1, '快捷面板', 'Shortcut panel', 'javascript:;', '', 10200, 1),
(20, 4, '内容管理', 'Content management', 'javascript:;', '', 40100, 1),
(21, 4, '相关设置', 'Related settings', 'javascript:;', '', 40200, 1),
(22, 20, '栏目管理', 'Manage column', 'Category', '', 40101, 1),
(23, 20, '内容管理', 'Manage content', 'Content', '', 40102, 1),
(24, 3, '模块管理', 'Manage module', 'javascript:;', '', 30100, 1),
(25, 24, '公告', 'Announcement', 'Announce', '', 30101, 1),
(26, 6, '扩展', 'Extensions', 'javascript:;', '', 60100, 1),
(27, 26, '来源管理', 'Source management', 'Copyfrom', '', 60101, 1),
(28, 5, '会员管理', 'Manage user', 'javascript:;', '', 50100, 1),
(29, 5, '会员组管理', 'Manage user group', 'javascript:;', '', 50200, 1),
(30, 28, '会员管理', 'Manage user', 'User', '', 50101, 1),
(31, 29, '管理会员组', 'Manage user group', 'Group', '', 50201, 1),
(32, 7, '模板管理', 'Manage template', 'javascript:;', '', 70100, 1),
(33, 32, '模板风格', 'Style template', 'Style', '', 70101, 1);


--
-- 表的结构 `announce`
--

CREATE TABLE IF NOT EXISTS `hd_announce` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `starttime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '开始时间',
  `endtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `hits` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `cateid` (`status`,`endtime`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;