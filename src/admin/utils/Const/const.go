// Package: Const
// File: const.go
// Created by mint
// Useage: 常量类
// DATE: 14-6-28 17:50
package Const

//解锁类型
const (
	C_UnLock_1 string = "1" //锁屏状态解锁成功
	C_UnLock_2 string = "2" //密码为空
	C_UnLock_3 string = "3" //密码错误
	C_UnLock_4 string = "4" //session过期，重新登陆
)

//桌面锁定状态
const (
	C_Lock_0 string = "0" //解锁状态
	C_Lock_1 string = "1" //锁屏状态
)

const (
	C_Founder_ID int64 = 1 //创始人，ID：1
)

//http提交方法
const (
	C_Method_Get  string = "GET"
	C_Method_Post string = "POST"
	C_Method_Put  string = "PUT"
)

//系统存储的session
const (
	C_Session_AdminID string = "AdminID"
	C_Session_Lang    string = "Lang"
	C_Session_LockS   string = "LockScreen"
)
