// Package: models
// File: openapi.app.models.member.go
// Created by mint
// Useage: 会员
// DATE: 14-7-7 13:49
package models
/*
import (
	"github.com/revel/revel"
	"github.com/go-xorm/xorm"
	"html/template"
	"admin/utils"
	"utils/page"
	"openapi/app/models/memberCard"
)

var (
	_ *xorm.Engine
)

type Member struct {
	Id           int64
	HongId       int64
	TelPhone     string `xorm:"char(11)"`
	Email        string `xorm:"varchar(255)"`
	GroupId      int64
	MemberCards  []memberCard.MemberCard
	NickName     string `xorm:"varchar(40) unique"`
	RealName     string `xorm:"varchar(40)"`
	Avatar       string `xorm:"varchar(255)"`
	Sex          byte   `xorm:"tinyint(1)"`
	Language     string `xorm:"varchar(10)"`
	RegTime      string `xorm:"DateTime"`
	RegIp        string `xorm:"char(15)"`
	LastTime     string `xorm:"DateTime"`
	LastIp       string `xorm:"char(15)"`
	HomePage     string
	Signature    string
	IsLock       byte   `xorm:"tinyint(1)"`
	Status       byte
	ProfileId    int64
}

type MemberProfile struct {
	Id           int64
	Area         int64
	Birthday     string `xorm:"Date"`
	CalenDarType byte
	Company      string `xorm:"varchar(100)"`
	EduSchool    string
	EduLevel     byte
}

//根据Id获取信息
func (m *Member) GetById(id int64) *Member {

	member := new(Member)
	_, err := DB_Read.Id(id).Get(member)
	if err != nil {
		revel.WARN.Printf("根据ID[%v]获取会员信息错误: %v", id, err)
	}

	return member
}

//根据HongId获取信息
func (m *Member) GetByHongId(hongId int64) *Member {

	member := new(Member)
	_, err := DB_Read.Where("hong_id = ?", hongId).Get(member)
	if err != nil {
		revel.WARN.Printf("根据HongID[%v]获取会员信息错误: %v", hongId, err)
	}

	return member
}

//根据telphone获取信息
func (m *Member) GetByTel(tel string) *Member {
	member := new(Member)
	_, err :=DB_Read.Where("tel_phone = ?", tel).Get(member)
	if err != nil {
		revel.WARN.Printf("根据TelPhone[%v]获取信息错误: %v", tel, err)
	}

	return member
}

//根据Email获取信息
func (m *Member) GetByEmail(mail string) *Member {
	member := new(Member)
	_, err :=DB_Read.Where("email = ?", mail).Get(member)
	if err != nil {
		revel.WARN.Printf("根据Email[%v]获取信息错误: %v", mail, err)
	}

	return member
}

//用户名是否已有
func (m *Member) HasName() bool {

	member := new(Member)
	_, err := DB_Read.Where("nick_name=?", m.NickName).Get(member)
	if err != nil {
		revel.WARN.Printf("获取用户[%v]信息错误: %v",m.Id, err)
		return false
	}

	if member.Id > 0 && member.Id != m.Id {
		return true
	}

	return false
}

//邮箱是否已有
func (m *Member) HasEmail() bool {

	member := new(Member)
	_, err := DB_Read.Where("email=?", m.Email).Get(member)
	if err != nil {
		revel.WARN.Printf("获取用户[%v]信息错误: %v", m.Id, err)
		return false
	}

	if member.Id > 0 && member.Id != m.Id {
		return true
	}

	return false
}

//获取会员列表
func (m *Member) GetUserList(search string, Page int64, Perpage int64) (member_arr []*Member, html template.HTML, where map[string]string) {
	//初始化菜单
	member_list := []*Member{}

	//查询条件
	var WhereStr string = " 1 AND "

	if len(search) > 0 {
		//解码
		where = utils.DecodeSegment(search)

		// 注册时间区间
		if len(where["start_time"]) > 0 {
			WhereStr += " `reg_time` >='" + where["start_time"] + " 00:00:00' AND "
		}
		if len(where["end_time"]) > 0 {
			WhereStr += " `reg_time` <='" + where["end_time"] + " 23:59:59' AND "
		}
		//
		if len(where["islock"]) > 0 && where["islock"] != "0" {
			WhereStr += " `is_lock` =" + where["islock"]
		}

		if len(where["type"]) > 0 && len(where["keyword"]) > 0 {

			if where["type"] == "1" {
				//用户名
				WhereStr += " `hong_id` ='" + where["keyword"] + "' AND "
			} else if where["type"] == "2" {
				//用户ID
				WhereStr += " `id` =" + where["keyword"] + " AND "
			} else if where["type"] == "3" {
				//邮箱
				WhereStr += " `email` ='" + where["keyword"] + "' AND "
			} else if where["type"] == "4" {
				//注册ip
				WhereStr += " `reg_ip` ='" + where["keyword"] + "' AND "
			} else if where["type"] == "5" {
				//昵称
				WhereStr += " `nick_name` like '%" + where["keyword"] + "%' AND "
			}
		}
	}

	WhereStr += " 1 "

	//查询总数
	member := new(Member)
	Total, err := DB_Read.Where(WhereStr).Count(member)
	if err != nil {
		revel.WARN.Printf("查询会员总数错误: %v", err)
	}

	//分页
	Pager := new(page.Page)
	if len(search) > 0 {
		Pager.SubPage_link = "/Member/" + search + "/"
	} else {
		Pager.SubPage_link = "/Member/"
	}

	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	DB_Read.Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&member_list)

	if len(member_list) > 0 {
//		user_group := new(User_Group)
//		for i, v := range user_list {
//			user_list[i].UserGroup = user_group.GetById(v.Groupid)
//		}
	}

	return member_list, pages, where
}


//移动
func (m *Member) Move(groupid int64, ids string) bool {
	member := new(Member)
	member.Groupid = groupid

	_, err := DB_Write.Where("id in (" + ids + ")").Cols("group_id").Update(member)
	if err != nil {
		revel.WARN.Printf("移动会员错误: %v", err)
		return false
	}

	return true
}

//解锁
func (m *Member) Unlock(ids string) bool {
	member := new(Member)
	member.Islock = 2

	_, err := DB_Write.Where("id in (" + ids + ")").Cols("is_lock").Update(member)
	if err != nil {
		revel.WARN.Printf("解锁会员错误: %v", err)
		return false
	}

	return true
}

//锁定
func (m *Member) Lock(ids string) bool {
	member := new(Member)
	member.Islock = 1

	_, err := DB_Write.Where("id in (" + ids + ")").Cols("is_lock").Update(member)
	if err != nil {
		revel.WARN.Printf("锁定会员错误: %v", err)
		return false
	}

	return true
}

//批量删除
//TODO 删除只是将status置为无效
func (m *Member) DelByIDS(ids string) bool {
	member := new(Member)
	_, err := DB_Write.Where("id in (" + ids + ")").Delete(member)
	if err != nil {
		revel.WARN.Printf("批量删除会员错误: %v", err)
		return false
	}

	return true
}

//删除用户
func (m *Member) DelByID(Id int64) bool {

	member := new(Member)
	_, err := DB_Write.Id(Id).Delete(member)
	if err != nil {
		revel.WARN.Printf("删除会员错误: %v", err)
		return false
	}

	return true
}
*/
