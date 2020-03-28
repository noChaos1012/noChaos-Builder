package controllers

import "com.waschild/noChaos-Server/models"

type DownLoadController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (c *DownLoadController) Create() { c.CreateWithModel(&models.NC_DownLoad{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *DownLoadController) GetMany() {
	c.GetManyWithModel(&models.NC_DownLoad{}, &[]models.NC_DownLoad{})
}

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *DownLoadController) GetDetail() { c.GetDetailWithModel(&models.NC_DownLoad{}) }
