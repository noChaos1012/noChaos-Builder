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

	// @router /{{.logicName}} [Post]
	func  (c *{{.logicName}}) {{.logicName}} (){
		var In T_In{{.logicID}}
		var Out T_Out{{.logicID}}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &In)
		if c.handlerErrOK(err) {
			c.LogicBody(&In,&Out)
			c.responseSuccess(map[string]interface{}{"output": Out})
		}
	}


	func (c *{{.logicName}}) LogicBody(In *T_In{{.logicID}},Out *T_Out{{.logicID}}) {
			{{.logic}}
	}
`

//获取代码
func (l NC_Logic) GetCode() string {
	code := strings.Replace(controllerCode, "{{.logicName}}", l.GetName(), -1)
	utils.GetPinYin(l.Name + strconv.Itoa(int(l.ID)))

	code = strings.Replace(code, "{{.inputStructCode}}", getStructCode(l.Input.Properties, utils.GetPinYin("T_In"+strconv.Itoa(int(l.ID)))), -1)
	code = strings.Replace(code, "{{.outputStructCode}}", getStructCode(l.Output.Properties, "T_Out"+strconv.Itoa(int(l.ID))), -1)
	code = strings.Replace(code, "{{.logicID}}", strconv.Itoa(int(l.ID)), -1)

	nodeCode := ""
	structCode := ""
	for _, node := range l.Nodes {
		nodeCode += l.getNodeCode(node)
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
	return utils.GetPinYin(logic.Name) + "_" + strconv.Itoa(int(logic.ID))
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
	Mark    string      //节点地址
	Type    string      //节点类型
	Declare interface{} //声明内容
}

//获取节点代码
func (logic *NC_Logic) getNodeCode(node Node) string {
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
			assignCode := strings.Join(split, "\t\n") + "\t\n"

			return assignCode
		}(node.Declare)

	case "variable":
		return "var " + node.Mark + " " + getStructName(node.Mark) + "\n"

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
		return getStructCode(variables, getStructName(node.Mark))
	}
	return code
}

//赋值
type Assign struct {
	Key   string
	Value string
}

func (as *Assign) KeyName() string {

	var split []string
	for i, val := range strings.Split(as.Key, ".") {
		if i > 0 {
			split = append(split, utils.GetPinYin(val))
		} else {
			split = append(split, val)
		}
	}
	return strings.Join(split, ".")
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

func getStructName(name string) string {
	return "t_" + utils.GetPinYin(name)
}

//定义节点
type NC_VariableNode struct {
	Mark       string        //地址编号
	Properties []NC_Property //类型结构属性
}

//单个定义的变量
type NC_Property struct {
	Name       string        //名称
	Mark       string        //标识
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

		pCode := utils.GetPinYin(p.Mark) + " "
		if p.Multiple {
			pCode += "[]"
		}

		switch p.Type {
		case "string", "int", "float", "bool":
			pCode += p.Type

		case "custom":
			pCode += getStructName(p.Mark)
			PrefixStructCode = getStructCode(p.Properties, getStructName(p.Mark))

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
