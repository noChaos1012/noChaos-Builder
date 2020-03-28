package controllers

import "com.waschild/noChaos-Server/models"

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
