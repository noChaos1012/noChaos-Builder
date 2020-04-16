package models

import (
	"com.waschild/noChaos-Server/utils"
	"github.com/jinzhu/gorm"
	"os"
	"path"
	"strconv"
	"strings"
)

//服务
type NC_Servlet struct {
	gorm.Model
	Name        string
	Description string
	Port        int         //服务端口号
	Database    NC_Database `gorm:"FOREIGNKEY:ServletId"`
	DatabaseID  uint
	FTP         NC_FTP `gorm:"FOREIGNKEY:ServletId"`
	FTPID       uint
	Deploys     []NC_Deploy `gorm:"FOREIGNKEY:ServletId"`
	Forms       []NC_Form   `gorm:"FOREIGNKEY:ServletId"`
	Logics      []NC_Logic  `gorm:"FOREIGNKEY:ServletId"`
}

//TODO NC_Servlet-初始化服务内容
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

	//main.go
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
	mainCode = strings.Replace(mainCode, "{{.ServletName}}", rootPath, -1)
	WriteAndFormat(appPath, "main.go", mainCode)

	//app.conf
	var confCode = `
	appname = {{.ServletName}}
	httpport = {{.Port}}
	runmode = dev
	autorender = false
	copyrequestbody = true
	EnableDocs = true
	sqlconn = 
	`
	confCode = strings.Replace(confCode, "{{.ServletName}}", servlet.GetName(), -1)
	confCode = strings.Replace(confCode, "{{.Port}}", strconv.Itoa(servlet.Port), -1)
	WriteAndFormat(path.Join(appPath, "conf"), "app.conf", confCode)

	//controllers.go
	var baseLogicCode = `
	package controllers
	import (
		"fmt"
		"github.com/astaxie/beego"
	)
	//统一处理控制器
	type NCController struct {
		beego.Controller
	}
	`
	baseLogicCodes := []string{
		baseLogicCode,
		ResponseSuccess,
		HandlerErrOK}
	WriteAndFormat(path.Join(appPath, "controllers"), "controllers.go", strings.Join(baseLogicCodes, "\n"))

	//数据库
	NCDB.Exec("Create Database If Not Exists " +
		servlet.GetName() +
		" Character Set UTF8 Collate utf8_general_ci ")
}

//TODO NC_Servlet-获取服务名称
func (servlet NC_Servlet) GetName() string {
	return utils.GetPinYin(servlet.Name) + "_" + strconv.Itoa(int(servlet.ID))
}
