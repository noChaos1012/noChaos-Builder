package controllers

import (
	"com.waschild/noChaos-Server/build/logic"
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
	build.BuildLogic()
	l.Data["json"] = "build success"
	l.ServeJSON()
}
