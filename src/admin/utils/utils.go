// Package: utils
// File: utils.go
// Created by mint
// Useage: 相关工具类集合
// DATE: 14-6-27 23:33
package utils

import (
	"github.com/revel/revel"
	"strconv"
	"strings"
	"utils/container"
	"utils/security"
)

//将搜索条件数组编码成一个字符串
//where 搜索条件Map
//例如:map转换成格式 name:张三|age:20 然后Base64
func EncodeSegment(where map[string]string) string {
	if !container.IsMap(where) {
		return ""
	}

	whereStr := ""
	for key, val := range where {
		whereStr += key + ":" + val + "|"
	}

	whereStr = security.Base64Encode([]byte(strings.Trim(whereStr, "|")))

	return whereStr
}

//将编码的搜索条件，解码成搜索条件Map数组
//where 搜索条件的编码字符串
func DecodeSegment(where string) map[string]string {
	where_map := make(map[string]string)

	where = security.Base64Decode(where)

	Search := strings.Split(where, "|")

	if len(Search) > 0 {
		for _, val := range Search {
			arr := strings.Split(val, ":")
			where_map[arr[0]] = arr[1]
		}
	}

	return where_map
}

//获取session的值
func GetSessionValue(key string, session revel.Session) string {
	if value, ok := session[key]; ok {
		return value
	}

	return ""
}

//解析管理员ID
func ParseAdminId(idStr string) (int64, bool) {
	return parseInt(idStr, "管理员ID")
}

//解析menuID
func ParseMenuId(idStr string) (int64, bool) {
	return parseInt(idStr, "菜单ID")
}

//解析menu的order
func ParseMenuOrder(orderStr string) (int64, bool) {
	return parseInt(orderStr, "菜单排序")
}

//解析menu的display
func ParseMenuDisplay(displayStr string) (int64, bool) {
	return parseInt(displayStr, "菜单显示")
}

//解析分页page
func ParsePage(pageStr string) (int64, bool) {
	page := int64(1)
	ok := false
	if page, ok = parseInt(pageStr, "分页"); ok {
		page = int64(1)
	}

	return page, ok
}

//解析roleID
func ParseRoleId(idStr string) (int64, bool) {
	return parseInt(idStr, "角色ID")
}

//解析启用Status
func ParseStatus(status string) (int64, bool) {
	return parseInt(status, "启用状态")
}

//解析AnnounceID
func ParseAnnounceId(aidStr string) (int64, bool) {
	return parseInt(aidStr, "公告ID")
}

//解析strin类型的int值
func parseInt(intStr, desc string) (int64, bool) {
	intVal, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		revel.WARN.Printf("%v[%v]解析错误: %v", desc, intStr, err)
		return 0, false
	}

	return intVal, true
}
