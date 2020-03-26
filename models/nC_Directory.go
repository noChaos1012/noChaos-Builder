package models

//文件夹结构 用来存储内容
type NC_Directory struct {
	ID        uint         `gorm:"primary_key"`
	ServletId uint         //服务ID
	Name      string       //名称
	DirId     uint         //上级文件夹ID
	Type      string       //类型为表单或逻辑
	Logics    []NC_Logic   //
	Forms     []NC_Form    //
	Servlet   []NC_Servlet //
}

var routerCode = `
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"com.waschild/noChaos-Server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/logic",
			beego.NSInclude(
				&controllers.LogicController{},
			),
		),
		beego.NSNamespace("/form",
			beego.NSInclude(
				&controllers.FormController{},
			),
		),
		beego.NSNamespace("/servlet",
			beego.NSInclude(
				&controllers.ServletController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
`
