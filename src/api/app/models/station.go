// Package: openapi.app.models.member
// File: openapi.app.models.member.go
// Created by mint
// Useage: 会员
// DATE: 14-7-8 14:51
package models

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
	"utils/errors"
	"utils/times"
	"github.com/go-xorm/xorm"
	."member"
)

var (
	_ *xorm.Engine
)

//新建加油站
func InsertNewGasStation(reqMsg *ReqNewStation) (*SationInfo, bool, errors.GlobalWaysError) {

	//TODO UUID VERSION 1
	uuid := uuid.NewUUID().String()
	timeNow := time.Now()

	//从Member表中抽取一个未使用的加油站记录
	memberNew, exist, gErr := GetUnUseMember(CMemberGroup_GasStation_ID, WriterEngine)
	if !exist {
		return nil, false, gErr
	}

	memberNew.UUID = uuid
	memberNew.TelPhone = reqMsg.TelPhone
	memberNew.Email = reqMsg.Email
	memberNew.NickName = reqMsg.StationName
	memberNew.Contact = reqMsg.Contact
	memberNew.Phone = reqMsg.Phone
	memberNew.Fax = reqMsg.Fax
	memberNew.PostCode = reqMsg.PostCode
	memberNew.LastTime = timeNow.Format(times.Time_Layout_1)
	memberNew.LastIp = reqMsg.ClientIp
	memberNew.HomePage = reqMsg.HomePage
	memberNew.Status = EMemberStatus_Enable

	//加油站详细信息
	memberProfile := &MemberProfile{
		UUID:         uuid,
		ProvinceCode: reqMsg.ProvinceCode,
		CityCode:     reqMsg.CityCode,
		CountyCode:   reqMsg.CountyCode,
		AdressDetail: reqMsg.Address,
		RegTime:      timeNow.Format(times.Time_Layout_1),
		BirthDay:     timeNow.Format(times.Time_Layout_2),
		CalenDarType: ECalenDarType_Gregorian,
		RegIp:        reqMsg.ClientIp,
	}

	if _, err := GenMember(memberNew, memberProfile, WriterEngine); err.IsError() {
		return nil, false, err
	}

	//TODO handle new station response
	stationInfo := &SationInfo{
		StationName:  reqMsg.StationName,
		Phone:        reqMsg.Phone,
		TelPhone:     reqMsg.TelPhone,
		Contact:      reqMsg.Contact,
		Fax:          reqMsg.Fax,
		Email:        reqMsg.Email,
		HomePage:     reqMsg.HomePage,
		PostCode:     reqMsg.PostCode,
		ProvinceCode: reqMsg.ProvinceCode,
		CityCode:     reqMsg.CityCode,
		CountyCode:   reqMsg.CountyCode,
		Address:      reqMsg.Address,
	}

	return stationInfo, true, errors.ErrorOK()
}

//根据ID更新加油站
func UpdateGasStation(reqMsg *ReqUpdStation, condiBean *Member) (*SationInfo, bool, errors.GlobalWaysError) {

	timeNow := time.Now()

	memberUdp := &Member{
		Id:       condiBean.Id,
		UUID:     condiBean.UUID,
		TelPhone: reqMsg.TelPhone,
		Email:    reqMsg.Email,
		NickName: reqMsg.StationName,
		Contact:  reqMsg.Contact,
		Phone:    reqMsg.Phone,
		Fax:      reqMsg.Fax,
		PostCode: reqMsg.PostCode,
		LastTime: timeNow.Format(times.Time_Layout_1),
		LastIp:   reqMsg.ClientIp,
		HomePage: reqMsg.HomePage,
	}

	memberProfile := &MemberProfile{
		ProvinceCode: reqMsg.ProvinceCode,
		CityCode:     reqMsg.CityCode,
		CountyCode:   reqMsg.CountyCode,
		AdressDetail: reqMsg.Address,
		CalenDarType: ECalenDarType_Gregorian,
	}

	if _, err := UpdaeteMember(memberUdp, memberProfile, WriterEngine); err.IsError() {
		return nil, false, err
	}

	//TODO handle upd station response
	stationInfo := &SationInfo{
		StationName:  reqMsg.StationName,
		Phone:        reqMsg.Phone,
		TelPhone:     reqMsg.TelPhone,
		Contact:      reqMsg.Contact,
		Fax:          reqMsg.Fax,
		Email:        reqMsg.Email,
		HomePage:     reqMsg.HomePage,
		PostCode:     reqMsg.PostCode,
		ProvinceCode: reqMsg.ProvinceCode,
		CityCode:     reqMsg.CityCode,
		CountyCode:   reqMsg.CountyCode,
		Address:      reqMsg.Address,
	}

	return stationInfo, true, errors.ErrorOK()
}

