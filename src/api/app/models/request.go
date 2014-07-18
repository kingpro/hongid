// Package: models
// File: request.go
// Created by mint
// Useage: http请求
// DATE: 14-7-16 10:46
package models

//请求新建加油站
type ReqNewStation struct {
	StationName  string `json:"stationName"`
	Phone        string `json:"phone"`
	TelPhone     string `json:"telPhone"`
	Contact      string `json:"contact"`
	Fax          string `json:"fax"`
	Email        string `json:"email"`
	HomePage     string `json:"homePage"`
	PostCode     string `json:"postCode"`
	ProvinceCode string `json:"provinceCode"`
	CityCode     string `json:"cityCode"`
	CountyCode   string `json:"countyCode"`
	Address      string `json:"address"`
	ClientIp     string `json:"-"`
}

//请求修改加油站
type ReqUpdStation struct {
	StationName  string `json:"stationName"`
	Phone        string `json:"phone"`
	TelPhone     string `json:"telPhone"`
	Contact      string `json:"contact"`
	Fax          string `json:"fax"`
	Email        string `json:"email"`
	HomePage     string `json:"homePage"`
	PostCode     string `json:"postCode"`
	ProvinceCode string `json:"provinceCode"`
	CityCode     string `json:"cityCode"`
	CountyCode   string `json:"countyCode"`
	Address      string `json:"address"`
	ClientIp     string `json:"-"`
}
