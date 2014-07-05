// Package: algorith
// File: luhn_test.go
// Created by mint
// Useage: luhn验证算法测试
// DATE: 14-7-3 18:53
package algorith

import "testing"

func TestGenerateChkDigit(t *testing.T) {
	card := "7992739871"
	if b := GenLuhnCheckDigit([]byte(card)); b != 3 {
		t.Errorf("Error: %v", b)
	}
}

func TestValidateLuhn1(t *testing.T) {
	card := "79927398713"
	if !ValidateLuhn(card) {
		t.Errorf("Error")
	}
}

func TestValidateLuhn2(t *testing.T) {
	card := "79927394713"
	if ValidateLuhn(card) {
		t.Errorf("Error")
	}
}
