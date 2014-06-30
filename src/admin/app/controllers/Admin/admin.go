// Package: Admin
// File: admin.go
// Created by mint
// Useage: 管理员管理
// DATE: 14-6-29 23:37
package Admin

import (
	"github.com/revel/revel"
	"admin/app/models"
	"admin/utils"
)

type Admin struct {
	*revel.Controller
}

//管理员首页
func (c Admin) Index(admin *models.Admin) revel.Result {

	page := c.Params.Get("page")
	where := make(map[string]string)

	if len(page) > 0 {
		if Page, ok := utils.ParsePage(page); ok {
			admin_list, pages := admin.GetByAll(0, where, Page, 10)
			c.Render(admin_list, pages)
		}
	} else {
		admin_list, pages := admin.GetByAll(0, where, 1, 10)
		c.Render(admin_list, pages)
	}

	return c.RenderTemplate("Setting/Admin/Index.html")
}
