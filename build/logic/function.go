package build

import (
	"com.waschild/noChaos-Server/models/noChaos"
	"com.waschild/noChaos-Server/utils"
	"encoding/json"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

var methodCode = `
package controllers

func {{.methodName}} ({{.inParams}}) {{.outParams}}{
	{{.outParamsVar}}
	{{.logics}}
	{{.return}}
}
`

type Method struct {
	Name    string     `json:"name"`
	Inputs  []Variable `json:"inputs"`  //输入变量
	Outputs []Variable `json:"outputs"` //输出变量
	Nodes   []Node     `json:"nodes"`   //运算节点
}

//获取输入参数
func (method Method) GetInParams() string {
	code := ""
	for _, v := range method.Inputs {
		code += v.Name + " " + v.Type + " , "
	}
	if len(code) > 1 {
		code = string([]byte(code)[:len(code)-2])
	}
	return code
}

//获取输出参数
func (method Method) GetOutParams() string {
	code := ""
	for _, v := range method.Outputs {
		code += v.Type + " , "
	}
	if len(code) > 1 {
		code = string([]byte(code)[:len(code)-2])
	}
	if len(method.Outputs) > 1 {
		return "( " + code + " )"
	}
	return code
}

func (method Method) GetOutParamsVar() string {
	outputNode := VariableNode{variables: method.Outputs}
	return outputNode.GetCode()
}

func (method Method) GetReturn() string {
	if len(method.Outputs) < 1 {
		return " "
	}

	code := "return "
	for _, v := range method.Outputs {
		code += v.Name + " , "
	}
	if len(method.Outputs) > 0 {
		code = string([]byte(code)[:len(code)-2])
	}
	return code
}

type Node struct {
	Name    string      `json:"name"`    //节点名称
	Type    string      `json:"type"`    //节点类型
	Declare interface{} `json:"declare"` //声明
}

//获取代码

func (node Node) GetCode() string {
	code := ""
	//fmt.Println(string(node.Declare.(type)))
	//switch  {
	//
	//}

	switch node.Type {
	case "assign":
		return node.GetAssignCode()
	//case "variable":
	//	var assignNode AssignNode
	//	assignNode.InitWithDeclare(node.Declare)
	//	return assignNode.GetCode()

	default:
		return code
	}
}

//获取赋值节点
func (node Node) GetAssignCode() string {
	code := ""
	assigns, _ := node.Declare.([]interface{})
	for _, assign := range assigns {
		newAssign, _ := assign.(map[string]interface{})
		code += newAssign["key"].(string) + " = " + newAssign["value"].(string) + "\n"
		//fmt.Printf("aa 反射的类型是 %v \n", reflect.TypeOf(aa))
	}
	return code
}

//赋值节点
type AssignNode struct {
	assigns []Assign
}

//获取代码
func (node AssignNode) GetCode() string {
	code := ""
	for _, assign := range node.assigns {
		code += assign.Key + " = " + assign.Value + "\n"
	}
	return code
}

//func (node AssignNode)GetAssignNodeCode() string{

//}

//赋值
type Assign struct {
	Key   string `json:"key"`
	Value string `json:"value"` //数据类型可能有多种
}

//获取代码
func (method Method) GetCode() string {
	code := methodCode
	code = strings.Replace(code, "{{.methodName}}", method.Name, -1)
	code = strings.Replace(code, "{{.inParams}}", method.GetInParams(), -1)
	code = strings.Replace(code, "{{.outParams}}", method.GetOutParams(), -1)
	code = strings.Replace(code, "{{.outParamsVar}}", method.GetOutParamsVar(), -1)

	nodeCode := ""
	for _, node := range method.Nodes {
		nodeCode += node.GetCode()
	}

	code = strings.Replace(code, "{{.logics}}", nodeCode, -1)
	code = strings.Replace(code, "{{.return}}", method.GetReturn(), -1)
	return code
}

//定义节点
type VariableNode struct {
	variables []Variable
}

//获取定义斌量
func (node VariableNode) GetCode() string {
	code := ""
	for _, v := range node.variables {
		code += v.Var() + "\n"
	}
	return code
}

type Variable struct {
	Name string `json:"name"`
	Type string `json:"type"`
	//Value interface{}
	//Label string //
}

//声明变量

func (variable Variable) Declare() string {
	return variable.Type + " " + variable.Name
}

func (variable Variable) Var() string {
	return "var " + variable.Name + " " + variable.Type
}

//测试使用
func initVariable(variableName, variableType string) Variable {
	variable := Variable{}
	variable.Name = variableName
	variable.Type = variableType
	return variable
}

func BuildMethod(code string) {
	appPath := noChaos.DeployPath + "/" + "testapp0316-1" //绝对路径
	utils.WriteToFile(path.Join(appPath, "controllers", "testMethod.go"), code)
	cmd := exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "controllers")+";"+"gofmt -w testMethod.go")
	fmt.Println(cmd.Run())
}

func Action() {

	//var a Variable
	//a.Name = "a"
	//a.Type = "string"
	//var b Variable
	//b.Name = "b"
	//b.Type = "string"
	a := initVariable("a", "string")
	b := initVariable("b", "string")

	method := Method{}
	method.Name = "testMethod"
	method.Inputs = []Variable{a, b}
	method.Outputs = []Variable{a, b}

	data, _ := json.Marshal(method)
	fmt.Println(string(data))

	//.Data["json"] = string(data)
	//l.ServeJSON()

	//fmt.Println(method.GetCode())

}
