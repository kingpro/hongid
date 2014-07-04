// Package: Module
// File: announce.go
// Created by mint
// Useage: 公告
// DATE: 14-7-2 20:03
package Module

import (
	"admin/app/models"
	"github.com/revel/revel"
	"admin/utils"
	"admin/utils/Const"
	"admin/app/controllers"
)

type Announce struct {
	*revel.Controller
}

//公告首页
func (c *Announce) Index(announce *models.Announce) revel.Result {

	page := c.Params.Get("page")

	if len(page) > 0 {
		Page, _ := utils.ParsePage(page)
		announce_list, pages := announce.GetByAll(Page, 10)

		c.Render(announce_list, pages)
	} else {
		announce_list, pages := announce.GetByAll(1, 10)

		c.Render(announce_list, pages)
	}

	return c.RenderTemplate("Module/Announce/Index.html")
}

//添加公告
func (c *Announce) Add(announce *models.Announce) revel.Result {

	if c.Request.Method == Const.C_Method_Get {

		return c.RenderTemplate("Module/Announce/Add.html")
	} else if c.Request.Method == Const.C_Method_Post {

		title := c.Params.Get("title")
		if len(title) > 0 {
			announce.Title = title
		} else {
			c.Flash.Error("请输入公告标题!")
			c.Flash.Out["url"] = "/Announce/add/"
			return c.Redirect("/Message/")
		}

		starttime := c.Params.Get("starttime")
		if len(starttime) > 0 {
			announce.Starttime = starttime
		} else {
			c.Flash.Error("请输入起始日期!")
			c.Flash.Out["url"] = "/Announce/add/"
			return c.Redirect("/Message/")
		}

		endtime := c.Params.Get("endtime")
		if len(endtime) > 0 {
			announce.Endtime = endtime
		} else {
			c.Flash.Error("请输入截止日期!")
			c.Flash.Out["url"] = "/Announce/add/"
			return c.Redirect("/Message/")
		}

		content := c.Params.Get("content")
		if len(content) > 0 {
			announce.Content = content
		} else {
			c.Flash.Error("请输入公告内容!")
			c.Flash.Out["url"] = "/Announce/add/"
			return c.Redirect("/Message/")
		}

		status := c.Params.Get("status")
		if len(status) > 0 {
			Status, _ := utils.ParseStatus(status)
			announce.Status = Status
		} else {
			c.Flash.Error("请选择是否启用!")
			c.Flash.Out["url"] = "/Announce/add/"
			return c.Redirect("/Message/")
		}

		if announce.Save() {
			//******************************************
			//管理员日志
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
				logs := new(models.Logs)
				desc := "添加公告:" + title
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			c.Flash.Success("添加公告成功!")
			c.Flash.Out["url"] = "/Announce/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加公告失败!")
			c.Flash.Out["url"] = "/Announce/add/"
			return c.Redirect("/Message/")
		}
	}else {

	}

	return c.Render()
}

//编辑栏目
func (c *Announce) Edit(announce *models.Announce) revel.Result {

	id := c.Params.Get("id")
	announceId := int64(0)
	if len(id) > 0 {
		announceId, _ = utils.ParseAnnounceId(id)
	}

	if c.Request.Method == Const.C_Method_Get {

		if len(id) > 0 {
			announce_info := announce.GetById(announceId)

			c.Render(announce_info)
		} else {
			c.Render()
		}

		return c.RenderTemplate("Module/Announce/Edit.html")
	} else if c.Request.Method == Const.C_Method_Post {

		if len(id) > 0 {

			title := c.Params.Get("title")
			if len(title) > 0 {
				announce.Title = title
			} else {
				c.Flash.Error("请输入公告标题!")
				c.Flash.Out["url"] = "/Announce/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			starttime := c.Params.Get("starttime")
			if len(starttime) > 0 {
				announce.Starttime = starttime
			} else {
				c.Flash.Error("请输入起始日期!")
				c.Flash.Out["url"] = "/Announce/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			endtime := c.Params.Get("endtime")
			if len(endtime) > 0 {
				announce.Endtime = endtime
			} else {
				c.Flash.Error("请输入截止日期!")
				c.Flash.Out["url"] = "/Announce/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			content := c.Params.Get("content")
			if len(content) > 0 {
				announce.Content = content
			} else {
				c.Flash.Error("请输入公告内容!")
				c.Flash.Out["url"] = "/Announce/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			status := c.Params.Get("status")
			if len(status) > 0 {
				Status, _ := utils.ParseStatus(status)
				announce.Status = Status
			} else {
				c.Flash.Error("请选择是否启用!")
				c.Flash.Out["url"] = "/Announce/edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if announce.Edit(announceId) {
				//******************************************
				//管理员日志
				if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
					logs := new(models.Logs)
					desc := "编辑公告|^|ID:" + id
					logs.Save(admin_info, c.Controller, desc)
				}
				//*****************************************

				c.Flash.Success("编辑公告成功!")
				c.Flash.Out["url"] = "/Announce/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑公告失败!")
				c.Flash.Out["url"] = "/Announce/edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑公告失败!")
			c.Flash.Out["url"] = "/Announce/edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	} else {

	}

	return c.Render()
}

const (
	C_DelAnnounce_Fail string = "0"
	C_DelAnnounce_Succ string = "1"
)

//删除公告
func (c *Announce) Delete(announce *models.Announce) revel.Result {

	id := c.Params.Get("id")
	data := make(map[string]string)

	if len(id) > 0 {
		announceId, _ := utils.ParseAnnounceId(id)
		if announce.DelByID(announceId) {
			//******************************************
			//管理员日志
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

				logs := new(models.Logs)
				desc := "删除公告|^|ID:" + id
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			data["status"] = C_DelAnnounce_Succ
			data["message"] = "删除成功!"
		} else {
			data["status"] = C_DelAnnounce_Fail
			data["message"] = "删除失败!"
		}
	} else {
		data["status"] = C_DelAnnounce_Fail
		data["message"] = "删除失败!"
	}

	return c.RenderJson(data)
}
