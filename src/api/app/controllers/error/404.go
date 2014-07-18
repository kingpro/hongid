// Package: error
// File: 404.go
// Created by mint
// Useage: http 404 error
// DATE: 14-7-17 23:34
package error

import (
	"github.com/revel/revel"
	"api/app/models"
	"net/http"
)

func (e *ErrorHandler) Handle404(rspErr *models.RspError) revel.Result {

	e.Response.Status = http.StatusNotFound

	return e.RenderJson(rspErr)
}
