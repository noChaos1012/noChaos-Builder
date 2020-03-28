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
func (c *DatabaseController) Create() { c.CreateWithModel(&models.NC_DataBase{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *DatabaseController) GetMany() {
	c.GetManyWithModel(&models.NC_DataBase{}, &[]models.NC_DataBase{})
}

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *DatabaseController) GetDetail() { c.GetDetailWithModel(&models.NC_DataBase{}) }

// @Title	Update
// @Description 修改表单
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /update [Post]
func (c *DatabaseController) Update() { c.UpdateWithModel(&models.NC_DataBase{}) }

// @Title	Delete
// @Description 删除
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /delete [Post]
func (c *DatabaseController) Delete() { c.DeleteWithModel(&models.NC_DataBase{}) }
