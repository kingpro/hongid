// Package: errors
// File: errorCode.go
// Created by mint
// Useage: 错误编号 & 错误信息
// DATE: 14-7-9 10:18
package errors

//错误编号
const (
	CODE_SUCCESS = 0

	//DB ERROR
	CODE_DB_ERR_BASE      = -100
	CODE_DB_ERR_BADCONN   = CODE_DB_ERR_BASE - 1
	CODE_DB_ERR_NODATA    = CODE_DB_ERR_BASE - 2
	CODE_DB_ERR_GETByID   = CODE_DB_ERR_BASE - 3
	CODE_DB_ERR_GETByCOL  = CODE_DB_ERR_BASE - 4
	CODE_DB_ERR_FINDByCol = CODE_DB_ERR_BASE - 5
	CODE_DB_ERR_INSERT    = CODE_DB_ERR_BASE - 6
	CODE_DB_ERR_UPDATE    = CODE_DB_ERR_BASE - 7
)

//错误信息
const (
	MSG_SUCCESS = "获取数据成功"

	//DB MSG
	MSG_DB_ERR_NODATA = "No data exist!"
)
