// Package: models
// File: mail.go
// Created by mint
// Useage: 邮件模型
// DATE: 14-7-6 13:28
package models

/*
CREATE TABLE IF NOT EXISTS `mailqueue` (
  `mailid` int(11) NOT NULL AUTO_INCREMENT,
  `toid` int(11) NOT NULL,
  `tomail` varchar(255) NOT NULL,
  `frommail` varchar(255) NOT NULL,
  `subject` varchar(255) NOT NULL,
  `message` text NOT NULL,
  `charset` varchar(255) NOT NULL,
  `htmlon` tinyint(4) NOT NULL,
  `level` tinyint(4) NOT NULL,
  `dateline` bigint(20) NOT NULL,
  `failures` int(11) NOT NULL,
  PRIMARY KEY (`mailid`),
  KEY `toid` (`toid`),
) ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;
 */

type MailQueue struct {
	MailId   int64 `xorm:"pk autoincr"`
	ToId     string `xorm:"notnull char(19) index"`
	ToMail   string `xorm:"varchar(255) index"`
	FromMail string `xorm:"varchar(255) index"`
	Subject  string `xorm:"varchar(255"`
	Message  string `xorm:"text notnull"`
	Htmlon   bool
	Level    uint8
	Status   bool
	DateLine string `xorm:"DateTime"`
}

