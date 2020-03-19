package controllers

import (
	build "com.waschild/noChaos-Server/build/logic"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type LogicController struct {
	beego.Controller
}

// @Title BuildLogic
// @Description 编译逻辑
// @Success 200 编译成功
// @Failure 403 body is empty
// @router / [get]
func (l *LogicController) Get() {
	//build.BuildLogic()

	l.Data["json"] = "build success"
	l.ServeJSON()
}

type Name struct {
	Name string `json:"name"`
}

type Receive_Data struct {
	Data interface{} `json:"data"`
}

// @Title BuildLogic
// @Description 编译逻辑
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /test [Post]
func (l *LogicController) Test() {
	//build.BuildLogic()
	fmt.Println(l.GetString("aa"))
	//l.Data["json"] = "build success"
	//l.ServeJSON()
	//fmt.Println(l.Ctx.Input)
	//var receivedJson Receive_Data

	var receiveMethod build.Method
	//build.Action()

	//var receiveMethod map[string]interface{}

	err := json.Unmarshal(l.Ctx.Input.RequestBody, &receiveMethod)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	build.BuildMethod(receiveMethod.GetCode())
	data, _ := json.Marshal(receiveMethod)
	l.Data["json"] = string(data)
	l.ServeJSON()

}

//{
//"name": "testMethod",
//"inputs": [{
//"name": "a",
//"type": "string"
//}, {
//"name": "b",
//"type": "string"
//}],
//"outputs": [{
//"name": "a",
//"type": "string"
//}, {
//"name": "b",
//"type": "string"
//}],
//}
