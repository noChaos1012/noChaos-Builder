package controllers

import (
	"com.waschild/noChaos-Server/models"
)

type DirectoryController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (d *DirectoryController) Create() { d.CreateWithModel(&models.NC_Directory{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (d *DirectoryController) GetMany() {
	d.GetManyWithModel(&models.NC_Directory{}, &[]models.NC_Directory{})
}

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (d *DirectoryController) GetDetail() { d.GetDetailWithModel(&models.NC_Directory{}) }

// @Title	Update
// @Description 修改表单
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /update [Post]
func (d *DirectoryController) Update() { d.UpdateWithModel(&models.NC_Directory{}) }

// @Title	Delete
// @Description 删除表单
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /delete [Post]
func (d *DirectoryController) Delete() { d.DeleteWithModel(&models.NC_Directory{}) }
