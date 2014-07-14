// Package: openapi.app.models.member
// File: openapi.app.models.member.go
// Created by mint
// Useage: 会员
// DATE: 14-7-8 14:51
package member

import (
	"openapi/app/models/memberCard"
	"utils/errors"
	"openapi/app/models"
	"github.com/revel/revel"
	"utils/page"
)

type EMemberStatus byte

const (
	EMemberStatus_Enable  EMemberStatus = iota //正常，用户已经在使用
	EMemberStatus_Disable                      //异常，用户在使用，但被锁定
	EMemberStatus_Del                          //删除，已经被用户删除，但系统还未回收
	EMemberStatus_NotUse                       //初始，新生成 & 系统回收
	EMemberStatus_Liang                        //备用，系统生成后经过帐号过滤，被系统留着备用的，靓号...
)

type ECalenDarType byte

const (
	ECalenDarType_Gregorian ECalenDarType = iota //公历
	ECalenDarType_Lunar                          //农历
)

type EIdentifyCardType byte

const (
	EIdentifyCardType_IDCard   EIdentifyCardType = iota //身份证
	EIdentifyCardType_Passport                          //护照
)

type ESexType   byte
const (
	ESexType_Male     ESexType = iota    //男
	ESexType_Female
)

type Member struct {
	Id               int64
	HongId           int64  `xorm:"unique"`
	TelPhone         string `xorm:"char(11) unique"`
	Email            string `xorm:"varchar(255) unique"`
	GroupId          int64
	MemberCards      []memberCard.MemberCard
	NickName         string `xorm:"varchar(40) unique"`
	IdentifyCard     string `xorm:"varchar(18) unique"`
	IdentifyCardType EIdentifyCardType
	Avatar           string `xorm:"varchar(255)"`
	Sex              ESexType   `xorm:"tinyint(1)"`
	GrowthPoints     uint32 //成长点数
	GrowthLevelId    int64  //成长级别
	Language         string `xorm:"varchar(10)"`
	RegTime          string `xorm:"DateTime created"`
	RegIp            string `xorm:"char(15)"`
	LastTime         string `xorm:"DateTime"`
	LastIp           string `xorm:"char(15)"`
	HomePage         string `xorm:"unique"`
	Signature        string
	ValidTime        string `xorm:"DateTime"`
	Status           EMemberStatus `xorm:"tinyint(1)"`
	ProfileId        int64
}

type MemberProfile struct {
	Id           int64
	RealName     string `xorm:"varchar(40)"`
	Area         int64  // 地区表
	Adress       string
	Birthday     string `xorm:"Date"`
	CalenDarType ECalenDarType
	Company      string `xorm:"varchar(100)"`
	EduSchool    string
	EduLevel     byte
}

func NewDefaultMember() {

}

//根据会员组筛选会员
func FindMemberListByGroup(groupId int64, pager *page.Page) (memberList []*Member, exist bool, err errors.GlobalWaysError) {
	memberList = make([]*Member, 0)
	condiBean := &Member{
		GroupId: groupId,
		Status:  EMemberStatus_Enable,
	}

	if exist, err = models.ReaderEngine.FindByCol(&memberList, condiBean, pager); err.IsError() {
		revel.WARN.Printf("FindMemberListByGroup return err: %v", err.ErrorMessage())
	}

	return
}

//根据ID查找会员
func GetMemberById(id int64) (member *Member, exist bool, err errors.GlobalWaysError) {

	member = &Member{
		Id:    id,
	}
	if exist, err = models.ReaderEngine.GetByCol("Status", "<=", EMemberStatus_Del, member); err.IsError() {
		revel.WARN.Printf("GetMemberById[%v] return error: %v", id, err.ErrorMessage())
	}

	return
}
