// Package: member
// File: member.go
// Created by mint
// Useage: 会员
// DATE: 14-7-18 17:33
package member

import (
	"memberCard"
	"github.com/go-xorm/xorm"

	"utils/errors"
	"utils/page"
)


type EMemberStatus byte

const (
	EMemberStatus_Enable  EMemberStatus = iota + 1 //正常，用户已经在使用
	EMemberStatus_Disable                          //异常，用户在使用，但被锁定
	EMemberStatus_Del                              //删除，已经被用户删除，但系统还未回收
	EMemberStatus_NotUse                           //初始，新生成 & 系统回收
	EMemberStatus_Liang                            //备用，系统生成后经过帐号过滤，被系统留着备用的，靓号...
)

type ECalenDarType byte

const (
	ECalenDarType_Gregorian ECalenDarType = iota + 1 //公历
	ECalenDarType_Lunar                              //农历
)

type EIdentifyCardType byte

const (
	EIdentifyCardType_IDCard   EIdentifyCardType = iota + 1 //身份证
	EIdentifyCardType_Passport                              //护照
)

type ESexType byte

const (
	ESexType_Male ESexType = iota + 1 //男
	ESexType_Female
)

type Member struct {
	Id               int64
	UUID             string `xorm:"unique"`
	HongId           string
	TelPhone         string `xorm:"char(11)"`
	Email            string `xorm:"varchar(255)"`
	GroupId          int64
	MemberCards      []memberCard.MemberCard
	NickName         string `xorm:"varchar(40)"`
	Contact          string
	IdentifyCard     string `xorm:"varchar(18)"`
	IdentifyCardType EIdentifyCardType
	Avatar           string   `xorm:"varchar(255)"`
	Sex              ESexType `xorm:"tinyint(1)"`
	Phone            string
	Fax              string
	PostCode         string
	GrowthPoints     uint32 //成长点数
	GrowthLevelId    int64  //成长级别
	Language         string `xorm:"varchar(10)"`
	LastTime         string `xorm:"DateTime"`
	LastIp           string `xorm:"char(15)"`
	HomePage         string
	Signature        string
	ValidTime        string        `xorm:"DateTime"`
	Status           EMemberStatus `xorm:"tinyint(1)"`
}

type MemberProfile struct {
	Id           int64
	UUID         string `xorm:"unique"`
	ProvinceCode string
	CityCode     string
	CountyCode   string
	AdressDetail string
	BirthDay     string `xorm:"Date"`
	CalenDarType ECalenDarType
	Company      string `xorm:"varchar(100)"`
	EduSchool    string
	EduLevel     byte
	RegTime      string `xorm:"DateTime created"`
	RegIp        string `xorm:"char(15)"`
}


//根据会员组筛选会员
func FindMemberListByGroup(groupId int64, pager *page.Page, engine *xorm.Engine) ([]*Member, bool, errors.GlobalWaysError) {
	memberList := make([]*Member, 0)
	condiBean := &Member{
		GroupId: groupId,
	}

	if err := engine.Limit(int(pager.Perpage), int((pager.Current_page-1)*pager.Perpage)).Where("Status < ?", EMemberStatus_Del).Find(&memberList, condiBean); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_FIND, "Find return error: %v", err)
	} else if len(memberList) == 0 {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return memberList, true, errors.ErrorOK()
}

//根据ID查找会员
func GetMemberById(id int64, engine *xorm.Engine) (*Member, bool, errors.GlobalWaysError) {

	member := &Member{
		Id: id,
	}

	if exist, err := engine.Where("Status < ?", EMemberStatus_Del).Get(member); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_GET, "Get return error: %v", err)
	} else if !exist {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return member, true, errors.ErrorOK()
}

//根据ID删除会员(更新状态)
func DelMemberById(id int64, engine *xorm.Engine) (bool, errors.GlobalWaysError) {

	member := &Member{
		Id:   id,
		Status: EMemberStatus_Del,
	}

	if _, err := engine.Update(member); err != nil {
		return false, errors.Newf(errors.CODE_DB_ERR_UPDATE, "Del member return error: %v", err)
	}

	return true, errors.ErrorOK()
}
