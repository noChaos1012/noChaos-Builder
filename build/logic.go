package build

import (
	"com.waschild/noChaos-Server/models/noChaos"
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

var controllerCode = `
package controllers

func {{.methodName}} ({{.inParams}}) {{.outParams}}{
{{.logic}}
return
}
`

type Logic struct {
	Name    string     `json:"name"`
	Inputs  []Variable `json:"inputs"`  //输入变量
	Outputs []Variable `json:"outputs"` //输出变量
	Nodes   []Node     `json:"nodes"`   //运算节点
}

//获取代码
func (l Logic) GetCode() string {
	code := strings.Replace(controllerCode, "{{.methodName}}", utils.GetEnglish(l.Name), -1)
	code = strings.Replace(code, "{{.inParams}}", l.GetInParams(), -1)
	code = strings.Replace(code, "{{.outParams}}", l.GetOutParams(), -1)

	nodeCode := ""
	for _, node := range l.Nodes {
		nodeCode += node.GetCode()
	}
	code = strings.Replace(code, "{{.logic}}", nodeCode, -1)
	return code
}

//获取输入参数
func (l Logic) GetInParams() string {
	var split []string
	for _, v := range l.Inputs {
		split = append(split, v.Name+" "+v.Type)
	}
	return strings.Join(split, ",")
}

//获取输出参数
func (l Logic) GetOutParams() string {

	var split []string
	for _, v := range l.Outputs {
		split = append(split, v.Name+" "+v.Type)
	}
	if len(split) == 0 {
		return "" //当长度为空不返回
	}
	return "(" + strings.Join(split, ",") + ")"
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
		return node.GetAssignCode()
	case "variable":
		return GetVariableCode(node.Declare)

	default:
		return code
	}
}

//获取赋值节点
func (node Node) GetAssignCode() string {

	var split []string
	assigns := []Assign{}
	utils.ReUnmarshal(node.Declare, &assigns)
	for _, assign := range assigns {
		split = append(split, assign.GetCode())
	}
	//fmt.Printf("aa 反射的类型是 %v \n", reflect.TypeOf(aa))
	return strings.Join(split, "\t\n") + "\t\n"
}

//赋值
type Assign struct {
	Key   string `json:"key"`
	Value string `json:"value"` //数据类型可能有多种
}

//获取赋值函数
func (assign Assign) GetCode() string {
	return assign.Key + " = " + assign.Value
}

//获取定义变量节点
func GetVariableCode(declare interface{}) string {
	var split []string
	variables := []Variable{}
	utils.ReUnmarshal(declare, &variables)
	for _, variable := range variables {
		split = append(split, variable.Var())
	}
	return strings.Join(split, "\t\n") + "\t\n"
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

//定义
func (variable Variable) Var() string {
	return "var " + variable.Name + " " + variable.Type
}

//以控制器创建逻辑方法
func BuildLogic(servletName string, logicName string, code string) {
	controllerDir := path.Join(noChaos.DeployPath, servletName, "controllers") //控制器文件夹
	utils.WriteToFile(path.Join(controllerDir, logicName+".go"), code)
	gofmtCMD := "gofmt -w  {{.logicName}}.go"
	gofmtCMD = strings.Replace(gofmtCMD, "{{.logicName}}", logicName, -1)
	cmd := exec.Command("/bin/bash", "-c", "cd "+controllerDir+";"+gofmtCMD)

	fmt.Println(cmd.Run())
}
