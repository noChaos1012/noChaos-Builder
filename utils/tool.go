package utils

import (
	"com.waschild/noChaos-Server/utils/pinyin"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
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

//重新解析json
func ReUnmarshal(inValue, outValue interface{}) {
	data, _ := json.Marshal(inValue)
	err := json.Unmarshal(data, outValue)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
}

//获取首字母大写的中文拼音
func GetPinYin(chinese string) string {
	preWord := pinyin.Romanize(chinese)
	strArry := []rune(preWord)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

//获取首字母小写的拼音
func GetPinYinLittle(chinese string) string {
	preWord := pinyin.Romanize(chinese)
	strArry := []rune(preWord)
	if strArry[0] >= 71 && strArry[0] <= 96 {
		strArry[0] += 32
	}
	return string(strArry)
}

//获取首字母大写的英文
func GetTitle(s string) string {
	strArry := []rune(s)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

//获取字符串中的中文连续词语，以(分割
func GetChinese(s string) []string {
	strArry := []rune(s)
	chineseArry := []rune{}
	for _, char := range strArry {
		if unicode.Is(unicode.Han, char) || char == 40 {
			chineseArry = append(chineseArry, char)
		}
	}
	return strings.Split(string(chineseArry), "(")
}
