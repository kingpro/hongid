// Package: openapi.app.models.member
// File: openapi.app.models.member.go
// Created by mint
// Useage: 会员
// DATE: 14-7-8 14:51
package member

import (
	"api/app/models"
	"api/app/models/memberCard"
	"code.google.com/p/go-uuid/uuid"
	"github.com/revel/revel"
	"time"
	"utils/errors"
	"utils/page"
	"utils/times"
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
		Id: id,
	}
	if exist, err = models.ReaderEngine.GetByCol("Status", "<=", EMemberStatus_Del, member); err.IsError() {
		revel.WARN.Printf("GetMemberById[%v] return error: %v", id, err.ErrorMessage())
	}

	return
}

//新建加油站
func InsertNewGasStation(reqMsg *models.ReqNewStation) (rspMsg *models.RspNewStation, err errors.GlobalWaysError) {

	//TODO UUID VERSION 1
	uuid := uuid.NewUUID().String()
	timeNow := time.Now()

	//从Member表中抽取一个未使用的加油站记录
	member := &Member{
		GroupId: CMemberGroup_GasStation_ID,
	}
	if _, err = models.ReaderEngine.GetByCol("Status", "=", EMemberStatus_NotUse, member); err.IsError() {
		revel.WARN.Printf("Get unuse member record return error: %v", err.ErrorMessage())

		return nil, err
	}

	//更新信息
	memberUpd := &Member{
		UUID:     uuid,
		TelPhone: reqMsg.TelPhone,
		Email:    reqMsg.Email,
		NickName: reqMsg.StationName,
		Contact:  reqMsg.Contact,
		Phone:    reqMsg.Phone,
		Fax:      reqMsg.Fax,
		PostCode: reqMsg.PostCode,
		LastTime: timeNow.Format(times.Time_Layout_1),
		LastIp:   reqMsg.ClientIp,
		HomePage: reqMsg.HomePage,
		Status:   EMemberStatus_Enable,
	}

	affacted := false
	if affacted, err = models.WriterEngine.Update(memberUpd, member); err.IsError() {
		revel.WARN.Printf("Turn unuse member record[%+v] return error: %v", member, err.ErrorMessage())

		return nil, err
	}

	//如果成功更新会员主表信息，则添加一条详细信息
	if affacted {
		//加油站详细信息
		memberProfile := &MemberProfile{
			UUID:         uuid,
			ProvinceCode: reqMsg.ProvinceCode,
			CityCode:     reqMsg.CityCode,
			CountyCode:   reqMsg.CountyCode,
			AdressDetail: reqMsg.Address,
			RegTime:      timeNow.Format(times.Time_Layout_1),
			BirthDay:     timeNow.Format(times.Time_Layout_2),
			CalenDarType: ECalenDarType_Gregorian,
			RegIp:        reqMsg.ClientIp,
		}
		if _, err = models.WriterEngine.Insert(memberProfile); err.IsError() {
			revel.WARN.Printf("InsertMemberProfile return error: %v", err.ErrorMessage())

			return nil, err
		}
	}

	rspMsg = &models.RspNewStation{
		StationName:  reqMsg.StationName,
		Phone:        reqMsg.Phone,
		TelPhone:     reqMsg.TelPhone,
		Contact:      reqMsg.Contact,
		Fax:          reqMsg.Fax,
		Email:        reqMsg.Email,
		HomePage:     reqMsg.HomePage,
		PostCode:     reqMsg.PostCode,
		ProvinceCode: reqMsg.ProvinceCode,
		CityCode:     reqMsg.CityCode,
		CountyCode:   reqMsg.CountyCode,
		Address:      reqMsg.Address,
	}

	return
}
