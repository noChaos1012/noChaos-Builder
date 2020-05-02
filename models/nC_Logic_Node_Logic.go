// @Title  nC_Logic_Node_Logic
// @Description  逻辑节点的相关处理
package models

import (
	"strings"
)

type Logic_Node struct {
	Package       string   //包名
	LogicId       uint     //逻辑编号
	InputAssigns  []Assign //输入赋值
	OutputAssigns []Assign //输出赋值
}

//TODO Logic_Node-获取源码
func (l_node *Logic_Node) GetCode(NodeMark string) string {

	logic := NC_Logic{}
	logic.ID = l_node.LogicId
	NCDB.First(&logic)

	LogicMark := "Logic" + NodeMark
	code := `
		{{LogicMark}} := {{Package}}{{LogicName}}{}
		{{inputAssigns}}
		{{LogicMark}}.LogicBody()
		{{outputAssigns}}
	`
	packageName := ""
	if l_node.Package != "self" {
		packageName = l_node.Package + "."
	}
	code = strings.Replace(code, "{{LogicMark}}", LogicMark, -1)
	code = strings.Replace(code, "{{Package}}", packageName, -1)
	code = strings.Replace(code, "{{LogicName}}", logic.GetName(), 1)

	inputAssigns := []string{}
	for _, assign := range l_node.InputAssigns {
		assign.Key = strings.Replace(assign.Key, "In", LogicMark+".In", 1)
		inputAssigns = append(inputAssigns, assign.AnalyzeExpression())
	}
	code = strings.Replace(code, "{{inputAssigns}}", strings.Join(inputAssigns, "\n"), 1)
	outputAssigns := []string{}
	for _, assign := range l_node.OutputAssigns {
		assign.Value = strings.Replace(assign.Value, "Out", LogicMark+".Out", 1)
		outputAssigns = append(outputAssigns, assign.AnalyzeExpression())
	}
	code = strings.Replace(code, "{{outputAssigns}}", strings.Join(outputAssigns, "\n"), 1)
	return code
}
