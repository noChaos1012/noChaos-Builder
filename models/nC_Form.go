package models

import (
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//数据结构
type NC_Form struct {
	ID        uint `gorm:"primary_key"`
	ServletId uint //服务ID
	DirId     uint //文件夹ID
	CreatedAt time.Time
	Name      string        //名称
	IsStore   bool          //需要存储
	Fields    []NC_Field    `gorm:"FOREIGNKEY:FormId"`     //字段
	Relations []NC_Relation `gorm:"FOREIGNKEY:FromFormId"` //关联
}

//数据结构 — 字段
type NC_Field struct {
	ID            uint `gorm:"primary_key"`
	FormId        uint
	Name          string
	Type          string //类型
	IsKey         bool   //主键
	AutoIncrement bool   //自增
	Index         bool   //索引
	Unique        bool   //唯一
	NotNull       bool   //非空
	Default       string //默认
	Size          int    //长度
	Maximum       int    //精度
	Decimal       int    //标度
	IsArray       bool   //是否是复数
}

//关联
type NC_Relation struct {
	Type       string //关联类型、一对多、多对多
	FromFormId uint   //来自对象
	ToFormId   uint   //关联对象
}

//类型字典
var TypeMap = map[string]map[string]string{}

var modelCode = `package models

type {{.ModelName}} struct{
	{{.ModelFields}}
}
`
var fieldCode = "{{.Name}} {{.Type}} {{.Tag}} "

func init() {
	goMap := map[string]string{
		"文本":   "string",
		"整数":   "int",
		"小数":   "float",
		"布尔":   "bool",
		"日期":   "string",
		"时间":   "string",
		"日期时间": "string",
		"时间戳":  "string",
	}
	sqlMap := map[string]string{
		"文本":   "varchar",
		"整数":   "int",   //"超大整数": "bigint",
		"小数":   "float", //"超长小数": "double",
		"布尔":   "int",
		"日期":   "date",
		"时间":   "time",
		"日期时间": "datetime",
		"时间戳":  "timestamp",
	}

	TypeMap["go"] = goMap
	TypeMap["sql"] = sqlMap

}

func (form NC_Form) GetCode() string {

	code := strings.Replace(modelCode, "{{.ModelName}}", utils.GetPinYin(form.Name), -1)

	codeArr := []string{}
	for _, f := range form.Fields {
		codeArr = append(codeArr, f.GetCode(form.IsStore))
	}
	code = strings.Replace(code, "{{.ModelFields}}", strings.Join(codeArr, "\n"), -1)

	fmt.Println(code)
	return code

}

//获取编码
func (f NC_Field) GetCode(IsStore bool) string {
	code := strings.Replace(fieldCode, "{{.Name}}", utils.GetPinYin(f.Name), -1)
	code = strings.Replace(code, "{{.Type}}", f.GetGoType(), -1)

	tagStr := `json:"` + utils.GetPinYin(f.Name) + `"`

	if !IsStore {
		tagStr = "`" + tagStr + "`"
		code = strings.Replace(code, "{{.Tag}}", tagStr, -1)
		return code
	}

	tagStr = tagStr + ` gorm:"column:` + utils.GetPinYin(f.Name) + `;`
	//tagStr := ` gorm:"column:`+ utils.GetPinYin(f.Name) + `";`
	tagStr = tagStr + f.GetSQLType()
	if f.IsKey {
		tagStr += "primary_key;"
	}
	if f.AutoIncrement {
		tagStr += "AUTO_INCREMENT;"
	}
	if f.Index {
		tagStr += "index;"
	}
	if f.Unique {
		tagStr += "unique;"
	}
	if f.NotNull {
		tagStr += "not null;"
	}

	if len(f.Default) > 0 {
		tagStr += `default:` + f.Default + `;`
	}

	if strings.HasSuffix(tagStr, ";") {
		tagStr = tagStr[:len(tagStr)-1] + `"`
	}
	tagStr = "`" + tagStr + "`"
	code = strings.Replace(code, "{{.Tag}}", tagStr, -1)
	return code
}

//获取程序字段类型
func (field NC_Field) GetGoType() string {
	fieldType, ok := TypeMap["go"][field.Type]
	if !ok {
		fieldType = field.Type
	}
	if field.Type == "整数" || field.Type == "小数" {
		fieldType += strconv.Itoa(field.Size)
		fmt.Println(fieldType, field)
	}

	//是数组
	if field.IsArray {
		fieldType = "[]" + fieldType
	}
	return fieldType
}

//获取数据库字段类型
func (field NC_Field) GetSQLType() string {
	fieldType, ok := TypeMap["sql"][field.Type]
	if !ok {
		return ""
	}
	switch field.Type {

	case "文本":
		if (field.Size <= 255) && (field.Size > 0) {
			fieldType = fieldType + "(" + strconv.Itoa(field.Size) + ")"
		}
		if (field.Size <= 65535) && (field.Size > 255) {
			fieldType = "text"
		}
		if (field.Size <= 16777215) && (field.Size > 65535) {
			fieldType = "mediumtext"
		}
		if (field.Size <= 4294967295) && (field.Size > 16777215) {
			fieldType = "longtext"
		}
	case "整数":
		if field.Size == 64 {
			fieldType = "bigint"
		}
	case "小数":
		if field.Size == 32 {
			fieldType = "float(" + strconv.Itoa(field.Maximum) + "," + strconv.Itoa(field.Decimal) + ")"
			fmt.Println(fieldType, field)
		}
		if field.Size == 64 {
			fieldType = "double(" + strconv.Itoa(field.Maximum) + "," + strconv.Itoa(field.Decimal) + ")"
		}
	}

	return "type:" + fieldType + ";"
}
