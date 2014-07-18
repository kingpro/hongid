// Package: Setting
// File: menu.go
// Created by mint
// Useage: 菜单类
// DATE: 14-6-29 20:35
package Setting

import (
	"admin/app/controllers"
	"admin/app/models"
	"admin/utils"
	"admin/utils/consts"
	"github.com/revel/revel"
)

type Menu struct {
	*revel.Controller
}

//菜单管理首页
func (c *Menu) Index(menu *models.Menu) revel.Result {

	//获取管理员信息
	if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
		menus := menu.GetMenuHtml(admin_info)
		c.Render(menus)
	}

	return c.RenderTemplate("Setting/Menu/Index.html")
}

//添加菜单
func (c *Menu) Add(menu *models.Menu) revel.Result {

	if c.Request.Method == consts.C_Method_Get {

		id := c.Params.Get("id")
		//添加子菜单
		if len(id) > 0 {
			if menuId, ok := utils.ParseMenuId(id); ok {

				if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

					//返回菜单Option的HTML
					menus := menu.GetMenuOptionHtml(menuId, admin_info)
					c.Render(menus, menuId)
				} else {
					c.Render(menuId)
				}
			}
		} else { //非子菜单
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

				//返回菜单Option的HTML
				menus := menu.GetMenuOptionHtml(0, admin_info)
				c.Render(menus)
			} else {
				c.Render()
			}
		}

		return c.RenderTemplate("Setting/Menu/Add.html")
	} else if c.Request.Method == consts.C_Method_Post {

		pid := c.Params.Get("pid")
		if len(pid) > 0 {
			if Pid, ok := utils.ParseMenuId(pid); ok {
				menu.Pid = Pid
			}
		} else {
			c.Flash.Error("请选择父菜单!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		name := c.Params.Get("name")
		if len(name) > 0 {
			menu.Name = name
		} else {
			c.Flash.Error("请输入中文语言名称!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		enname := c.Params.Get("enname")
		if len(enname) > 0 {
			menu.EName = enname
		} else {
			c.Flash.Error("请输入英文语言名称!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		url := c.Params.Get("url")
		if len(url) > 0 {
			menu.Url = url
		} else {
			c.Flash.Error("请输入菜单地址!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		order := c.Params.Get("order")
		if len(order) > 0 {
			if Order, ok := utils.ParseMenuOrder(order); ok {
				menu.Order = Order
			} else {
				c.Flash.Error("输入的排序非数值，请重新输入!")
				c.Flash.Out["url"] = "/Menu/Add/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("请输入排序!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		data := c.Params.Get("data")
		menu.Data = data

		display := c.Params.Get("display")
		if len(display) > 0 {
			if Display, ok := utils.ParseMenuDisplay(display); ok {
				menu.Display = Display
			}
		} else {
			c.Flash.Error("请选择是否显示菜单!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		if menu.Save() {
			//******************************************
			//管理员日志
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

				logs := new(models.Logs)
				desc := "添加菜单:" + name + "|^|菜单管理"
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			c.Flash.Success("添加菜单成功")
			c.Flash.Out["url"] = "/Menu/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加菜单失败")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}
	} else {

	}

	return c.Render()
}

//删除状态
const (
	C_MenuDel_Fail string = "0"
	C_MenuDel_Succ string = "1"
)

//删除
func (c Menu) Delete(menu *models.Menu) revel.Result {
	id := c.Params.Get("id")

	data := make(map[string]string)

	if len(id) <= 0 {
		data["status"] = C_MenuDel_Fail
		data["message"] = "参数错误!"
	}

	if menuId, ok := utils.ParseMenuId(id); ok {
		if menu.DelByID(menuId) {
			//******************************************
			//管理员日志
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
				logs := new(models.Logs)
				desc := "删除菜单|^|ID:" + id
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			data["status"] = C_MenuDel_Succ
			data["message"] = "删除成功!"
		} else {
			data["status"] = C_MenuDel_Fail
			data["message"] = "删除失败!"
		}
	}

	return c.RenderJson(data)
}

//编辑菜单
func (c Menu) Edit(menu *models.Menu) revel.Result {

	if c.Request.Method == consts.C_Method_Get {

		id := c.Params.Get("id")
		if len(id) > 0 {
			if menuId, ok := utils.ParseMenuId(id); ok {
				//获取菜单信息
				menu_info := menu.GetById(menuId)

				if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
					//返回菜单Option的HTML
					menus := menu.GetMenuOptionHtml(menu_info.Pid, admin_info)

					c.Render(menu_info, menus)
				} else {
					c.Render(menu_info)
				}
			}
		} else {
			if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {

				//返回菜单Option的HTML
				menus := menu.GetMenuOptionHtml(0, admin_info)
				c.Render(menus)
			} else {
				c.Render()
			}
		}
		return c.RenderTemplate("Setting/Menu/Edit.html")
	} else if c.Request.Method == consts.C_Method_Post {

		id := c.Params.Get("id")
		if len(id) > 0 {
			if menuId, ok := utils.ParseMenuId(id); ok {

				pid := c.Params.Get("pid")
				if len(pid) > 0 {
					if Pid, ok := utils.ParseMenuId(pid); ok {
						menu.Pid = Pid
					}
				} else {
					c.Flash.Error("请选择父菜单!")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				name := c.Params.Get("name")
				if len(name) > 0 {
					menu.Name = name
				} else {
					c.Flash.Error("请输入中文语言名称!")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				enname := c.Params.Get("enname")
				if len(enname) > 0 {
					menu.EName = enname
				} else {
					c.Flash.Error("请输入英文语言名称!")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				url := c.Params.Get("url")
				if len(url) > 0 {
					menu.Url = url
				} else {
					c.Flash.Error("请输入菜单地址!")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				order := c.Params.Get("order")
				if len(order) > 0 {
					if Order, ok := utils.ParseMenuOrder(order); ok {
						menu.Order = Order
					} else {
						c.Flash.Error("输入的排序非数值，请重新输入!")
						c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
						return c.Redirect("/Message/")
					}
				} else {
					c.Flash.Error("请输入排序!")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				data := c.Params.Get("data")
				menu.Data = data

				display := c.Params.Get("display")
				if len(display) > 0 {
					if Display, ok := utils.ParseMenuDisplay(display); ok {
						menu.Display = Display
					}
				} else {
					c.Flash.Error("请选择是否显示菜单!")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				if menu.Edit(menuId) {
					//******************************************
					//管理员日志
					if admin_info, ok := controllers.GetAdminInfoBySession(c.Session); ok {
						logs := new(models.Logs)
						desc := "编辑菜单:" + name + "|^|菜单管理|^|ID:" + id
						logs.Save(admin_info, c.Controller, desc)
					}
					//*****************************************

					c.Flash.Success("编辑菜单成功")
					c.Flash.Out["url"] = "/Menu/"
					return c.Redirect("/Message/")
				} else {
					c.Flash.Error("编辑菜单失败")
					c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}
			}
		}
	} else {

	}

	return c.Render()
}
