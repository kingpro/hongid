// Package: gasStations
// File: gasStations
// Created by mint
// Useage: 会员组：加油站
// DATE: 14-7-12 20:04
package gasStations

import (
	"api/app/models"
	"encoding/json"
	"github.com/revel/revel"
	"regexp"
	"strconv"
	"utils/page"
	"api/app/routes"
	"member"
)

type GasStations struct {
	*revel.Controller
}

const (
	DigitalRegexp string = "^[0-9]*$"
)

//加油站列表
/*
/api/stations(GET): retrieve a list of gas stations by paging

Default size is 10, customize by /api/stations?page=2&size=20
*/
func (g *GasStations) StationsList() revel.Result {

	//==================================================================================================================
	//参数验证
	pageNum, pageSize := 1, 10

	pageStr := g.Params.Get("page")
	g.Validation.Match(pageStr, regexp.MustCompile(DigitalRegexp))
	if len(pageStr) != 0 && !g.Validation.HasErrors() {
		pageNum, _ = strconv.Atoi(pageStr)
	}

	pageSizeStr := g.Params.Get("size")
	g.Validation.Match(pageSizeStr, regexp.MustCompile(DigitalRegexp))
	if len(pageSizeStr) != 0 && !g.Validation.HasErrors() {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}
	//==================================================================================================================
	//获取加油站列表
	pager := &page.Page{
		Perpage:      int64(pageSize),
		Current_page: int64(pageNum),
	}

	stations, exist, err := member.FindMemberListByGroup(member.CMemberGroup_GasStation_ID, pager, models.ReaderEngine)
	if !exist {
		revel.WARN.Println(err.ErrorMessage())

		//TODO Handle 404 error
		rspErr := &models.RspError{err.GetCode(), err.GetMessage()}
		return g.Redirect(routes.ErrorHandler.Handle404(rspErr))
		return g.NotFound(err.ErrorMessage())
	}
	//==================================================================================================================

	return g.RenderJson(stations)
}

//根据ID获取某一加油站的信息
/*
/api/stations/{id}(GET): retrieve a specific gas
Error 404 if not found
*/
func (g *GasStations) GetStationById() revel.Result {

	//==================================================================================================================
	//参数解析
	stationId := 0
	stationIdStr := g.Params.Get("stationId")
	g.Validation.Match(stationIdStr, regexp.MustCompile(DigitalRegexp))
	if len(stationIdStr) != 0 && !g.Validation.HasErrors() {
		stationId, _ = strconv.Atoi(stationIdStr)
	} else {
		return g.NotFound("StationID[%v]", stationIdStr)
	}
	//==================================================================================================================
	//根据ID获取加油站信息
	station, exist, err := member.GetMemberById(int64(stationId), models.ReaderEngine)
	if !exist {
		revel.WARN.Println(err.ErrorMessage())

		//TODO handle 404 error
		return g.NotFound(err.ErrorMessage())
	}
	//==================================================================================================================

	return g.RenderJson(station)
}

/*
/api/stations(POST): create gas station
Error if mandatory fields not filled.
{
    "fieldErrors": [
        {
            "field": "provinceCode",
            "message": "province should not be empty"
        },
        {
            "field": "countryCode",
            "message": "country should not be empty"
        },
        {
            "field": "contact",
            "message": "contact should not be empty"
        },
        {
            "field": "addressDetails",
            "message": "address details should not be empty"
        },
        {
            "field": "cityCode",
            "message": "city should not be empty"
        },
        {
            "field": "name",
            "message": "name should not be empty"
        }
    ]
}
*/
func (g *GasStations) SaveStation() revel.Result {

	//==================================================================================================================
	//解析request body（json）
	reqData := make([]byte, 512)
	//TODO  read body error handle
	n, _ := g.Request.Body.Read(reqData)
	reqData = reqData[0:n]

	reqMsg := new(models.ReqNewStation)
	if err := json.Unmarshal(reqData, reqMsg); err != nil {

		//TODO Unmarshal req body error handle
		revel.WARN.Println(err)
	}

	//客户端IP
	clientIp := g.Request.Header["X-Forwarded-For"]
	reqMsg.ClientIp = clientIp[0]

	rspMsg := new(models.RspNewStation)
	//==================================================================================================================
	//错误验证
	fieldErrors := make([]models.FieldError, 0)
	if len(reqMsg.StationName) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "stationName", Message: "stationName should not be empty"})
	}
	if len(reqMsg.Contact) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "contact", Message: "contact should not be empty"})
	}
	if len(reqMsg.ProvinceCode) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "provinceCode", Message: "provinceCode should not be empty"})
	}
	if len(reqMsg.CityCode) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "cityCode", Message: "cityCode should not be empty"})
	}
	if len(reqMsg.CountyCode) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "countyCode", Message: "countyCode should not be empty"})
	}
	if len(reqMsg.Address) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "address", Message: "address should not be empty"})
	}
	if len(fieldErrors) != 0 {
		rspMsg.FieldErrors = fieldErrors

		return g.RenderJson(rspMsg)
	}
	//==================================================================================================================
	//新建加油站
	stationInfo, flag, err := models.InsertNewGasStation(reqMsg)
	if !flag {
		revel.WARN.Println(err.ErrorMessage())
		//TODO 新建加油站错误
	}
	rspMsg.StationInfo = stationInfo
	//==================================================================================================================

	return g.RenderJson(rspMsg)
}

/**
/api/stations/{id}(PUT): update gas station
Error 404 if not found
 */
func (g *GasStations) UpdateStation() revel.Result {

	//==================================================================================================================
	//参数解析
	stationId := 0
	stationIdStr := g.Params.Get("stationId")
	g.Validation.Match(stationIdStr, regexp.MustCompile(DigitalRegexp))
	if len(stationIdStr) != 0 && !g.Validation.HasErrors() {
		stationId, _ = strconv.Atoi(stationIdStr)
	} else {
		return g.NotFound("StationID[%v]", stationIdStr)
	}
	//==================================================================================================================
	//解析request body(json)
	reqData := make([]byte, 512)
	//TODO  read body error handle
	n, _ := g.Request.Body.Read(reqData)
	reqData = reqData[0:n]

	reqMsg := new(models.ReqUpdStation)
	if err := json.Unmarshal(reqData, reqMsg); err != nil {

		//TODO Unmarshal req body error handle
		revel.WARN.Println(err)
	}

	//客户端IP
	clientIp := g.Request.Header["X-Forwarded-For"]
	reqMsg.ClientIp = clientIp[0]

	rspMsg := new(models.RspUpdStation)
	//==================================================================================================================
	//错误验证
	fieldErrors := make([]models.FieldError, 0)
	if len(reqMsg.StationName) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "stationName", Message: "stationName should not be empty"})
	}
	if len(reqMsg.Contact) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "contact", Message: "contact should not be empty"})
	}
	if len(reqMsg.ProvinceCode) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "provinceCode", Message: "provinceCode should not be empty"})
	}
	if len(reqMsg.CityCode) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "cityCode", Message: "cityCode should not be empty"})
	}
	if len(reqMsg.CountyCode) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "countyCode", Message: "countyCode should not be empty"})
	}
	if len(reqMsg.Address) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "address", Message: "address should not be empty"})
	}
	if len(fieldErrors) != 0 {
		rspMsg.FieldErrors = fieldErrors

		return g.RenderJson(rspMsg)
	}
	//==================================================================================================================
	//验证加油站是否存在
	station, exist, err := member.GetMemberById(int64(stationId), models.ReaderEngine)
	if !exist {
		//TODO handle 404 error
		return g.NotFound(err.ErrorMessage())
	}
	//==================================================================================================================
	//更新加油站
	stationInfo, flag, err := models.UpdateGasStation(reqMsg, station)
	if !flag {
		revel.WARN.Println(err.ErrorMessage())
		//TODO handle error
	}
	rspMsg.StationInfo = stationInfo
	//==================================================================================================================

	return g.RenderJson(rspMsg)
}

/*
/api/stations/{id}(DELETE): delete gas station
Error 404 if not found
Status API Training Shop Blog About
*/
func (g *GasStations) DeleteStation() revel.Result {

	//==================================================================================================================
	//参数解析
	stationId := 0
	stationIdStr := g.Params.Get("stationId")
	g.Validation.Match(stationIdStr, regexp.MustCompile(DigitalRegexp))
	if len(stationIdStr) != 0 && !g.Validation.HasErrors() {
		stationId, _ = strconv.Atoi(stationIdStr)
	} else {
		return g.NotFound("StationID[%v]", stationIdStr)
	}
	//==================================================================================================================
	//删除加油站
	flag, err := member.DelMemberById(int64(stationId), models.WriterEngine)
	if !flag {
		revel.WARN.Println(err.ErrorMessage())
	}

	return g.Render()
}
