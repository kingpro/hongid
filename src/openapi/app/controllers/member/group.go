// Package: member
// File: group.go
// Created by mint
// Useage: 会员组controller
// DATE: 14-7-11 17:32
package member

import (
	"github.com/revel/revel"
	"strconv"
	"openapi/app/models/member"
)

type MemberGroup struct {
	*revel.Controller
}

type RspMemberGroup struct {
	GroupName    string
	GroupDesc    string
	GrowthPoints uint32
	Contribution uint32
	Status       member.EMemberGroupStatus
}

func (c *MemberGroup) GetMemberGroupById() revel.Result {
	groupIdStr := c.Params.Get("groupId")
	groupId, _ := strconv.Atoi(groupIdStr)

	group, _, _ := member.GetMemberGroupById(int64(groupId))

//	rspMsg := &RspMemberGroup{
//		GroupName: group.GetMemberGroupName(),
//		GroupDesc: group.GetMemberGroupDesc(),
//		GrowthPoints: group.GetMemberGroupGPoints(),
//		Contribution: group.GetMemberGroupContribution(),
//		Status: group.GetMemberGroupStatus(),
//	}

	return c.RenderJson(group)
}

func (c *MemberGroup) GetMemberGroupByName() revel.Result {
	groupNameStr := c.Params.Get("groupName")

	group, _, _ := member.GetMemberGroupByName(groupNameStr)

	rspMsg := &RspMemberGroup{
		GroupName: group.GetMemberGroupName(),
		GroupDesc: group.GetMemberGroupDesc(),
		Contribution: group.GetMemberGroupContribution(),
		Status: group.GetMemberGroupStatus(),
	}

	return c.RenderJson(rspMsg)
}
