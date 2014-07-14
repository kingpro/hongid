// Package: openapi.app.models.memberCard
// File: format_test.go
// Created by mint
// Useage: format测试类
// DATE: 14-7-3 19:52
package memberCard

import "testing"

func TestFmtToInside(t *testing.T) {
	strCard := "1234-5674-8594-8c832-789"
	if card, _ := FmtToInside(strCard); card != "1234567485948832789" {
		t.Logf("strCard: %v", card)
		t.Error("转换错误")
	}
}

func TestFmtToDisplay1(t *testing.T) {
	card := CardNumber("6320860000000012343")
	if strCard, _ := FmtToDisplay(card, EMemberCardDisPlayType_SPACE); strCard != "6320 8600 0000 0012 343" {
		t.Errorf("Error: %v", strCard)
	}
}

func TestFmtToDisplay2(t *testing.T) {
	card := CardNumber("6320860000000012343")
	if strCard, _ := FmtToDisplay(card, EMemberCardDisplayType_DASH); strCard != "6320-8600-0000-0012-343" {
		t.Errorf("Error: %v", strCard)
	}
}

func TestFmtToDisplay3(t *testing.T) {
	card := CardNumber("6320860000000012343")
	if strCard, _ := FmtToDisplay(card, EMemberCardDisplayType_COMMA); strCard != "6320,8600,0000,0012,343" {
		t.Errorf("Error: %v", strCard)
	}
}
