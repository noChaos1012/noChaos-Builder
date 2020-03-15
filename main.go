package main

import (
	"com.waschild/noChaos-Server/models/noChaos"
	_ "com.waschild/noChaos-Server/routers"
	"os"

	"github.com/astaxie/beego"
)

func main() {

	os.MkdirAll(noChaos.DeployPath, 0755)

	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "test" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
