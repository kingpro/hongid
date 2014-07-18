// Package: models
// File: response.go
// Created by mint
// Useage: http响应
// DATE: 14-7-16 11:11
package models

type SationInfo struct {
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
}

type RspNewStation struct {
	StationInfo *SationInfo  `json:"stationInfo"`
	FieldErrors []FieldError `json:"fieldErrors"`
}

type RspUpdStation struct {
	StationInfo *SationInfo  `json:"stationInfo"`
	FieldErrors []FieldError `json:"fieldErrors"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RspError struct {
	ErrorCode   int8    `json:"errorCode"`
	Message     string  `json:"message"`
}
