package simulation

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetTime() string {
	now := time.Now()
	timestamp := now.UnixNano() / 1e6                                           //13位豪秒时间戳
	lRand := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000000) // 生成9位随机数
	ret := fmt.Sprintf("%v%v", timestamp, lRand)
	return ret
}

func GetUt(account string, pwdmd5 string, timemd5 string) string {
	const fixer = "loginEncryptionParameter"
	logininfo := account + pwdmd5 + fixer + timemd5
	var infomd5 string = GetMd5(logininfo)
	// 在时间戳md5 第10位 插入logininfo的md5加密结果
	ut := timemd5[:10] + infomd5 + timemd5[10:]
	return ut
}

func GetSubmitUt(tasksJson string, timemd5 string) string {
	const fixer = "submitLogToken"
	submitInfo := tasksJson + fixer + timemd5
	var infomd5 string = GetMd5(submitInfo)
	// 在时间戳md5 第10位 插入logininfo的md5加密结果
	ut := timemd5[:10] + infomd5 + timemd5[10:]
	return ut
}
