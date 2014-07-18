// Package: models
// File: admin_panel.go
// Created by mint
// Useage: 快捷菜单
// DATE: 14-6-26 10:58
package models

import (
	"github.com/revel/revel"
	"time"
	"utils/times"
)

type AdminPanel struct {
	Id         int64  `xorm:"pk"`
	Mid        int64  `xorm:"int(11)"`
	Menu       *Menu  `xorm:"- <- ->"`
	Aid        int64  `xorm:"int(11)"`
	Name       string `xorm:"varchar(40)"`
	Url        string `xorm:"char(100)"`
	CreateTime string `xorm:"DateTime"`
}

//根据Id获取信息
func (a *AdminPanel) GetById(Id int64) *AdminPanel {

	admin_panel := new(AdminPanel)

	_, err := DB_Read.Id(Id).Get(admin_panel)
	if err != nil {
		revel.WARN.Printf("根据Id[%v]获取快捷菜单错误: %v", Id, err)
	}

	return admin_panel
}

//获取快捷方式列表
func (a *AdminPanel) GetPanelList(Admin_Info *Admin) []*AdminPanel {
	//初始化菜单
	admin_panel := []*AdminPanel{}

	err := DB_Read.Where("aid=?", Admin_Info.Id).Find(&admin_panel)

	if err != nil {
		revel.WARN.Printf("获取管理员[%v]快捷方式列表错误: %v", Admin_Info.RealName, err)
	} else {
		menu := new(Menu)

		for i, v := range admin_panel {
			admin_panel[i].Menu = menu.GetById(v.Mid)
		}
	}

	return admin_panel
}

//根据mid获取快捷方式
func (a *AdminPanel) GetByMid(Mid int64, Admin_Info *Admin) *AdminPanel {
	admin_panel := new(AdminPanel)

	_, err := DB_Read.Where("mid=? and aid=?", Mid, Admin_Info.Id).Get(admin_panel)
	if err != nil {
		revel.WARN.Printf("根据mid[%v]获取管理员[%v]快捷菜单错误: %v", Mid, Admin_Info.RealName, err)
	}

	return admin_panel
}

//是否已添加快捷方式
func (a *AdminPanel) IsAdd(Mid int64, Admin_Info *Admin) bool {
	admin_panel := new(AdminPanel)

	_, err := DB_Read.Where("mid=? AND aid=?", Mid, Admin_Info.Id).Get(admin_panel)
	if err != nil {
		revel.WARN.Printf("根据mid[%v]获取管理员[%v]快捷菜单错误: %v", Mid, Admin_Info.RealName, err)
	}

	if admin_panel.Id > 0 {
		return true
	}

	return false
}

//删除快捷方式
func (a *AdminPanel) DelPanel(Mid int64, Admin_Info *Admin) bool {
	admin_panel := new(AdminPanel)

	_, err := DB_Write.Where("mid=? AND aid=?", Mid, Admin_Info.Id).Delete(admin_panel)
	if err != nil {
		revel.WARN.Printf("删除管理员[%v]快捷方式错误: %v", Admin_Info.RealName, err)
		return false
	}

	return true
}

//添加快捷方式
func (a *AdminPanel) AddPanel(Mid int64, Admin_Info *Admin) bool {
	admin_panel := new(AdminPanel)

	admin_panel.Mid = Mid
	admin_panel.Aid = Admin_Info.Id

	menu := new(Menu)
	//获取菜单信息
	menu_info := menu.GetById(Mid)

	admin_panel.Name = menu_info.Name
	admin_panel.Url = menu_info.Url
	admin_panel.CreateTime = time.Now().Format(times.Time_Layout_1)

	_, err := DB_Write.Insert(admin_panel)
	if err != nil {
		revel.WARN.Printf("添加快捷方式错误: %v", err)
		return false
	}

	return true
}
