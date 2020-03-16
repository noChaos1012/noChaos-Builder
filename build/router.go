package build

import (
	"com.waschild/noChaos-Server/utils"
	"path"
	"strings"
)

var router_Code = `// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"{{.Appname}}/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
`

//更新路由
func updateRouter(appPath string, packPath string) {

	utils.WriteToFile(path.Join(appPath, "routers", "router.go"), strings.Replace(router_Code, "{{.Appname}}", packPath, -1))

}
