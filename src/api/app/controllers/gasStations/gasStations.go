// Package: gasStations
// File: gasStations
// Created by mint
// Useage: 会员组：加油站
// DATE: 14-7-12 20:04
package gasStations

import (
	"api/app/models"
	"api/app/models/member"
	"encoding/json"
	"github.com/revel/revel"
	"regexp"
	"strconv"
	"utils/page"
)

/*
/api/stations/{id}(PUT): update gas station
Error 404 if not found
/api/stations/{id}(DELETE): delete gas station
Error 404 if not found
Status API Training Shop Blog About
*/

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

	pager := &page.Page{
		Perpage:      int64(pageSize),
		Current_page: int64(pageNum),
	}

	stations, exist, err := member.FindMemberListByGroup(member.CMemberGroup_GasStation_ID, pager)
	if !exist {
		g.NotFound(err.ErrorMessage())
	}

	return g.RenderJson(stations)
}

//根据ID获取某一加油站的信息
/*
/api/stations/{id}(GET): retrieve a specific gas
Error 404 if not found
*/
func (g *GasStations) GetStationById() revel.Result {

	stationId := 0

	stationIdStr := g.Params.Get("stationId")
	g.Validation.Match(stationIdStr, regexp.MustCompile(DigitalRegexp))
	if len(stationIdStr) != 0 && !g.Validation.HasErrors() {
		stationId, _ = strconv.Atoi(stationIdStr)
	} else {
		return g.NotFound("StationID[%v]", stationIdStr)
	}

	station, exist, err := member.GetMemberById(int64(stationId))
	if !exist {
		return g.NotFound(err.ErrorMessage())
	}

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

加油站
stationName 名称
phone 电话
telPhone 手机
contact 联系人
fax 传真
email 邮箱
homePage 主页
postCode 邮编
provinceCode 省
cityCode 市
countyCode 区/县
address 地址详细
*/

func (g *GasStations) SaveStation() revel.Result {

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
	if len(reqMsg.Email) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "email", Message: "email should not be empty"})
	}
	if len(reqMsg.HomePage) == 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "homePage", Message: "homePage should not be empty"})
	}
	if len(fieldErrors) != 0 {
		errorFields := &models.FieldErrors{
			FieldErrors: fieldErrors,
		}

		return g.RenderJson(errorFields)
	}
	//==================================================================================================================

	rspMsg, err := member.InsertNewGasStation(reqMsg)
	if err.IsError() {
		//TODO 新建加油站错误
	}

	return g.RenderJson(rspMsg)
}
