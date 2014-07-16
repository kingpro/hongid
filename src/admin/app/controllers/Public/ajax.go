// Package: Public
// File: ajax.go
// Created by mint
// Useage: AJAX操作
// DATE: 14-6-28 14:17
package Public

import (
	"admin/app/models"
	"admin/utils"
	"admin/utils/consts"
	"github.com/dchest/captcha"
	"github.com/revel/revel"
	"utils/security"
)

type Ajax struct {
	*revel.Controller
}

//获取验证码
func (c *Ajax) GetCaptcha() revel.Result {
	CaptchaId := captcha.NewLen(6)
	return c.RenderText(CaptchaId)
}

//锁屏
func (c *Ajax) ScreenLock() revel.Result {
	data := make(map[string]string)

	c.Session[consts.C_Session_LockS] = consts.C_Lock_1

	data["status"] = consts.C_Lock_1
	data["message"] = "锁屏!"
	return c.RenderJson(data)
}

//解锁
func (c *Ajax) ScreenUnlock(admin *models.Admin) revel.Result {
	lock_password := c.Params.Get("lock_password")

	if lock_password == "" || len(lock_password) <= 0 {
		return c.RenderText(consts.C_UnLock_2)
	}

	if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {

		admin_info := admin.GetById(adminId)

		if !security.CompareHashAndPassword(admin_info.Password, lock_password) {
			return c.RenderText(consts.C_UnLock_3)
		} else {
			c.Session[consts.C_Session_LockS] = consts.C_Lock_0
			return c.RenderText(consts.C_UnLock_1)
		}
	}

	return c.RenderText(consts.C_UnLock_4)
}

//当前位置
func (c *Ajax) Pos(menu *models.Menu) revel.Result {
	id := c.Params.Get("id")

	if menuId, ok := utils.ParseMenuId(id); ok {
		if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {
			//获取登陆用户信息
			admin := new(models.Admin)
			admin_info := admin.GetById(adminId)

			menu_str := menu.GetPos(menuId, admin_info)

			return c.RenderText(menu_str)
		}
	}

	return c.RenderText("")
}

//检查消息
func (c *Ajax) GetMessage() revel.Result {
	data := make(map[string]string)

	data["status"] = "0"
	data["message"] = "请填写用户名!"
	return c.RenderJson(data)
}

//获取快捷方式
func (c *Ajax) GetPanel(admin_panel *models.AdminPanel) revel.Result {

	if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {

		mid := c.Params.Get("mid")
		if menuId, ok := utils.ParseMenuId(mid); ok {

			//获取登陆用户信息
			admin := new(models.Admin)
			admin_info := admin.GetById(adminId)

			panel_info := admin_panel.GetByMid(menuId, admin_info)
			if panel_info.Id > 0 {
				Html := "<span><a target='right' href='/" + panel_info.Url + "/'>" + panel_info.Name + "</a><a class='panel-delete' href='javascript:delete_panel();'></a></span>"
				return c.RenderText(Html)
			} else {
				Html := ""
				return c.RenderText(Html)
			}
		}
	}

	Html := "<span><a href='javascript:;'>未登陆</a></span>"
	return c.RenderText(Html)
}

//添加快捷方式
func (c *Ajax) AddPanel(admin_panel *models.AdminPanel) revel.Result {

	if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {

		mid := c.Params.Get("mid")
		if menuId, ok := utils.ParseMenuId(mid); ok {

			//获取登陆用户信息
			admin := new(models.Admin)
			admin_info := admin.GetById(adminId)

			//是否已添加快捷方式
			isAdd := admin_panel.IsAdd(menuId, admin_info)
			if isAdd {
				panel_info := admin_panel.GetByMid(menuId, admin_info)

				Html := "<span><a target='right' href='/" + panel_info.Url + "/'>" + panel_info.Name + "</a><a class='panel-delete' href='javascript:delete_panel();'></a></span>"
				return c.RenderText(Html)
			} else {
				isFinish := admin_panel.AddPanel(menuId, admin_info)

				if isFinish {
					panel_info := admin_panel.GetByMid(menuId, admin_info)

					Html := "<span><a target='right' href='/" + panel_info.Url + "/'>" + panel_info.Name + "</a><a class='panel-delete' href='javascript:delete_panel();'></a></span>"
					return c.RenderText(Html)
				} else {
					Html := "<span><a href='javascript:;'>请重新添加</a></span>"
					return c.RenderText(Html)
				}
			}
		}
	}

	Html := "<span><a href='javascript:;'>未登陆</a></span>"
	return c.RenderText(Html)
}

const (
	C_PanelDel_Succ     string = "0"
	C_PanelDel_NotLogin string = "1"
	C_PanelDel_Fail     string = "2"
)

//删除快捷方式
func (c *Ajax) DelPanel(admin_panel *models.AdminPanel) revel.Result {

	data := make(map[string]string)
	if adminId, ok := utils.ParseAdminId(utils.GetSessionValue(consts.C_Session_AdminID, c.Session)); ok {

		mid := c.Params.Get("mid")
		if menuId, ok := utils.ParseMenuId(mid); ok {
			//获取登陆用户信息
			admin := new(models.Admin)
			admin_info := admin.GetById(adminId)

			ok := admin_panel.DelPanel(menuId, admin_info)
			if ok {
				data["status"] = C_PanelDel_Succ
				data["message"] = "取消成功!"
				return c.RenderJson(data)
			} else {
				data["status"] = C_PanelDel_Fail
				data["message"] = "取消失败!"
				return c.RenderJson(data)
			}
		}
	}

	data["status"] = C_PanelDel_NotLogin
	data["message"] = "未登陆!"
	return c.RenderJson(data)
}
