// @Title  nC_Logic
// @Description  逻辑的处理

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

	NodesJson interface{} `gorm:"type:text"`
	FlowsJson interface{} `gorm:"type:text"`

	Input  NC_VariableNode    //输入变量
	Output NC_VariableNode    //输出变量
	Nodes  map[string]NC_Node `gorm:"-"` //运算节点
	Flows  []NC_Flow          //运算节点
}

var Code = ""

//TODO LOGIC-组装逻辑体代码
func (l *NC_Logic) GetCode() string {
	var controllerCode = `
	package controllers
	import "encoding/json"
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
			{{.core}}
	}
	`

	code := strings.Replace(controllerCode, "{{.logicName}}", l.GetName(), -1)
	code = strings.Replace(code, "{{.logicID}}", strconv.Itoa(int(l.ID)), -1)
	code = strings.Replace(code, "{{.core}}", l.GetCoreCode(), -1)
	code = strings.Replace(code, "{{.variableStructCode}}", l.GetStructCode(), -1)

	regex, _ := regexp.Compile(`

`)
	code = regex.ReplaceAllString(code, `
`)
	return code
}

//TODO LOGIC-获取逻辑名称 eg:Luoji_1
func (logic *NC_Logic) GetName() string {
	return utils.GetPinYin(logic.Name) + "_" + strconv.Itoa(int(logic.ID))
}

//TODO LOGIC-编译属性，将json属性转换成对应属性
func (logic *NC_Logic) CompileProperties() {
	nodes := []NC_Node{}
	properties := []interface{}{&nodes, &logic.Flows}
	for i, jsonCode := range []interface{}{logic.NodesJson, logic.FlowsJson} {
		t, ok := jsonCode.([]uint8)
		if ok {
			err := json.Unmarshal([]byte(string(t)), properties[i])
			if err != nil {
				fmt.Println("json.Unmarshal is err:", err.Error())
			}
		}
	}

	logic.Nodes = map[string]NC_Node{}
	for _, node := range nodes {
		logic.Nodes[node.Mark] = node
	}

	for _, flow := range logic.Flows {
		fromNode := logic.Nodes[flow.From]
		fromNode.OutLines = append(fromNode.OutLines, flow)
		logic.Nodes[flow.From] = fromNode

		toNode := logic.Nodes[flow.To]
		toNode.InLines = append(toNode.InLines, flow)
		logic.Nodes[flow.To] = toNode
	}

	logic.Input.Mark = logic.Nodes["In"].Mark
	utils.ReUnmarshal(logic.Nodes["In"].Declare, &logic.Input.Properties)
	logic.Input.Mark = logic.Nodes["Out"].Mark
	utils.ReUnmarshal(logic.Nodes["Out"].Declare, &logic.Output.Properties)
	//初始化最高层上级节点
	logic.Nodes["Logic"] = NC_Node{}
}

//TODO LOGIC-获取逻辑中用到的结构体代码
func (logic *NC_Logic) GetStructCode() string {
	structCode := ""
	for _, node := range logic.Nodes {
		if node.Type == "variable" {
			properties := []NC_Property{}
			utils.ReUnmarshal(node.Declare, &properties)
			switch node.Mark {
			case "In":
				structCode += "\n" + recurseStructCode("T_In"+strconv.Itoa(int(logic.ID)), properties)
			case "Out":
				structCode += "\n" + recurseStructCode("T_Out"+strconv.Itoa(int(logic.ID)), properties)
			default:
				structCode += "\n" + recurseStructCode(getStructName(node.Mark), properties)
			}
		}
	}
	return structCode
}

//TODO LOGIC-获取逻辑体核心代码
func (logic *NC_Logic) GetCoreCode() string {
	return logic.AssembleNodes("In", "Logic", "")
}

//TODO LOGIC-组装各逻辑节点代码
func (logic *NC_Logic) AssembleNodes(currentMark, superiorMark, code string) string {

	if currentMark == superiorMark {
		//循环到初始节点
		return "}\n"
	}
	nextMark := ""
	currentNode := logic.Nodes[currentMark]
	if len(currentNode.InLines) > 1 {
		//有多个来源节点->判断的结束节点
		return "}\n"
	}
	if len(currentNode.OutLines) > 0 {
		//有下一个节点->普通节点
		nextMark = currentNode.OutLines[0].To
	} else {
		//没有下一个节点->输出节点
		return ""
	}

	switch currentNode.Type {
	case AssignNode:
		code += currentNode.getAssignCode()

	case VariableNode:
		if len(currentNode.InLines) > 0 && len(currentNode.OutLines) > 0 {
			code += currentNode.getVariableCode()
		}
	case JudgeNode:
		for _, flow := range currentNode.OutLines {
			if len(flow.Judges) == 0 {
				currentNode.Type = CycleNode
				break
			}
		}
		if currentNode.Type == CycleNode {
			//循环节点
			childMark := ""
			childFlow := NC_Flow{}
			for _, flow := range currentNode.OutLines {
				if len(flow.Judges) > 0 {
					//循环体内节点
					childMark = flow.To
					childFlow = flow
				} else {
					//出循环的节点
					nextMark = flow.To
				}
			}
			code += logic.AssembleNodes(childMark, currentMark, currentNode.getCycleCode(childFlow))
		} else {
			//判断节点
			for _, flow := range currentNode.OutLines {
				code += logic.AssembleNodes(flow.To, currentMark, currentNode.getJudgeCode(flow))
			}
			return code
		}
	case FormNode:
		code += currentNode.getFormCode()

	case LogicNode:
		code += currentNode.getLogicCode()
	}
	code += logic.AssembleNodes(nextMark, superiorMark, "\n") //继续执行下一个节点
	return code
}

//TODO 获取结构体名称 eg:t_Jiegou
func getStructName(name string) string {
	return "t_" + utils.GetPinYin(name)
}

//TODO 递归获取结构代码
func recurseStructCode(structName string, properties []NC_Property) string {
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
			PrefixStructCode = recurseStructCode(getStructName(p.Mark), p.Properties)
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
