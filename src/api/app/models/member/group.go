// Package: openapi.app.models.member
// File: group.go
// Created by mint
// Useage: 会员组: 个人会员、企业会员、机构会员...
// DATE: 14-7-8 14:55
package member

import (
	"api/app/models"
	"github.com/revel/revel"
	"time"
	"utils/errors"
	"utils/times"
)

const (
	CMemberGroup_GasStation_ID int64 = 1 //加油站会员组ID，默认(固定)：1
)

type EMemberGroupStatus byte

const (
	EMemberGroupStatus_Enable EMemberGroupStatus = iota
	EMemberGroupStatus_Distable
	EMemberGroupStatus_Del
)

type MemberGroup struct {
	Id           int64
	GroupName    string `xorm:"varchar(20) unique"`
	GroupDesc    string
	Contribution uint32             //会费
	Status       EMemberGroupStatus `xorm:"tinyint(1)"`
	CreateTime   string             `xorm:"DateTime created"`
	UpdateTime   string             `xorm:"DateTime updated"`
}

func NewMemberGroup(name, desc string, contribution uint32) *MemberGroup {
	return &MemberGroup{
		GroupName:    name,
		GroupDesc:    desc,
		Contribution: contribution,
		Status:       EMemberGroupStatus_Enable,
		CreateTime:   time.Now().Format(times.Time_Layout_1),
		UpdateTime:   time.Now().Format(times.Time_Layout_1),
	}
}

func (g *MemberGroup) GetMemberGroupName() string {
	return g.GroupName
}

func (g *MemberGroup) SetMemberGroupName(name string) {
	g.GroupName = name
}

func (g *MemberGroup) GetMemberGroupDesc() string {
	return g.GroupDesc
}

func (g *MemberGroup) SetMemberGroupDesc(desc string) {
	g.GroupDesc = desc
}

func (g *MemberGroup) GetMemberGroupContribution() uint32 {
	return g.Contribution
}

func (g *MemberGroup) SetMemberGroupContribution(contribution uint32) {
	g.Contribution = contribution
}

func (g *MemberGroup) GetMemberGroupStatus() EMemberGroupStatus {
	return g.Status
}

func (g *MemberGroup) SetMemberGroupStatus(status EMemberGroupStatus) {
	g.Status = status
}

func (g *MemberGroup) GetMemberGroupCreateTime() string {
	return g.CreateTime
}

func (g *MemberGroup) GetMemberGroupUpdateTime() string {
	return g.UpdateTime
}

func (g *MemberGroup) SetMemberGroupUpdateTime(time string) {
	g.UpdateTime = time
}

//数据库操作
//Insert
func AddMemberGroup(memberGroup *MemberGroup) bool {

	return true
}

//Select
func GetMemberGroupById(id int64) (memberGroup *MemberGroup, exist bool, err errors.GlobalWaysError) {
	memberGroup = new(MemberGroup)
	if exist, err = models.ReaderEngine.GetById(id, memberGroup); err.IsError() {
		revel.WARN.Printf("GetMemberGroupById[%v] ret err: %v", id, err.ErrorMessage())
	}
	return
}

func GetMemberGroupByName(name string) (memberGroup *MemberGroup, exist bool, err errors.GlobalWaysError) {
	memberGroup = new(MemberGroup)
	if exist, err = models.ReaderEngine.GetByCol("GroupName", "=", name, memberGroup); err.IsError() {
		revel.WARN.Printf("GetMemberGroupByName[%v] ret err: %v", name, err.ErrorMessage())
	}
	return
}
