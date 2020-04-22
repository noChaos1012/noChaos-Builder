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
	code := `
		Logic{{NodeMark}} := {{Package}}{{LogicName}}{}
		In{{NodeMark}} := {{Package}}T_In26{}
		{{inputAssigns}}
		Out{{NodeMark}} := {{Package}}T_Out26{}
		Logic{{NodeMark}}.LogicBody(&In{{NodeMark}},&Out{{NodeMark}})
		{{outputAssigns}}
	`
	if l_node.Package == "self" {
		l_node.Package = ""
	} else {
		l_node.Package += "."
	}
	code = strings.Replace(code, "{{NodeMark}}", NodeMark, -1)
	code = strings.Replace(code, "{{Package}}", l_node.Package, -1)
	code = strings.Replace(code, "{{LogicName}}", logic.GetName(), 1)

	inputAssigns := []string{}
	for _, assign := range l_node.InputAssigns {
		assign.Key = strings.Replace(assign.Key, "In", "In"+NodeMark, 1)
		inputAssigns = append(inputAssigns, assign.AnalyzeExpression())
	}
	code = strings.Replace(code, "{{inputAssigns}}", strings.Join(inputAssigns, "\n"), 1)
	outputAssigns := []string{}
	for _, assign := range l_node.OutputAssigns {
		assign.Value = strings.Replace(assign.Value, "Out", "Out"+NodeMark, 1)
		outputAssigns = append(outputAssigns, assign.AnalyzeExpression())
	}
	code = strings.Replace(code, "{{outputAssigns}}", strings.Join(outputAssigns, "\n"), 1)
	return code
}
