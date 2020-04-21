// @Title  nC_Logic_Expression
// @Description  逻辑赋值表达式处理
package models

//运算 以单个数据结果输出，需要进行初始化数据配置，以备前端调取使用
type Method struct {
	Type       string   //所属类型
	Name       string   //名称
	Expression string   //表达式
	InputTypes []string //输入类型
	OutPutType string   //输出类型
}

var MethodDict = map[string]string{
	//通用
	"长度": "len",
	//固定值
	"文本": "string",
	//转换
	"文本转整数":   "strconv.Atoi",
	"文本转布尔":   "strconv.ParseBool",  //返回两个参数，取第一个，第二个为判断是否成功
	"文本转小数":   "strconv.ParseFloat", //返回两个参数，取第一个，第二个为判断是否成功
	"文本转指定进制": "strconv.ParseInt",   //返回两个参数，取第一个，第二个为判断是否成功
	"整数转文本":   "strconv.Itoa",
	"布尔转文本":   "strconv.FormatBool",
	"小数转文本":   "strconv.FormatFloat",
	"指定进制转文本": "strconv.FormatInt",
	//文本
	"子文本数":     "strings.Count",
	"文本第一位置":   "strings.Index",
	"任一文本第一位置": "strings.IndexAny",
	"文本最后位置":   "strings.LastIndex",
	"任一文本最后位置": "strings.LastIndexAny",
	"拆分定额字符串":  "strings.SplitN", //不保留分隔符
	"拆分全部字符串":  "strings.Split",
	"分割定额字符串":  "strings.SplitAfterN",
	"分割全部字符串":  "strings.SplitAfter", //保留分隔符
	"空格分割字符串":  "strings.Fields",
	"是否包含":     "strings.Contains",
	"是否包含任一":   "strings.ContainsAny",
	"开头包含":     "strings.HasPrefix",
	"结尾包含":     "strings.HasSuffix",
	"重复拼接":     "strings.Repeat",
	"转换大写":     "strings.ToUpper", //HELLO WORLD
	"转换小写":     "strings.ToLower", //hello world
	"转换标题":     "strings.ToTitle", //Hello World
	"删除首尾文本":   "strings.Trim",
	"删除首尾空格":   "strings.TrimSpace",
	"删除头部文本":   "strings.TrimPrefix",
	"删除尾部文本":   "strings.TrimSuffix",
	"替换文本":     "strings.Replace",
	"替换定额文本":   "strings.Replace",
	"小写相等":     "strings.EqualFold",
	"高效相等":     "strings.Compare",
	"从前截取文本":   "", //a[2:]
	"从后截取文本":   "", //a[:2]
	//文本复数
	"拼接字符串": "strings.Join",
	//数组
	"添加元素":   "append",
	"获取元素":   "",       //a[1]
	"合并数组":   "append", //append({{A}},{{B}}...)
	"从前截取数组": "",       //a[2:]
	"从后截取数组": "",       //a[:2]
	//数学
	"随机数":  "rand.Intn",        //rand.Intn(3) math
	"绝对值":  "math.Abs(float64", //math.Abs(float64({{i}})))
	"向上取整": "math.Ceil",
	"向下取整": "math.Floor",
	"取余":   "math.Mod",
	"分离小数": "math.Modf", //取整数，取小数
	"幂":    "math.Pow",
	"十的次方": "math.Pow10",
	"开平方":  "math.Sqrt",
	"开立方":  "math.Cbrt",
	"圆周率":  "math.Pi",
}

//特殊方法 调用特定逻辑
//FieldsFunc(s string:f func(rune) bool) []string：以一个或多个满足f(rune)的字符为分隔符，将s切分成多个子串，结果中不包含分隔符本身。如果s中没有满足f(rune)的字符，则返回一个空列表。
//TrimLeftFunc(s string:f func(rune) bool) string：删除s头部连续的满足f(rune)的字符。
//TrimRightFunc(s string:f func(rune) bool) string：删除s尾部连续的满足f(rune)的字符。
//TrimFunc(s string:f func(rune) bool) string：删除s首尾连续的满足f(rune)的字符。
//IndexFunc(s string:f func(rune) bool) int：返回s中第一个满足f(rune) 的字符的字节位置。如果没有满足 f(rune) 的字符，则返回 -1。
//LastIndexFunc(s string:f func(rune) bool) int：返回s中最后一个满足f(rune)的字符的字节位置。如果没有满足 f(rune) 的字符，则返回 -1。
//ContainsRune(s string:r rune) bool：判断字符串s中是否包含字符r。
//IndexRune(s string:r rune) int ：返回字符r在字符串s中第一次出现的位置。如果找不到则返回-1。
//Map 将 s 中满足 mapping(rune) 的字符替换为 mapping(rune) 的返回值。
