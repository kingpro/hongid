// Package: openapi.app.models.memberCard
// File: group.go
// Created by mint
// Useage: 会员卡分组：实名卡、匿名卡、车队卡、品牌宣传卡、礼品卡...
// DATE: 14-7-8 14:37
package memberCard

type CardGroup struct {
	Id         int64
	GroupName  string  `xorm:"varchar(20) unique"`
	GroupEName string  `xorm:"varchar(20) unique"`
	GroupDesc  string
	Points     uint32
}

func (g *CardGroup) GetCardGroupName() string {
	return g.GroupName
}

func (g *CardGroup) SetCardGroupName(name string) {
	g.GroupName = name
}

func (g *CardGroup) GetCardGroupEName() string {
	return g.GroupEName
}

func (g *CardGroup) SetCardGroupEName(ename string) {
	g.GroupEName = ename
}

func (g *CardGroup) GetCardGroupDesc() string {
	return g.GroupDesc
}

func (g *CardGroup) SetCardGroupDesc(desc string) {
	g.GroupDesc = desc
}

func (g *CardGroup) GetCardGroupPoints() uint32 {
	return g.Points
}

func (g *CardGroup) SetCardGroupPoints(points uint32) {
	g.Points = points
}

func NewCardGroup(name, ename, desc string, points uint32) *CardGroup {
	return &CardGroup{
		GroupName:   name,
		GroupEName:  ename,
		GroupDesc:   desc,
		Points:      points,
	}
}
