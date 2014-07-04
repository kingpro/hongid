// Package: memberCard
// File: validation_test.go
// Created by mint
// Useage: 会员卡验证测试
// DATE: 14-7-3 19:34
package memberCard

import "testing"

func TestValidateCard1(t *testing.T) {
	var card1 CardNumber = "6225880169758712"
	if ValidateCard(&card1) {
		t.Logf("Ok")
	}else {
		t.Errorf("error")
	}
}

func TestValidateCard2(t *testing.T) {
	var card2 CardNumber = "6225880169752712"
	if ValidateCard(&card2) {
		t.Errorf("error")
	}else {
		t.Logf("ok")
	}
}
