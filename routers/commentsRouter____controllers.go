package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/create`,
			AllowHTTPMethods: []string{"Post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
		beego.ControllerComments{
			Method:           "Package",
			Router:           `/package`,
			AllowHTTPMethods: []string{"Post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
