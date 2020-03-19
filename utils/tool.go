package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5加密
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//获取双引号加内容的字符串
func ImportStr(str string) string {
	return `"` + str + `"`
}
