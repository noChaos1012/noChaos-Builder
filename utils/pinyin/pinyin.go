package pinyin

import (
	"regexp"
	"strings"
)

func prepare(s string) string {
	var re *regexp.Regexp

	re = regexp.MustCompile(`[a-zA-Z0-9_-]+`)
	s = re.ReplaceAllStringFunc(s, func(repl string) string {
		return "\t" + repl
	})

	re = regexp.MustCompile(`[^\p{Han}\p{P}\p{Z}\p{M}\p{N}\p{L}\t]`)
	s = re.ReplaceAllString(s, "")

	return s
}

func Romanize(s string) string {
	s = prepare(s)
	for i := 0; i < len(dict); i += 2 {
		s = strings.Replace(s, dict[i], dict[i+1], -1)
	}
	s = strings.Replace(s, "\t", " ", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.TrimSpace(s)
	return s
}

//// ToSlice 转换为字符串数组
//func toSlice(s string) []string {
//	var split []string
//	re := regexp.MustCompile(`[^a-zA-Z1-4]+`)
//	for _, str := range re.Split(s, -1) {
//		if str != "" {
//			split = append(split, str)
//		}
//	}
//	return split
//}
