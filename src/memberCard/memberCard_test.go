// Package: openapi.app.models.memberCard
// File: memberCard_test.go
// Created by mint
// Useage: 测试类
// DATE: 14-7-4 17:13
package memberCard

import "testing"

func TestMemberCardString(t *testing.T) {
	card := &MemberCard{
		MII: 6,
		CPI: 32,
		CDI: 86,
		PII: 1,
		IVC: 1,
	}
	if card.String() != CardNumber("6320860000000000011") {
		t.Errorf("Error: %v", card.String())
	}
}

func TestNewMemberCard(t *testing.T) {
	card := NewMemberCard(6, 32, 86, 1)
	if card.String() != CardNumber("6320860000000000013") {
		t.Errorf("Error: %v", card.String())
	}
}
