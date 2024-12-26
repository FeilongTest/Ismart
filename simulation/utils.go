package simulation

import (
	"math/rand"
	"strings"
	"time"
)

// CompareCompare 字符串比较
func CompareCompare(src, dest string) bool {
	if strings.Compare(src, dest) == 0 {
		return true
	}
	return false
}

// RandInt64 生成随机数
func RandInt64(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// DealString 处理字符串
func DealString(str string) string {
	//括号中文变英文
	str = strings.ReplaceAll(str, "（", "(")
	str = strings.ReplaceAll(str, "）", ")")
	str = strings.TrimSpace(str)
	return str
}
