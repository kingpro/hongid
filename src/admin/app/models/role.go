// Package: models
// File: role.go
// Created by mint
// Useage: 角色表
// DATE: 14-6-25 20:03
package models

//角色表
import (
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"html/template"
	"time"
	"utils/page"
	"utils/times"
)

var _ *xorm.Engine

type Role struct {
	Id         int64
	RoleName   string `xorm:"unique varchar(255)"`
	Desc       string `xorm:"varchar:(255)"`
	Data       string `xorm:"text"` //菜单列表
	Status     int64  `xorm:"bool"`
	CreateTime string `xorm:"DateTime"`
}

//根据ID获取角色信息
func (r *Role) GetById(Id int64) *Role {

	role := new(Role)
	_, err := DB_Read.Id(Id).Get(role)
	if err != nil {
		revel.WARN.Printf("根据Id[%v]获取角色信息错误: %v", Id, err)
	}

	return role
}

//获取角色列表
func (r *Role) GetByAll(Page int64, Perpage int64) ([]*Role, template.HTML) {
	role_list := []*Role{}

	//查询总数
	role := new(Role)
	Total, err := DB_Read.Count(role)
	if err != nil {
		revel.WARN.Printf("获取角色总数错误: %v", err)
	}

	// 分页
	Pager := new(page.Page)
	Pager.SubPage_link = "/Role/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Find(&role_list)

	return role_list, pages
}

//获取角色
func (r *Role) GetRoleList() []*Role {
	roleList := []*Role{}
	DB_Read.Find(&roleList)
	return roleList
}

//添加角色
func (r *Role) Save() bool {

	role := new(Role)
	role.RoleName = r.RoleName
	role.Desc = r.Desc
	role.Data = r.Data
	role.Status = r.Status
	role.CreateTime = time.Now().Format(times.Time_Layout_1)

	_, err := DB_Write.Insert(role)
	if err != nil {
		revel.WARN.Printf("添加角色错误: %v", err)
		return false
	}

	return true
}

//编辑角色
func (r *Role) Edit(Id int64) bool {
	role := new(Role)

	if len(r.RoleName) > 0 {
		role.RoleName = r.RoleName
	}
	if len(r.Desc) > 0 {
		role.Desc = r.Desc
	}
	if len(r.Data) > 0 {
		role.Data = r.Data
	}
	role.Status = r.Status

	_, err := DB_Write.Id(Id).Cols("rolename", "desc", "data", "status").Update(role)
	if err != nil {
		revel.WARN.Printf("编辑角色错误: %v", err)
		return false
	}

	return true
}

//设置状态
func (r *Role) SetStatus(Id int64) bool {
	role := new(Role)

	role.Status = r.Status
	_, err := DB_Write.Id(Id).Cols("status").Update(role)
	if err != nil {
		revel.WARN.Printf("设置角色状态错误: %v", err)
		return false
	}

	return true
}

//删除角色
func (r *Role) DelByID(Id int64) bool {

	_, err := DB_Write.Id(Id).Delete(new(Role))
	if err != nil {
		revel.WARN.Printf("删除角色错误: %v", err)
		return false
	}

	return true
}
