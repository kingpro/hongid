// Package: Public
// File: public.go
// Created by mint
// Useage: 后台公共处理
// DATE: 14-6-28 16:51
package Public

import "github.com/revel/revel"

type Public struct {
	*revel.Controller
}

//消息提示
func (c *Public) Message() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/message.html")
}
