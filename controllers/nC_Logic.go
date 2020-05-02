package controllers

import (
	"com.waschild/noChaos-Server/models"
	"encoding/json"
	"fmt"
)

type LogicController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (c *LogicController) Create() {
	model := models.NC_Logic{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &model)

	if c.handlerErrOK(err) {

		NodesData, _ := json.Marshal(model.NodesJson)
		nodes := []models.NC_Node{}
		err := json.Unmarshal(NodesData, &nodes)
		if err != nil {
			fmt.Println("json.Unmarshal is err:", err.Error())
		}
		model.InputJson, _ = json.Marshal(nodes[0])
		model.OutputJson, _ = json.Marshal(nodes[1])
		model.NodesJson = NodesData
		model.FlowsJson, _ = json.Marshal(model.FlowsJson)

		models.NCDB.Create(&model)
		c.responseSuccess(map[string]interface{}{"model": model})
	}
}

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *LogicController) GetMany() {
	c.GetManyWithModel(&models.NC_Logic{}, &[]models.NC_Logic{})
}

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *LogicController) GetDetail() { c.GetDetailWithModel(&models.NC_Logic{}) }

// @Title	GetDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDemo [Post]
func (c *LogicController) GetDemo() {
	model := models.NC_Logic{}
	c.responseSuccess(map[string]interface{}{"model": model})
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

// @Title BuildLogic
// @Description 编译逻辑
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /testBuildAssignNode [Post]
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

// @Title	Update
// @Description 修改逻辑
// @Param	name	string	true	"逻辑名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /update [Post]
func (l *LogicController) Update() { l.UpdateWithModel(&models.NC_Servlet{}) }

// @Title	Delete
// @Description 删除逻辑
// @Param	name	string	true	"逻辑名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /delete [Post]
func (l *LogicController) Delete() { l.DeleteWithModel(&models.NC_Servlet{}) }
