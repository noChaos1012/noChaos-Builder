package models

import (
	"com.waschild/noChaos-Server/utils"
	"strconv"
	"strings"
)

type NC_Logic struct {
	ID        uint `gorm:"primary_key"`
	ServletId uint
	DirId     uint       //文件夹ID
	Name      string     `json:"name"`
	Inputs    []Variable `json:"inputs"`  //输入变量
	Outputs   []Variable `json:"outputs"` //输出变量
	Nodes     []Node     `json:"nodes"`   //运算节点
	Flows     []Flow     `json:"flows"`   //流向
}

var controllerCode = `
package controllers

func {{.methodName}} ({{.inParams}}) {{.outParams}}{
{{.logic}}
return
}
`

//获取代码
func (l NC_Logic) GetCode() string {
	code := strings.Replace(controllerCode, "{{.methodName}}", utils.GetPinYin(l.Name), -1)
	code = strings.Replace(code, "{{.inParams}}",
		func(vars []Variable) string {
			var split []string
			for _, v := range vars {
				split = append(split, v.Name+" "+v.Type)
			}
			return strings.Join(split, ",")
		}(l.Inputs), -1)
	code = strings.Replace(code, "{{.outParams}}",
		func(vars []Variable) string {
			var split []string
			for _, v := range vars {
				split = append(split, v.Name+" "+v.Type)
			}
			if len(split) == 0 {
				return "" //当长度为空不返回
			}
			return "(" + strings.Join(split, ",") + ")"
		}(l.Outputs), -1)

	nodeCode := ""
	for _, node := range l.Nodes {
		nodeCode += node.GetCode()
	}
	code = strings.Replace(code, "{{.logic}}", nodeCode, -1)
	return code
}

//数据库名称
func (logic NC_Logic) GetName() string {
	return utils.GetPinYin(logic.Name) + "_nc_" + strconv.Itoa(int(logic.ID))
}

//操作节点
type Node struct {
	Name    string      `json:"name"`    //节点名称
	Type    string      `json:"type"`    //节点类型
	Declare interface{} `json:"declare"` //声明
}

//获取节点代码
func (node Node) GetCode() string {
	code := ""
	switch node.Type {
	case "assign":
		return func(declare interface{}) string {
			var split []string
			assigns := []Assign{}
			utils.ReUnmarshal(declare, &assigns)
			for _, assign := range assigns {
				split = append(split, assign.Key+" = "+assign.Value)
			}
			return strings.Join(split, "\t\n") + "\t\n"
		}(node.Declare)

	case "variable":
		return func(declare interface{}) string {
			var split []string
			variables := []Variable{}
			utils.ReUnmarshal(declare, &variables)
			for _, variable := range variables {
				split = append(split, "var "+variable.Name+" "+variable.Type)
			}
			return strings.Join(split, "\t\n") + "\t\n"
		}(node.Declare)

	default:
		return code
	}
}

//赋值
type Assign struct {
	Key   string `json:"key"`
	Value string `json:"value"` //数据类型可能有多种
}

//变量
type Variable struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

//声明
func (variable Variable) Declare() string {
	return variable.Type + " " + variable.Name
}

//顺序流向
type Flow struct {
	From   string
	To     string
	Judges [][]Judge
}

//判断条件
type Judge struct {
	Param     string
	Condition interface{} //条件
}
