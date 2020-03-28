package controllers

import (
	"com.waschild/noChaos-Server/build"
	"com.waschild/noChaos-Server/models"
)

type ServletController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (c *ServletController) Create() { c.CreateWithModel(&models.NC_Servlet{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *ServletController) GetMany() {
	c.GetManyWithModel(&models.NC_Servlet{}, &[]models.NC_Servlet{})
}

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *ServletController) GetDetail() { c.GetDetailWithModel(&models.NC_Servlet{}) }

// @Title	Update
// @Description 修改表单
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /update [Post]
func (c *ServletController) Update() { c.UpdateWithModel(&models.NC_Servlet{}) }

// @Title	Delete
// @Description 删除
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /delete [Post]
func (c *ServletController) Delete() { c.DeleteWithModel(&models.NC_Servlet{}) }

// @Title PackageServlet
// @Description 打包服务
// @Param	name	string	true		"服务名称"
// @Param	mode	string	true		"模式（mac/linux）"
// @Success 200 打包成功
// @Failure 403 body is empty
// @router /package [Post]
func (s *ServletController) Package() {
	build.PackageServlet(s.GetString("name"), s.GetString("mode"))
	s.Data["json"] = "package servlet success"
	s.ServeJSON()
}

// @Title	GetDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDemo [Post]
func (c *ServletController) GetDemo() {
	model := models.NC_Servlet{}
	c.responseSuccess(map[string]interface{}{"model": model})
}
