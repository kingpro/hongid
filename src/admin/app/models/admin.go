// Package: models
// File: admin.go
// Created by mint
// Useage: 管理员表
// DATE: 14-6-25 19:54
package models

import (
	"github.com/revel/revel"
	"html/template"
	"regexp"
	"time"
	"utils/ipv4"
	"utils/page"
	"utils/security"
	"utils/times"
)

type Admin struct {
	Id            int64  `xorm:"pk autoincr"`
	Username      string `xorm:"unique index varchar(255)"`
	Password      string `xorm:"varchar:(255)"`
	Roleid        int64  `xorm:"index"`
	Role          *Role  `xorm:"- <- ->"`
	Lastloginip   string `xorm:"varchar(32)"`
	Lastlogintime string `xorm:"varchar(32)"`
	Email         string `xorm:"varchar(32)"`
	Realname      string `xorm:"varchar(32)"`
	Lang          string `xorm:"varchar(6)"`
	Status        int64  `xorm:"bool"`
	Createtime    string `xorm:"DateTime"`
}

type Password struct {
	Password        string
	PasswordConfirm string
}

func (a *Admin) Validate(v *revel.Validation) {
	v.Required(a.Username).Message("请输入用户名!")
	valid := v.Match(a.Username, regexp.MustCompile("^\\w*$")).Message("只能使用字母、数字和下划线!")
	if valid.Ok {
		if a.HasName() {
			err := &revel.ValidationError{
				Message: "该用户名已经注册过!",
				Key:     "a.Username",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(a.Email).Message("请输入Email")
	valid = v.Email(a.Email).Message("无效的电子邮件!")
	if valid.Ok {
		if a.HasEmail() {
			err := &revel.ValidationError{
				Message: "该邮件已经注册过!",
				Key:     "a.Email",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(a.Password).Message("请输入密码!")
	v.MinSize(a.Password, 6).Message("密码最少六位!")
}

//验证密码
func (P *Password) ValidatePassword(v *revel.Validation) {
	v.Required(P.Password).Message("请输入密码!")
	v.Required(P.PasswordConfirm).Message("请输入确认密码!")

	v.MinSize(P.Password, 6).Message("密码最少六位!")
	v.Required(P.Password == P.PasswordConfirm).Message("两次密码不相同!")
}

//获取管理员列表
func (a *Admin) GetByAll(RoleId int64, Page int64, Perpage int64) ([]*Admin, template.HTML) {
	admin_list := []*Admin{}

	if RoleId > 0 {

		//查询总数
		admin := new(Admin)
		Total, err := DB_Read.Where("roleid=?", RoleId).Count(admin)
		if err != nil {
			revel.WARN.Printf("获取角色[%v]的管理员列表错误: %v", RoleId, err)
		}

		//分页
		Pager := new(page.Page)
		Pager.SubPage_link = "/Admin/"
		Pager.Nums = Total
		Pager.Perpage = Perpage
		Pager.Current_page = Page
		Pager.SubPage_type = 2
		pages := Pager.Show()

		DB_Read.Where("roleid=?", RoleId).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Find(&admin_list)

		if len(admin_list) > 0 {
			role := new(Role)

			for i, v := range admin_list {
				admin_list[i].Role = role.GetById(v.Roleid)
			}
		}

		return admin_list, pages
	} else {

		//查询总数
		admin := new(Admin)
		Total, err := DB_Read.Count(admin)
		if err != nil {
			revel.WARN.Printf("获取全部管理员列表错误: %v", err)
		}

		//分页
		Pager := new(page.Page)
		Pager.SubPage_link = "/Admin/"
		Pager.Nums = Total
		Pager.Perpage = Perpage
		Pager.Current_page = Page
		Pager.SubPage_type = 2
		pages := Pager.Show()

		DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Find(&admin_list)

		if len(admin_list) > 0 {
			role := new(Role)

			for i, v := range admin_list {
				admin_list[i].Role = role.GetById(v.Roleid)
			}
		}

		return admin_list, pages
	}
}

//验证用户名是否已经注册过
func (a *Admin) HasName() bool {

	admin := new(Admin)
	_, err := DB_Read.Where("username=?", a.Username).Get(admin)
	if err != nil {
		revel.WARN.Printf("验证管理员用户名时错误: %v", err)
		return false
	}

	if admin.Id > 0 && admin.Id != a.Id {
		return true
	}

	return false
}

//验证邮箱是否已经注册过
func (a *Admin) HasEmail() bool {

	admin := new(Admin)
	_, err := DB_Read.Where("email=?", a.Email).Get(admin)
	if err != nil {
		revel.WARN.Printf("验证管理员Email错误: %v", err)
		return false
	}

	if admin.Id > 0 && admin.Id != a.Id {
		return true
	}

	return false
}

//根据Id获取管理员信息
func (a *Admin) GetById(Id int64) *Admin {
	admin := new(Admin)

	_, err := DB_Read.Id(Id).Get(admin)
	if err != nil {
		revel.WARN.Printf("根据Id[%v]获取管理员信息错误: %v", Id, err)
	} else {
		role := new(Role)
		admin.Role = role.GetById(admin.Roleid)
	}

	return admin
}

//根据真实姓名获取管理员信息
func (a *Admin) GetByRealName(name string) *Admin {
	admin := new(Admin)

	_, err := DB_Read.Where("realname=?", name).Get(admin)
	if err != nil {
		revel.WARN.Printf("根据真实姓名[%v]获取管理员信息错误: %v", name, err)
	} else {
		role := new(Role)
		admin.Role = role.GetById(admin.Roleid)
	}

	return admin
}

//根据用户名获取管理员信息
func (a *Admin) GetByName(name string) *Admin {
	admin := new(Admin)

	_, err := DB_Read.Where("username=?", name).Get(admin)
	if err != nil {
		revel.WARN.Printf("根据用户名[%v]获取管理员信息错误: %v", name, err)
	} else {
		role := new(Role)
		admin.Role = role.GetById(admin.Roleid)
	}

	return admin
}

//添加管理员
func (a *Admin) Save() bool {

	admin := new(Admin)
	admin.Username = a.Username
	admin.Password = security.GenerateFromPassword(a.Password)
	admin.Roleid = a.Roleid

	var err error
	admin.Lastloginip, err = ipv4.GetClientIP()
	if err != nil {
		revel.WARN.Printf("获取客户端IP错误: %v", err)
	}

	admin.Email = a.Email
	admin.Realname = a.Realname
	admin.Lang = a.Lang
	admin.Lastlogintime = "0000-00-00 00:00:00"
	admin.Status = a.Status
	admin.Createtime = time.Now().Format(times.Time_Layout_1)

	_, err = DB_Write.Insert(admin)
	if err != nil {
		revel.WARN.Printf("添加管理员错误: %v", err)
		return false
	}

	return true
}

//更新登录IP&时间
func (a *Admin) UpdateLoginTime(Id int64) bool {
	admin := new(Admin)

	admin.Lastloginip, _ = ipv4.GetClientIP()
	admin.Lastlogintime = time.Now().Format(times.Time_Layout_1)

	_, err := DB_Write.Id(Id).Cols("lastloginip", "lastlogintime").Update(admin)
	if err != nil {
		revel.WARN.Printf("更新管理员登录IP&时间错误: %v", err)
		return false
	}

	return true
}

//修改个人信息
func (a *Admin) EditInfo(Id int64) bool {
	admin := new(Admin)

	if len(a.Email) > 0 {
		admin.Email = a.Email
	}

	if len(a.Realname) > 0 {
		admin.Realname = a.Realname
	}

	if len(a.Lang) > 0 {
		admin.Lang = a.Lang
	}

	_, err := DB_Write.Id(Id).Cols("email", "realname", "lang").Update(admin)
	if err != nil {
		revel.WARN.Printf("更新管理员[%v]个人信息错误: %v", a.Realname, err)
		return false
	}

	return true
}

//修改密码
func (a *Admin) EditPwd(Id int64) bool {
	admin := new(Admin)

	if len(a.Password) > 0 {
		admin.Password = security.GenerateFromPassword(a.Password)
	}

	_, err := DB_Write.Id(Id).Cols("password").Update(admin)
	if err != nil {
		revel.WARN.Printf("修改管理员密码错误: %v", err)
		return false
	}

	return true
}

//编辑管理员
func (a *Admin) Edit(Id int64) bool {

	admin := new(Admin)

	if len(a.Username) > 0 {
		admin.Username = a.Username
	}

	if len(a.Password) > 0 {
		admin.Password = security.GenerateFromPassword(a.Password)
	}

	if a.Roleid > 0 {
		admin.Roleid = a.Roleid
	}

	if len(a.Email) > 0 {
		admin.Email = a.Email
	}

	if len(a.Realname) > 0 {
		admin.Realname = a.Realname
	}

	if len(a.Lang) > 0 {
		admin.Lang = a.Lang
	}

	admin.Status = a.Status

	if len(a.Password) > 0 {
		_, err := DB_Write.Id(Id).Cols("username", "password", "email", "realname", "roleid", "lang", "status").Update(admin)
		if err != nil {
			revel.WARN.Printf("编辑管理员错误: %v", err)
			return false
		}
	} else {
		_, err := DB_Write.Id(Id).Cols("username", "email", "realname", "roleid", "lang", "status").Update(admin)
		if err != nil {
			revel.WARN.Printf("编辑管理员错误: %v", err)
			return false
		}
	}

	return true
}

//删除管理员
func (a *Admin) DelByID(Id int64) bool {

	admin := new(Admin)

	_, err := DB_Write.Id(Id).Delete(admin)
	if err != nil {
		revel.WARN.Printf("删除管理员错误: %v", err)
		return false
	}
	return true
}

//获取MySQL版本
func (a *Admin) GetMysqlVer() string {
	sql := "SELECT VERSION() AS version;"
	results, err := DB_Read.Query(sql)

	if err != nil {
		revel.WARN.Printf("获取MySQL版本错误: %v", err)
	}

	return string(results[0]["version"])
}
