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
	member := &Member{
		GroupId: CMemberGroup_GasStation_ID,
	}
	if exist, err := ReaderEngine.Where("Status = ?", EMemberStatus_NotUse).Get(member); err != nil {
		return nil, false, errors.Newf(errors.CODE_DB_ERR_GET, "Get unuse member record return error: %v", err)
	} else if !exist {
		return nil, false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	//事务处理
	session := WriterEngine.NewSession()
	defer session.Close()

	err := session.Begin()
	//TODO memberCard handle
	//更新信息
	memberUpd := &Member{
		UUID:     uuid,
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
		Status:   EMemberStatus_Enable,
	}

	affacted := true
	if _, err = WriterEngine.Update(memberUpd, member); err != nil {
		affacted = false
		session.Rollback()
		return nil, false, errors.Newf(errors.CODE_DB_ERR_UPDATE, "Turn unuse member record[%+v] return error :%v", member, err)
	}

	//如果成功更新会员主表信息，则添加一条详细信息
	if affacted {
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
		if _, err = WriterEngine.Insert(memberProfile); err != nil {
			session.Rollback()
			return nil, false, errors.Newf(errors.CODE_DB_ERR_INSERT, "InsertMemberProfile return error: %v", err)
		}
	}
	err = session.Commit()
	if err != nil {
		return nil, false, errors.New(errors.CODE_DB_ERR_COMMIT, errors.MSG_DB_ERR_COMMIT)
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

	session := WriterEngine.NewSession()
	defer session.Close()

	//事务处理
	err := session.Begin()

	memberUdp := &Member{
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
	affacted := true
	if _, err = WriterEngine.Update(memberUdp, condiBean); err != nil {
		affacted = false
		session.Rollback()
		return nil, false, errors.Newf(errors.CODE_DB_ERR_UPDATE, "Update Gas Station return error: %v", err)
	}

	//如果成功更新主表信息，则更新profile表信息
	if affacted {
		memberProfile := &MemberProfile{
			ProvinceCode: reqMsg.ProvinceCode,
			CityCode:     reqMsg.CityCode,
			CountyCode:   reqMsg.CountyCode,
			AdressDetail: reqMsg.Address,
			BirthDay:     timeNow.Format(times.Time_Layout_2),
			CalenDarType: ECalenDarType_Gregorian,
		}
		if _, err = WriterEngine.Where("UUID = ?", condiBean.UUID).Update(memberProfile); err != nil {
			session.Rollback()
			return nil, false, errors.Newf(errors.CODE_DB_ERR_UPDATE, "Update Gas Station Profile return error: %v", err)
		}
	}
	err = session.Commit()
	if err != nil {
		return nil, false, errors.New(errors.CODE_DB_ERR_COMMIT, errors.MSG_DB_ERR_COMMIT)
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

