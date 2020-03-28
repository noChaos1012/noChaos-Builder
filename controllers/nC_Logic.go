package controllers

import (
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
	//
	//var receivedLogic build.Logic
	//
	//err := json.Unmarshal(l.Ctx.Input.RequestBody, &receivedLogic)
	//if err != nil {
	//	fmt.Println("json.Unmarshal is err:", err.Error())
	//}
	//
	//fmt.Println(receivedLogic)
	//fmt.Println(utils.GetPinYin(receivedLogic.Name), receivedLogic.GetCode())
	//
	////
	//build.BuildLogic("testapp0316-1", utils.GetPinYin(receivedLogic.Name), receivedLogic.GetCode())
	//data, _ := json.Marshal(receivedLogic.GetCode())
	//l.Data["json"] = string(data)
	//l.ServeJSON()

}
