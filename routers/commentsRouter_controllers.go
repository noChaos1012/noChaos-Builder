package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"],
        beego.ControllerComments{
            Method: "BuildForms",
            Router: `/buildForms`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"],
        beego.ControllerComments{
            Method: "BuildLogic",
            Router: `/buildLogic`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"],
        beego.ControllerComments{
            Method: "BuildServlet",
            Router: `/buildServlet`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:BuildController"],
        beego.ControllerComments{
            Method: "TestStruct",
            Router: `/testStruct`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"],
        beego.ControllerComments{
            Method: "CreateDB",
            Router: `/createDB`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DatabaseController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DeployController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DirectoryController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:DownLoadController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "Build",
            Router: `/build`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "GetFieldDemo",
            Router: `/getFieldDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "InitForm",
            Router: `/initForm`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:FormController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:LogicController"],
        beego.ControllerComments{
            Method: "Test",
            Router: `/testBuildAssignNode`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:Panduanceshi20041301_26"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:Panduanceshi20041301_26"],
        beego.ControllerComments{
            Method: "Panduanceshi20041301_26",
            Router: `/Panduanceshi20041301_26`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
        beego.ControllerComments{
            Method: "GetDemo",
            Router: `/getDemo`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/getDetail`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
        beego.ControllerComments{
            Method: "GetMany",
            Router: `/getMany`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
        beego.ControllerComments{
            Method: "Package",
            Router: `/package`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:ServletController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"] = append(beego.GlobalControllerRouter["com.waschild/noChaos-Server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
