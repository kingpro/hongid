// Package: memberCard
// File: validation.go
// Created by mint
// Useage: 会员卡验证
// DATE: 14-7-3 19:32
package memberCard

var DELTA = []int{0, 1, 2, 3, 4, -4, -3, -2, -1, 0}

func ValidateCard(cc *CardNumber) bool {
	checksum := 0
	bOdd := false
	card := []byte(*cc)
	for i := len(card) - 1; i > -1; i-- {
		cn := int(card[i]) - 48
		checksum += cn

		if bOdd {
			checksum += DELTA[cn]
		}
		bOdd = !bOdd
	}
	if checksum%10 == 0 {
		return true
	}

	return false
}
