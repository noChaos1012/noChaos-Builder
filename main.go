package main

import (
	_ "com.waschild/noChaos-Server/routers"
	"com.waschild/noChaos-Server/utils"
	"github.com/astaxie/beego"
	"os"
)

func main() {

	os.MkdirAll(utils.DeployPath, 0755)
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "testBuildAssignNode" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
