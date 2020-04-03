package models

import (
	"com.waschild/noChaos-Server/utils"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type NC_Logic struct {
	ID        uint `gorm:"primary_key"`
	ServletId uint
	DirId     uint //文件夹ID
	Name      string
	//Inputs    string //输入变量

	InputJson  interface{} `gorm:"type:text"`
	OutputJson interface{} `gorm:"type:text"`
	NodesJson  interface{} `gorm:"type:text"`
	FlowsJson  interface{} `gorm:"type:text"`

	Input  NC_VariableNode //输入变量
	Output NC_VariableNode //输出变量
	Nodes  []Node          //运算节点
	Flows  []Flow          //流向
}

var controllerCode = `
	package controllers
	
	import "encoding/json"

	{{.inputStructCode}}
	{{.outputStructCode}}
	{{.variableStructCode}}
	
	type {{.logicName}} struct {
		NCController
	}
	
	func  (c *{{.logicName}}) {{.logicName}} (){
	var inputNode t_inputNode
	var outputNode t_outputNode
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &inputNode)

	if c.handlerErrOK(err) {
		{{.logic}}
		c.responseSuccess(map[string]interface{}{"output": outputNode})
	}
	}
`

//获取代码
func (l NC_Logic) GetCode() string {
	code := strings.Replace(controllerCode, "{{.logicName}}", utils.GetPinYin(l.Name), -1)
	code = strings.Replace(code, "{{.inputStructCode}}", getStructCode(l.Input.Properties, "t_inputNode"), -1)
	code = strings.Replace(code, "{{.outputStructCode}}", getStructCode(l.Output.Properties, "t_outputNode"), -1)

	nodeCode := ""
	structCode := ""
	for _, node := range l.Nodes {
		nodeCode += node.getCode()
		structCode += "\n" + node.getStructCode()
	}
	code = strings.Replace(code, "{{.logic}}", nodeCode, -1)
	code = strings.Replace(code, "{{.variableStructCode}}", structCode, -1)
	//code = strings.Replace(code,"\n\n","\n",-1)

	regex, _ := regexp.Compile(`

`)
	code = regex.ReplaceAllString(code, `
`)

	return code
}

//数据库
func (logic NC_Logic) GetName() string {
	return utils.GetPinYin(logic.Name) + "_nc_" + strconv.Itoa(int(logic.ID))
}

//编译属性，将json属性转换成对应属性
func (logic *NC_Logic) CompileProperties() {
	properties := []interface{}{&logic.Nodes, &logic.Flows, &logic.Input, &logic.Output}
	for i, jsonCode := range []interface{}{logic.NodesJson, logic.FlowsJson, logic.InputJson, logic.OutputJson} {
		t, ok := jsonCode.([]uint8)
		if ok {
			err := json.Unmarshal([]byte(string(t)), properties[i])
			if err != nil {
				fmt.Println("json.Unmarshal is err:", err.Error())
			}
		}
	}
}

//操作节点
type Node struct {
	Name    string      //节点名称
	Index   int         //节点地址
	Type    string      //节点类型
	Declare interface{} //声明内容
}

//获取节点代码
func (node Node) getCode() string {
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
		return "var " + node.Name + " " + getStructName(node.Name) + "\n"

	default:
		return code
	}
}

//获取定义节点的结构代码
func (node Node) getStructCode() string {
	code := ""
	if node.Type == "variable" {
		variables := []NC_Property{}
		utils.ReUnmarshal(node.Declare, &variables)
		return getStructCode(variables, getStructName(node.Name))
	}
	return code
}

//赋值
type Assign struct {
	Key   string
	Value string
}

//顺序流向
type Flow struct {
	From   int
	To     int
	Judges [][]Judge
}

//判断条件
type Judge struct {
	Param     string
	Condition interface{} //条件
}

func getStructName(name string) string {
	return "t_" + utils.GetPinYin(name)
}

//定义节点
type NC_VariableNode struct {
	Index      int           //地址编号
	Properties []NC_Property //类型结构属性
}

//单个定义的变量
type NC_Property struct {
	Name       string        //名称
	Type       string        //类型
	Multiple   bool          //多个
	Properties []NC_Property //类型结构属性
}

//获取结构代码
func getStructCode(properties []NC_Property, structName string) string {

	code := `
	{{.PrefixStruct}}
	
	type {{.StructName}} struct{
		{{.Properties}}
	}`

	PrefixStructCode := ""
	PropertiesCode := ""
	for _, p := range properties {

		pCode := utils.GetPinYinLittle(p.Name) + " "
		if p.Multiple {
			pCode += "[]"
		}

		switch p.Type {
		case "string", "int", "float", "bool":
			pCode += p.Type

		case "custom":
			pCode += getStructName(p.Name)
			fmt.Println()
			PrefixStructCode = getStructCode(p.Properties, getStructName(p.Name))

		default:
			pCode += getStructName(p.Type)
		}
		PropertiesCode += "\n\t" + pCode
	}

	code = strings.Replace(code, "{{.StructName}}", structName, -1)
	code = strings.Replace(code, "{{.PrefixStruct}}", PrefixStructCode, -1)
	code = strings.Replace(code, "{{.Properties}}", PropertiesCode, -1)

	return code
}
