package controllers

import "com.waschild/noChaos-Server/models"

type DeployController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (c *DeployController) Create() { c.CreateWithModel(&models.NC_Deploy{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *DeployController) GetMany() { c.GetManyWithModel(&models.NC_Deploy{}, &[]models.NC_Deploy{}) }

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *DeployController) GetDetail() { c.GetDetailWithModel(&models.NC_Deploy{}) }

// @Title	Delete
// @Description 删除
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /delete [Post]
func (c *DeployController) Delete() { c.DeleteWithModel(&models.NC_Deploy{}) }
