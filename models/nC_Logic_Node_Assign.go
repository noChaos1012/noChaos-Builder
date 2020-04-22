// @Title  nC_Logic_Node_Assign
// @Description  赋值节点的相关处理
package models

import (
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//赋值
type Assign struct {
	Key        string   //被赋值参数
	Value      string   //赋值公式
	Expression []string //各个赋值表达式
}

//TODO Assign-解析表达式并获取源码
func (as *Assign) AnalyzeExpression() string {
	expression := as.Value

	//替换中文方法名
	for _, chinese := range utils.GetChinese(as.Value) {
		//去除空串
		if len(chinese) > 0 {
			expression = strings.Replace(expression, chinese, MethodDict[chinese], 1)
		}
	}
	//替换参数
	rex := regexp.MustCompile(`{{.*?}}`)
	out := rex.FindAllStringSubmatch(expression, -1) //以空格区分
	for _, i := range out {
		for _, j := range i {
			p := strings.TrimPrefix(string(j), "{{")
			p = strings.TrimRight(p, "}}")
			p = utils.GetTitle(p)
			expression = strings.Replace(expression, string(j), p, 1)
		}
	}
	ex := utils.GetTitle(as.Key) + "=" + as.AssembleExpressions(expression)
	as.Expression = append(as.Expression, ex)
	return strings.Join(as.Expression, "\n")
}

//TODO Assign-组装表达式
func (as *Assign) AssembleExpressions(formula string) string {
	if strings.Contains(formula, "(") {
		endPoint := strings.Index(formula, ")")
		startPoint := strings.LastIndex(formula[:endPoint], "(")
		leftFormula := formula[:startPoint]
		name := strings.ReplaceAll(as.Key, ".", "") + "_" + strconv.Itoa(len(as.Expression)+1)
		fmt.Println(startPoint, endPoint, formula, leftFormula)
		//如果还有括号或者是开头
		if strings.Contains(leftFormula, "(") || startPoint == 0 {
			startPoint = strings.LastIndex(leftFormula, "(")
			as.Expression = append(as.Expression, name+":="+formula[startPoint+1:endPoint+1])
			left := formula[:startPoint+1]
			right := formula[endPoint+1:]
			formula = left + name + right
			return as.AssembleExpressions(formula)
		}
		as.Expression = append(as.Expression, name+":="+formula[startPoint+1:endPoint+1])
		left := formula[:startPoint]
		right := formula[endPoint+1:]
		formula = left + name + right
		return formula
	} else {
		return formula
	}
}
