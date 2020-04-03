package models

import (
	"com.waschild/noChaos-Server/utils"
	"github.com/jinzhu/gorm"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

//服务
type NC_Servlet struct {
	gorm.Model
	Name        string
	Description string
	DataBase    NC_Database `gorm:"FOREIGNKEY:ServletId"`
	Deploys     []NC_Deploy `gorm:"FOREIGNKEY:ServletId"`
	Forms       []NC_Form   `gorm:"FOREIGNKEY:ServletId"`
	Logics      []NC_Logic  `gorm:"FOREIGNKEY:ServletId"`
}

//初始化服务内容
func (servlet NC_Servlet) Initialize() {

	appPath := utils.DeployPath + "/" + servlet.GetName()       //绝对路径
	rootPath := utils.PackageRootPath + "/" + servlet.GetName() //相对包内路径

	os.MkdirAll(appPath, 0755)

	//创建文件夹
	os.Mkdir(path.Join(appPath, "conf"), 0755)
	os.Mkdir(path.Join(appPath, "controllers"), 0755)
	os.Mkdir(path.Join(appPath, "models"), 0755)
	os.Mkdir(path.Join(appPath, "routers"), 0755)
	os.Mkdir(path.Join(appPath, "tests"), 0755)

	var mainCode = `package main

	import (
		_ "{{.ServletName}}/routers"
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

	utils.WriteToFile(path.Join(appPath, "main.go"), strings.Replace(mainCode, "{{.ServletName}}", rootPath, -1))
	gofmtCMD := "gofmt -w  {{.File}}"
	cmd := exec.Command("/bin/bash", "-c", "cd "+appPath+";"+strings.Replace(gofmtCMD, "{{.File}}", "main.go", -1))
	cmd.Run()

	var confCode = `appname = {{.ServletName}}
	httpport = {{.Port}}
	runmode = {{.Mode}}
	autorender = false
	copyrequestbody = true
	EnableDocs = true
	sqlconn = 
	`

	utils.WriteToFile(path.Join(appPath, "conf", "app.conf"), strings.Replace(confCode, "{{.ServletName}}", servlet.GetName(), -1))
	cmd = exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "conf")+";"+strings.Replace(gofmtCMD, "{{.File}}", "app.conf", -1))
	cmd.Run()

	var baseLogicCode = `
	package controllers
	
	import (
		"encoding/json"
		"fmt"
		"github.com/astaxie/beego"
	)
	
	//统一处理控制器
	type NCController struct {
		beego.Controller
	}
	
	// 处理是否解析成功
	func (ncc *NCController) handlerErrOK(err error) bool {
	
		if err != nil {
			fmt.Println("come to error ")
			res := make(map[string]interface{})
			res["status"] = "failure"
			res["description"] = "接收数据解析失败"
			ncc.Data["json"] = res
			panic(err)
			return false
		}
		return true
	}
	
	// 返回成功
	func (ncc *NCController) responseSuccess(result interface{}) {
	
		res := make(map[string]interface{})
		res["list"] = result
		res["status"] = "success"
		res["description"] = "操作成功"
	
		ncc.Data["json"] = res
		ncc.ServeJSON()
	}
	`
	utils.WriteToFile(path.Join(appPath, "controllers", "NCController.go"), baseLogicCode)
	cmd = exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "controllers")+";"+strings.Replace(gofmtCMD, "{{.File}}", "NCController.go", -1))
	cmd.Run()

	//数据库创建
	NCDB.Exec("Create Database If Not Exists " +
		servlet.GetName() +
		" Character Set UTF8 Collate utf8_general_ci ")
}

//数据库名称
func (servlet NC_Servlet) GetName() string {
	return utils.GetPinYin(servlet.Name) + "_nc_" + strconv.Itoa(int(servlet.ID))
}
