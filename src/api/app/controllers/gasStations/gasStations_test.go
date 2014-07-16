// Package: gasStations
// File:
// Created by mint
// Useage:
// DATE: 14-7-15 22:03
package gasStations

import (
	"api/app/models"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	WEB_SERVER_ADDR = "http://localhost:9000"
)

func SendHttpPost(dataBuf io.Reader, protoAddr string) (outbuffer []byte, err error) {
	resp, err := http.Post(WEB_SERVER_ADDR+protoAddr, "application/binary", dataBuf)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	outbuffer, err = ioutil.ReadAll(resp.Body)

	return outbuffer, err
}

func TestSaveStationErrorField(t *testing.T) {
	reqMsg := models.ReqNewStation{
		StationName: "中和加油站",
		Phone:        "028-8348573",
		TelPhone:     "13498572301",
		Contact:      "中和",
		Fax:          "028-4758293",
		Email:        "zi__chen@183.com",
		HomePage:     "/zhonghe",
		PostCode:     "610500",
		ProvinceCode: "519845",
		CityCode:     "885933",
		CountyCode:   "587342",
		Address:      "成都市中和菜市场门口",
	}

	reqData, _ := json.Marshal(reqMsg)

	out, _ := SendHttpPost(bytes.NewReader(reqData), "/api/stations")

	t.Errorf("error: %+v", string(out))
}
