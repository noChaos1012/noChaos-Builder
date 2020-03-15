package build

import (
	"com.waschild/noChaos-Server/models/noChaos"
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"os/exec"

	"os"
	"strings"

	path "path/filepath"
)

var conf_Code = `appname = {{.Appname}}
httpport = 8080
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true
sqlconn = {{.SQLConnStr}}
`

var main_Code = `package main

import (
	_ "{{.Appname}}/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
`

//初始化服务框架
func InitServlet(name string) {

	fmt.Println("buildServlet")
	appPath := noChaos.DeployPath + "/" + name
	fmt.Println("appPath: " + appPath)

	os.MkdirAll(appPath, 0755)

	//cmd := exec.Command("/bin/bash", "-c", "cd "+ appPath )
	//fmt.Println(cmd.Run())
	utils.WriteToFile(path.Join(appPath, "main.go"), strings.Replace(main_Code, "{{.Appname}}", name, -1))

	os.Mkdir(path.Join(appPath, "conf"), 0755)
	utils.WriteToFile(path.Join(appPath, "conf", "app.conf"), strings.Replace(conf_Code, "{{.Appname}}", name, -1))
	os.Mkdir(path.Join(appPath, "controllers"), 0755)
	os.Mkdir(path.Join(appPath, "models"), 0755)
	os.Mkdir(path.Join(appPath, "routers"), 0755)
	os.Mkdir(path.Join(appPath, "tests"), 0755)
}

//打包服务框架
func PackageServlet(name, mode string) {

	fmt.Println("buildServlet")
	appPath := noChaos.DeployPath + "/" + name

	beePackCode := ""
	switch mode {
	case "mac":
		beePackCode = "bee pack"
	case "linux":
		beePackCode = "bee pack -be GOOS=linux"

	default:
		beePackCode = ""
	}

	fmt.Println(appPath + beePackCode)

	cmd := exec.Command("/bin/bash", "-c", "cd "+appPath+";"+beePackCode)
	fmt.Println(cmd.Run())

}
