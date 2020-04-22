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
		beego.NSNamespace("/dir",
			beego.NSInclude(
				&controllers.DirectoryController{},
			),
		),
		beego.NSNamespace("/db",
			beego.NSInclude(
				&controllers.DatabaseController{},
			),
		),
		beego.NSNamespace("/deploy",
			beego.NSInclude(
				&controllers.DeployController{},
			),
		),
		beego.NSNamespace("/build",
			beego.NSInclude(
				&controllers.BuildController{},
			),
		),
		beego.NSNamespace("/testBuildAssignNode",
			beego.NSNamespace("/testBuildAssignNode"), //beego.NSInclude(
			//&controllers.Ceshiluoji20{},
			//),

			beego.NSInclude(
			//&controllers.Ceshiluoji20{},
			),
		),
	)

	beego.AddNamespace(ns)
}
