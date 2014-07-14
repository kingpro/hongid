// Package: gasStations
// File: gasStations
// Created by mint
// Useage: 会员组：加油站
// DATE: 14-7-12 20:04
package gasStations

import (
	"github.com/revel/revel"
	"openapi/app/models/member"
	"strconv"
	"regexp"
	"utils/page"
)

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
	CMemberGroup_GasStation_ID   int64 = 1  //加油站会员组ID，默认(固定)：1
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
		Perpage:       int64(pageSize),
		Current_page:  int64(pageNum),
	}

	stations, exist, err := member.FindMemberListByGroup(CMemberGroup_GasStation_ID, pager)
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

func (g *GasStations) SaveStation() revel.Result {


	return g.Render()
}
