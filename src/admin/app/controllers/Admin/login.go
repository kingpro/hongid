// Package: Admin
// File: login.go
// Created by mint
// Useage: 管理员后台登录
// DATE: 14-6-27 22:42
package Admin

import (
	"admin/app/models"
	"admin/utils/Const"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/revel/revel"
	"utils/security"
)

type Admin struct {
	*revel.Controller
}

//登陆
func (c *Admin) Login(admin *models.Admin) revel.Result {

	if c.Request.Method == Const.C_Method_Get {
		CaptchaId := captcha.NewLen(6)
		return c.Render(CaptchaId)
	} else if c.Request.Method == Const.C_Method_Post {

		username := c.Params.Get("username")
		password := c.Params.Get("password")

		captchaId := c.Params.Get("captchaId")
		verify := c.Params.Get("verify")

		data := make(map[string]string)

		if LANG, ok := c.Session[Const.C_Session_Lang]; ok {
			//设置语言
			c.Request.Locale = LANG
		} else {
			//设置默认语言
			c.Request.Locale = "zh"
		}

		if !captcha.VerifyString(captchaId, verify) {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("verification_code")
			return c.RenderJson(data)
		}

		if len(username) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_user_name")
			return c.RenderJson(data)
		}

		if len(password) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_password")
			return c.RenderJson(data)
		}

		if len(verify) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_verification_code")
			return c.RenderJson(data)
		}

		admin_info := admin.GetByName(username)

		if admin_info.Id <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_username_error")
		} else if admin_info.Status == 0 && admin_info.Id != Const.C_Founder_ID {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_forbid_login")
		} else if admin_info.Role.Status == 0 && admin_info.Id != Const.C_Founder_ID {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_forbid_role_login")
			//TODO 密码加密解密
		} else if username == admin_info.Username && security.CompareHashAndPassword(admin_info.Password, password) {
			c.Session[Const.C_Session_AdminID] = fmt.Sprintf("%d", admin_info.Id)
			c.Session[Const.C_Session_Lang] = admin_info.Lang

			c.Flash.Success(c.Message("login_success"))
			c.Flash.Out["url"] = "/"

			//更新登陆时间
			admin.UpdateLoginTime(admin_info.Id)

			//******************************************
			//管理员日志
			logs := new(models.Logs)
			desc := "登陆用户名:" + admin_info.Username + "|^|登陆系统!|^|登陆ID:" + fmt.Sprintf("%d", admin_info.Id)
			logs.Save(admin_info, c.Controller, desc)
			//*****************************************

			data["status"] = "1"
			data["url"] = "/Message/"
			data["message"] = c.Message("login_success")
		} else {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_password_error")
		}

		return c.RenderJson(data)
	}

	return c.Render()
}
