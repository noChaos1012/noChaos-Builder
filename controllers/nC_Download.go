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
func (c *DownLoadController) Create() { c.CreateWithModel(&models.NC_DeployDownload{}) }

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *DownLoadController) GetMany() {
	c.GetManyWithModel(&models.NC_DeployDownload{}, &[]models.NC_DeployDownload{})
}

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *DownLoadController) GetDetail() { c.GetDetailWithModel(&models.NC_DeployDownload{}) }

// @Title	GetDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDemo [Post]
func (c *DownLoadController) GetDemo() {
	model := models.NC_DeployDownload{}
	c.responseSuccess(map[string]interface{}{"model": model})
}
