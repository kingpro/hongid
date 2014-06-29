// Package: security
// File: rand.go
// Created by mint
// Useage: 随机数工具
// DATE: 14-6-27 23:30
package security

import (
	"bytes"
	"math/rand"
	"time"
)

//生成随机字符串
func RandomString(num int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < num; {
		if string(RandomInt(65, 90)) != temp {
			temp = string(RandomInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

//生成随机数字
func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
