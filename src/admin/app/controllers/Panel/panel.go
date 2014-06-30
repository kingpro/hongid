// Package: Panel
// File: panel.go
// Created by mint
// Useage: 我的面板
// DATE: 14-6-28 21:32
package Panel

import "github.com/revel/revel"

type Panel struct {
	*revel.Controller
}

func (c *Panel) Index() revel.Result {
	return c.RenderTemplate("Panel/Index.html")
}
