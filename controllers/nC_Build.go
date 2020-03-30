package controllers

import (
	"com.waschild/noChaos-Server/models"
	"encoding/json"
)

type BuildController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (c *BuildController) Create() { c.CreateWithModel(&models.NC_Build{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *BuildController) GetMany() { c.GetManyWithModel(&models.NC_Build{}, &[]models.NC_Build{}) }

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *BuildController) GetDetail() { c.GetDetailWithModel(&models.NC_Build{}) }

// @Title	GetDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDemo [Post]
func (c *BuildController) GetDemo() {
	model := models.NC_Build{}
	c.responseSuccess(map[string]interface{}{"model": model})
}

// @Title	BuildForm
// @Description 构建表单源码
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildForm [Post]
func (c *BuildController) BuildForm() {
	form := models.NC_Form{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if c.handlerErrOK(err) {
		form.BuildForm()
		c.responseSuccess(map[string]interface{}{"model": form})
	}
}

// @Title	BuildModelsGo
// @Description 构建服务models.go源码
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildModelsGo [Post]
func (c *BuildController) BuildModelsGo() {
	servlet := models.NC_Servlet{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &servlet)
	if c.handlerErrOK(err) {
		success, desc := servlet.BuildModels()
		if success {
			c.responseSuccess(map[string]interface{}{"model": servlet})
		} else {
			res := make(map[string]interface{})
			res["status"] = "failure"
			res["description"] = desc
			c.Data["json"] = res
			c.ServeJSON()
		}
	}
}

// @Title	BuildRouterGo
// @Description 构建服务router.go源码
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildRouterGo [Post]
func (c *BuildController) BuildRouterGo() {
	servlet := models.NC_Servlet{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &servlet)
	if c.handlerErrOK(err) {
		servlet.BuildRouter()
		c.responseSuccess(map[string]interface{}{"model": servlet})
	}
}

// @Title	BuildForms
// @Description 构建服务所有form源码
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildForms [Post]
func (c *BuildController) BuildForms() {
	servlet := models.NC_Servlet{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &servlet)
	if c.handlerErrOK(err) {
		models.NCDB.Debug().First(&servlet)
		models.NCDB.Debug().Model(&servlet).Related(&servlet.Forms, "ServletId")
		for _, form := range servlet.Forms {
			form.BuildForm()
		}
		c.responseSuccess(map[string]interface{}{"model": servlet})
	}
}

// @Title	BuildServlet
// @Description 构建服务源码
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildServlet [Post]
func (c *BuildController) BuildServlet() {
	//form := models.NC_Servlet{}
	//err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	//if c.handlerErrOK(err) {
	//	form.Build()
	//	c.responseSuccess(map[string]interface{}{"models": form})
	//}
}
