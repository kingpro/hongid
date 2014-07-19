// Package: member
// File: member.go
// Created by mint
// Useage: 会员
// DATE: 14-7-18 17:33
package member

import (
	"github.com/go-xorm/xorm"
	"memberCard"
	"utils/algorith"
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

type EEduLevel byte

const (
	EEduLevel_None      EEduLevel = iota + 1 //文盲
	EEduLevel_Primary                        //小学
	EEduLevel_Secondary                      //中学
	EEduLevel_Senior                         //高中
	EEduLevel_Bachelor                       //本科
	EEduLevel_Master                         //研究生
	EEduLevel_Doctor                         //博士
)

type Member struct {
	Id               int64
	UUID             string `xorm:"default ''"`
	HongId           int64
	TelPhone         string `xorm:"char(11) default ''"`
	Email            string `xorm:"default ''"`
	CryptPwd         string `xorm:"default ''"`
	GroupId          int64  `xorm:"default '0'"`
	MemberCards      []memberCard.MemberCard
	NickName         string            `xorm:"default ''"`
	Contact          string            `xorm:"default ''"`
	IdentifyCard     string            `xorm:"varchar(18) default ''"`
	IdentifyCardType EIdentifyCardType `xorm:"tinyint(1) default '1'"`
	Avatar           string            `xorm:"varchar(255) default ''"`
	Sex              ESexType          `xorm:"tinyint(1) default '1'"`
	Phone            string            `xorm:"char(12) default ''"`
	Fax              string            `xorm:"char(12) default ''"`
	PostCode         string            `xorm:"char(6) default ''"`
	GrowthPoints     uint32            `xorm:"default '0'"` //成长点数
	GrowthLevelId    int64             `xorm:"default '1'"` //成长级别
	Language         string            `xorm:"varchar(10) default ''"`
	LastTime         string            `xorm:"DateTime default '0000-00-00'"`
	LastIp           string            `xorm:"char(15) default '000.000.000.000'"`
	HomePage         string            `xorm:"default ''"`
	Signature        string            `xorm:"default ''"`
	ValidTime        string            `xorm:"DateTime default '0000-00-00'"`
	Status           EMemberStatus     `xorm:"tinyint(1) default '4'"`
}

type MemberProfile struct {
	Id           int64
	UUID         string        `xorm:"unique"`
	ProvinceCode string        `xorm:"char(6) default ''"`
	CityCode     string        `xorm:"char(6) default ''"`
	CountyCode   string        `xorm:"char(6) default ''"`
	AdressDetail string        `xorm:"default ''"`
	BirthDay     string        `xorm:"Date default '0000-00-00'"`
	CalenDarType ECalenDarType `xorm:"tinyint(1) default '1'"`
	Company      string
	EduSchool    string
	EduLevel     EEduLevel `xorm:"tinyint(1) default '1'"`
	RegTime      string    `xorm:"DateTime created"`
	RegIp        string    `xorm:"char(15) default '000.000.000.000'"`
}

//查找
//======================================================================================================================
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

//根据Hongid查找会员
func GetMemberByHongId(hongId int64, engine *xorm.Engine) (*Member, bool, errors.GlobalWaysError) {
	member := &Member{
		HongId: hongId,
	}

	if exist, err := engine.Where("Status < ?", EMemberStatus_Del).Get(member); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_GET, "Get return error: %v", err)
	} else if !exist {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return member, true, errors.ErrorOK()
}

//根据TelPhone查找会员
func GetMemberByTel(tel string, engine *xorm.Engine) (*Member, bool, errors.GlobalWaysError) {
	member := &Member{
		TelPhone: tel,
	}

	if exist, err := engine.Where("Status < ?", EMemberStatus_Del).Get(member); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_GET, "Get return error: %v", err)
	} else if !exist {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return member, true, errors.ErrorOK()
}

//根据Email查找会员
func GetMemberByEmail(email string, engine *xorm.Engine) (*Member, bool, errors.GlobalWaysError) {
	member := &Member{
		Email: email,
	}

	if exist, err := engine.Where("Status < ?", EMemberStatus_Del).Get(member); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_GET, "Get return error: %v", err)
	} else if !exist {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return member, true, errors.ErrorOK()
}

//根据NickName查找会员
func GetMemberByNickName(nickName string, engine *xorm.Engine) (*Member, bool, errors.GlobalWaysError) {
	member := &Member{
		NickName: nickName,
	}

	if exist, err := engine.Where("Status < ?", EMemberStatus_Del).Get(member); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_GET, "Get return error: %v", err)
	} else if !exist {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return member, true, errors.ErrorOK()
}

//返回一个未使用的会员
func GetUnUseMember(groupId int64, engine *xorm.Engine) (*Member, bool, errors.GlobalWaysError) {
	member := &Member{
		GroupId: groupId,
		Status:  EMemberStatus_NotUse,
	}

	if exist, err := engine.Get(member); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_GET, "Get return error: %v", err)
	} else if !exist {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return member, true, errors.ErrorOK()
}

//======================================================================================================================

//======================================================================================================================
//删除会员(更新状态)
//将其状态改为 用户删除 态。过一定时间未重新启用，则被系统收回。期间用户个人信息不被清除
func DelMember(member *Member, engine *xorm.Engine) (bool, errors.GlobalWaysError) {

	member.Status = EMemberStatus_Del

	if _, err := engine.Id(member.Id).Update(member); err != nil {
		return false, errors.Newf(errors.CODE_DB_ERR_UPDATE, "Del member return error: %v", err)
	}

	return true, errors.ErrorOK()
}

//系统回收（定期执行）
//会员状态: 用户删除 ---> 未使用
//会员详细信息: 删除
func RecoverMember(engine *xorm.Engine) (affactedTotal int64, idList []int64) {
	affactedTotal = 0
	idList = make([]int64, 0)

	memberList := make([]*Member, 0)
	if err := engine.Where("Status = ?", EMemberStatus_Del).Find(&memberList); err != nil {
		return
	}

	for _, member := range memberList {
		session := engine.NewSession()
		defer session.Close()

		session.Begin()
		{
			member.Status = EMemberStatus_NotUse

			var affacted int64 = 0
			var err error

			if affacted, err = engine.Id(member.Id).Update(member); err != nil {
				session.Rollback()
				continue
			}

			if _, err = engine.Where("UUID = ?", member.UUID).Delete(new(MemberProfile)); err != nil {
				session.Rollback()
				continue
			}

			affactedTotal += affacted
			idList = append(idList, member.Id)
		}
		session.Commit()
	}

	return
}

//======================================================================================================================

//======================================================================================================================
//生成会员
func GenMember(member *Member, memberProfile *MemberProfile, engine *xorm.Engine) (int64, errors.GlobalWaysError) {

	session := engine.NewSession()
	defer session.Close()

	session.Begin()
	{
		if _, err := engine.Insert(member); err != nil {
			session.Rollback()
			return 0, errors.Newf(errors.CODE_DB_ERR_INSERT, "Generator Member Return Error: %v", err)
		}

		if memberProfile != nil {
			if _, err := engine.Insert(memberProfile); err != nil {
				session.Rollback()
				return 0, errors.Newf(errors.CODE_DB_ERR_INSERT, "Generator MemberProfile Return Error: %v", err)
			}
		}
	}
	err := session.Commit()
	if err != nil {
		session.Rollback()
		return 0, errors.New(errors.CODE_DB_ERR_COMMIT, errors.MSG_DB_ERR_COMMIT)
	}

	return 1, errors.ErrorOK()
}

//批量生成会员
func GenMembers(minNo int64, maxNo int64, count int64, groupId int64, engine *xorm.Engine) (affactedTotal int64) {
	affactedTotal = 0

	for affactedTotal <= count {
		randId := algorith.RandomInt64(minNo, maxNo)
		//如果存在当前hongid，放弃randId
		if has, err := engine.Where("HongId = ?", randId).Get(new(Member)); has || err != nil {
			continue
		}

		//TODO 靓号逻辑判断
		member := &Member{
			HongId:  randId,
			GroupId: groupId,
			Status:  EMemberStatus_NotUse,
		}

		if affacted, err := GenMember(member, nil, engine); err.IsError() {
			continue
		} else {
			affactedTotal += affacted
		}

		//loop break logic
		if cnt, _ := engine.Where("HongId >= ?", minNo).And("HongId <= ?", maxNo).Count(new(Member)); cnt >= count {
			break
		}
	}

	return
}

//======================================================================================================================
//更新会员
func UpdaeteMember(member *Member, memberProfile *MemberProfile, engine *xorm.Engine) (bool, errors.GlobalWaysError) {
	session := engine.NewSession()
	defer session.Close()

	session.Begin()
	{
		if _, err := engine.Id(member.Id).Update(member); err != nil {
			session.Rollback()
			return false, errors.Newf(errors.CODE_DB_ERR_UPDATE, "Update Member Return Error: %v", err)
		}

		if memberProfile != nil {
			if _, err := engine.Where("UUID = ?", member.UUID).Update(memberProfile); err != nil {
				session.Rollback()
				return false, errors.Newf(errors.CODE_DB_ERR_INSERT, "Update MemberProfile Return Error: %v", err)
			}
		}
	}
	err := session.Commit()
	if err != nil {
		session.Rollback()
		return false, errors.New(errors.CODE_DB_ERR_COMMIT, errors.MSG_DB_ERR_COMMIT)
	}

	return true, errors.ErrorOK()
}
