package models

import (
	"fmt"
	"strings"
)

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
	return strings.Join(codes, "||")
}

//判断条件
type Judge struct {
	Param     string
	Condition string  //条件
	Value     string  //取值表达式
	SubJudges []Judge //子（且）判断
}

// TODO NC_Flow-获取判断源码
func (judge *Judge) getCode()string {

	//valueAssign := Assign{}
	//valueAssign.Value = judge.Value
	//valueAssign.Key = judge.Param + "Value"
	//valueAssign.AnalyzeExpression() +
	code := ""

	code += strings.Join([]string{judge.Param, judge.Condition, judge.Value}, " ")
	if len(judge.SubJudges) > 0 {
		subCodes := []string{}
		for _, subJudge := range judge.SubJudges {
			subCodes = append(subCodes, subJudge.getCode())
		}
		code = fmt.Sprintf("((%s) && (%s))", code, strings.Join(subCodes, "||"))
	}
	return code
}
