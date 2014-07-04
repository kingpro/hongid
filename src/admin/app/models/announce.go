// Package: models
// File: announce.go
// Created by mint
// Useage: 公告
// DATE: 14-7-2 19:50
package models

import (
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"html/template"
	"utils/page"
	"time"
	"utils/times"
)

var _ *xorm.Engine

type Announce struct {
	Id         int64  `xorm:"pk"`
	Title      string `xorm:"varchar(255)"`
	Content    string `xorm:"text"`
	Starttime  string `xorm:"DateTime"`
	Endtime    string `xorm:"DateTime"`
	Hits       int64  `xorm:"int(11)"`
	Status     int64  `xorm:"default 1"`
	Createtime string `xorm:"DateTime"`
}

//根据ID获取公告信息
func (a *Announce) GetById(Id int64) *Announce {
	announce := new(Announce)

	_, err := DB_Read.Id(Id).Get(announce)
	if err != nil {
		revel.WARN.Printf("根据ID[%v]获取公告信息错误: %v", err)
	}

	return announce
}

//获取公告列表
func (a *Announce) GetByAll(Page int64, Perpage int64) ([]*Announce, template.HTML) {
	announce_list := []*Announce{}

	//查询总数
	announce := new(Announce)
	total, err := DB_Read.Count(announce)
	if err != nil {
		revel.WARN.Printf("查询公告总数错误: %v", err)
	}

	//分页
	Pager := new(page.Page)
	Pager.SubPage_link = "/Announce/"
	Pager.Nums = total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Limit(int(Perpage), int((Page - 1)*Pager.Perpage)).Desc("id").Find(&announce_list)

	return announce_list, pages
}

//添加公告
func (a *Announce) Save() bool {

	announce := new(Announce)
	announce.Title = a.Title
	announce.Content = a.Content
	announce.Starttime = a.Starttime
	announce.Endtime = a.Endtime
	announce.Hits = 0
	announce.Status = a.Status
	announce.Createtime = time.Now().Format(times.Time_Layout_1)

	_, err := DB_Write.Insert(announce)
	if err != nil {
		revel.WARN.Printf("添加公告错误: %v", err)
		return false
	}

	return true
}

//编辑公告
func (a *Announce) Edit(Id int64) bool {
	announce := new(Announce)

	if len(a.Title) > 0 {
		announce.Title = a.Title
	}

	if len(a.Content) > 0 {
		announce.Content = a.Content
	}

	if len(a.Starttime) > 0 {
		announce.Starttime = a.Starttime
	}

	if len(a.Endtime) > 0 {
		announce.Endtime = a.Endtime
	}

	announce.Status = a.Status

	_, err := DB_Write.Id(Id).Cols("title", "content", "starttime", "endtime", "status").Update(announce)
	if err != nil {
		revel.WARN.Printf("编辑公告错误: %v", err)
		return false
	}

	return true
}

//删除公告
func (a *Announce) DelByID(Id int64) bool {
	announce := new(Announce)

	_, err := DB_Write.Id(Id).Delete(announce)
	if err != nil {
		revel.WARN.Printf("删除公告错误: %v", err)
		return false
	}

	return true
}
