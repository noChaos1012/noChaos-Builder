package controllers

import (
	"com.waschild/noChaos-Server/build"
	"com.waschild/noChaos-Server/models"
	"github.com/astaxie/beego"
)

type ServletController struct {
	beego.Controller
}

// @Title	CreatedServlet
// @Description 创建服务
// @Param	name	string	true		"服务名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (s *ServletController) Create() {

	servlet := models.NC_Servlet{}
	servlet.Name = s.GetString("name")
	models.NCDB.Create(&servlet)
	//servlet.Create()
	s.Data["json"] = "create servlet success"
	s.ServeJSON()
}

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
