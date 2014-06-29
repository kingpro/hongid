// Package: Public
// File: captcha.go
// Created by mint
// Useage: 验证码服务
// DATE: 14-6-28 14:49
package Public

import (
	"github.com/dchest/captcha"
	"github.com/revel/revel"
)

type Captcha struct {
	*revel.Controller
}

//首页
func (c *Captcha) Index() revel.Result {
	captcha.Server(250, 62)
	var CaptchaId string = c.Params.Get("CaptchaId")
	captcha.WriteImage(c.Response.Out, CaptchaId, 250, 62)
	return nil
}

//返回验证码
func (c *Captcha) GetCaptchaId() revel.Result {
	CaptchaId := captcha.NewLen(6)
	return c.RenderText(CaptchaId)
}
