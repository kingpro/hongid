package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	data := make(map[string]string)
	data["1"] = "1"
	data["2"] = "2"
	return c.RenderJson(data)
}
