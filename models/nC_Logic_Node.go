// @Title  nC_Logic_Node
// @Description  逻辑节点的相关处理
package models

import (
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"strings"
)

//操作节点
type NC_Node struct {
	Name     string      //节点名称
	Mark     string      //节点地址
	Type     NodeType    //节点类型
	Declare  interface{} //声明内容
	InLines  []NC_Flow   //入线
	OutLines []NC_Flow   //出线
}

type NodeType string

const (
	AssignNode   NodeType = "assign"   //赋值节点
	VariableNode NodeType = "variable" //定义节点
	CycleNode    NodeType = "cycle"    //循环节点
	JudgeNode    NodeType = "judge"    //判断节点
	LogicNode    NodeType = "logic"    //逻辑节点
	FormNode     NodeType = "form"     //表单处理节点
)

//定义节点
type NC_VariableNode struct {
	Mark       string        //地址编号
	Properties []NC_Property //类型结构属性
}

//单个定义的变量
type NC_Property struct {
	Name         string        //名称
	Mark         string        //标识
	Type         string        //类型
	Multiple     bool          //多个
	InitialValue string        //初始值
	Properties   []NC_Property //类型结构属性
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
type NC_Flow struct {
	From   string
	To     string
	Judges []Judge
}

// TODO NC_Flow-获取判断条件源码
func (flow *NC_Flow) getJudgeCode() string {
	codes := []string{}
	for _, judge := range flow.Judges {
		codes = append(codes, judge.getCode())
	}

	fmt.Println("flow code is ", strings.Join(codes, "||"))
	return strings.Join(codes, "||")
}

//判断条件
type Judge struct {
	Param     string
	Condition string //条件
	Value     string
	SubJudges []Judge //子（且）判断
}

// TODO NC_Flow-获取判断源码
func (judge *Judge) getCode() string {
	code := strings.Join([]string{judge.Param, judge.Condition, judge.Value}, " ")
	if len(judge.SubJudges) > 0 {
		subCodes := []string{}
		for _, subJudge := range judge.SubJudges {
			subCodes = append(subCodes, subJudge.getCode())
		}
		code = fmt.Sprintf("((%s) && (%s))", code, strings.Join(subCodes, "||"))
	}
	fmt.Println("judge code is ", code)
	return code
}

// TODO NC_Node-获取赋值节点源码assign
func (node *NC_Node) getAssignCode() string {
	code := ""
	var split []string
	assigns := []Assign{}
	utils.ReUnmarshal(node.Declare, &assigns)
	for _, assign := range assigns {
		split = append(split, assign.Key+" = "+assign.Value)
	}
	code = strings.Join(split, "\t\n") + "\t\n"
	return code
}

// TODO NC_Node-获取定义节点源码variable
func (node *NC_Node) getVariableCode() string {
	code := "var " + node.Mark + " " + getStructName(node.Mark) + "\n"
	variables := []NC_Property{}
	utils.ReUnmarshal(node.Declare, &variables)
	for _, p := range variables {
		if p.InitialValue != "" {
			code += node.Mark + "." + p.Mark + "=" + p.InitialValue + "\n"
		}
	}
	return code
}

// TODO NC_Node-获取循环节点源码cycle
func (node *NC_Node) getCycleCode(flow NC_Flow) string {
	judgeCode := fmt.Sprintf("if %s == false {\n\tbreak\n}\n", flow.getJudgeCode())
	code := "for { \n" + judgeCode
	return code
}

// TODO NC_Node-获取判断节点源码judge
func (node *NC_Node) getJudgeCode(flow NC_Flow) string {
	code := fmt.Sprintf("if %s {", flow.getJudgeCode())
	return code
}

// TODO NC_Node-获取逻辑节点源码logic
func (node *NC_Node) getLogicCode() string {
	code := ""
	return code
}

// TODO NC_Node-获取表单节点源码form
func (node *NC_Node) getFormCode() string {
	code := ""
	return code
}
