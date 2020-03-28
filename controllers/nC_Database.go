package controllers

import "com.waschild/noChaos-Server/models"

type DatabaseController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (c *DatabaseController) Create() { c.CreateWithModel(&models.NC_Database{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *DatabaseController) GetMany() {
	c.GetManyWithModel(&models.NC_Database{}, &[]models.NC_Database{})
}

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *DatabaseController) GetDetail() { c.GetDetailWithModel(&models.NC_Database{}) }

// @Title	Update
// @Description 修改表单
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /update [Post]
func (c *DatabaseController) Update() { c.UpdateWithModel(&models.NC_Database{}) }

// @Title	Delete
// @Description 删除
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /delete [Post]
func (c *DatabaseController) Delete() { c.DeleteWithModel(&models.NC_Database{}) }

// @Title	GetDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDemo [Post]
func (c *DatabaseController) GetDemo() {
	model := models.NC_Database{}
	c.responseSuccess(map[string]interface{}{"model": model})
}
