// Package: memberCard
// File: trim.go
// Created by mint
// Useage: 格式化
// DATE: 14-7-3 19:45
package memberCard

import (
	"unicode"
	"fmt"
)

// 展示型的数据 ---> 系统内部处理形式
// xxxx xxxx xxxx xxxx xxx --->  xxxxxxxxxxxxxxxxxxx
// xxxx-xxxx-xxxx-xxxx-xxx --->  xxxxxxxxxxxxxxxxxxx
// ......
func FmtToInside(strCard string) (CardNumber, bool) {
	if strCard == "" {
		return CardNumber(""), false
	}

	runeCard := []rune(strCard)
	runeCard = fmtToDigit(runeCard)

	return CardNumber(runeCard), true
}

func fmtToDigit(runeCard []rune) []rune {
	card := make([]rune, 0)
	for _, byte := range runeCard {
		if isDigit(byte) {
			card = append(card, byte)
		}
	}

	return card
}

//判断一个字符是否是10进制数值
func isDigit(digit rune) bool {
	return unicode.IsDigit(digit)
}

//======================================================================================================================
//会员卡展示形式
type EMemberCardDisPlayType uint16

const (
	EMemberCardDisPlayType_SPACE   EMemberCardDisPlayType = iota
	EMemberCardDisplayType_DASH
	EMemberCardDisplayType_COMMA
)

// 系统内部处理形式  --->  展示形式
// xxxxxxxxxxxxxxxxxxx ---> xxxx xxxx xxxx xxxx xxx
// xxxxxxxxxxxxxxxxxxx ---> xxxx-xxxx-xxxx-xxxx-xxx
// ......
func FmtToDisplay(card CardNumber, displayType EMemberCardDisPlayType) (string, bool) {
	//TODO 会员卡的展示形式扩展

	switch displayType{
	case EMemberCardDisPlayType_SPACE:
		return fmt.Sprintf("%s %s %s %s %s", card[0:4], card[4:8], card[8:12], card[12:16], card[16:]), true
	case EMemberCardDisplayType_DASH:
		return fmt.Sprintf("%s-%s-%s-%s-%s", card[0:4], card[4:8], card[8:12], card[12:16], card[16:]), true
	case EMemberCardDisplayType_COMMA:
		return fmt.Sprintf("%s,%s,%s,%s,%s", card[0:4], card[4:8], card[8:12], card[12:16], card[16:]), true
	default:
		return string(card), true
	}

	return "", false
}



