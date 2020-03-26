package models

import (
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"os/exec"
	"path"
	"strings"
)

//服务
type NC_Servlet struct {
	gorm.Model
	Name        string
	Description string
	DataBase    NC_DataBase `gorm:"FOREIGNKEY:ServletId"`
	Deploys     []NC_Deploy `gorm:"FOREIGNKEY:ServletId"`
	Forms       []NC_Form   `gorm:"FOREIGNKEY:ServletId"`
}

var confCode = `appname = {{.Appname}}
httpport = {{.Port}}
runmode = {{.Mode}}
autorender = false
copyrequestbody = true
EnableDocs = true
sqlconn = 
`

var mainCode = `package main

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

var modelsCode = `
package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var NCDB *gorm.DB

func init() {
	var err error

	NCDB, err = gorm.Open("mysql", "{{.DBConnection}}")

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		NCDB.DB().SetMaxIdleConns(10)
		NCDB.DB().SetMaxOpenConns(100)
	}
	//defer NCDB.Close()
	NCDB.AutoMigrate({{.MigrateForms}})
}

//分页结构体
type NC_Page struct {
	PageNo     int
	PageSize   int
	TotalCount int
	List       interface{}
}

func PageUtil(count int, pageNo int, pageSize int, list interface{}) NC_Page {

	return NC_Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalCount: count,
		List:       list,
	}
}
`

//初始化服务内容
func (servlet NC_Servlet) Initialize() {

	appPath := utils.DeployPath + "/" + servlet.Name       //绝对路径
	rootPath := utils.PackageRootPath + "/" + servlet.Name //相对包内路径

	fmt.Println("appPath: " + appPath)
	os.MkdirAll(appPath, 0755)

	//创建文件夹
	os.Mkdir(path.Join(appPath, "conf"), 0755)
	os.Mkdir(path.Join(appPath, "controllers"), 0755)
	os.Mkdir(path.Join(appPath, "models"), 0755)
	os.Mkdir(path.Join(appPath, "routers"), 0755)
	os.Mkdir(path.Join(appPath, "tests"), 0755)

	utils.WriteToFile(path.Join(appPath, "main.go"), strings.Replace(mainCode, "{{.Appname}}", rootPath, -1))
	gofmtCMD := "gofmt -w  {{.File}}"
	cmd := exec.Command("/bin/bash", "-c", "cd "+appPath+";"+strings.Replace(gofmtCMD, "{{.File}}", "main.go", -1))
	cmd.Run()

	utils.WriteToFile(path.Join(appPath, "conf", "app.conf"), strings.Replace(confCode, "{{.Appname}}", servlet.Name, -1))
	cmd = exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "conf")+";"+strings.Replace(gofmtCMD, "{{.File}}", "app.conf", -1))
	cmd.Run()

	utils.WriteToFile(path.Join(appPath, "models", "models.go"), modelsCode)
	cmd = exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "models")+";"+strings.Replace(gofmtCMD, "{{.File}}", "models.go", -1))
	cmd.Run()
}

//获取同步表单代码
func (servlet NC_Servlet) GetMigrateFormsCode() string {

	NCDB.Debug().Model(&servlet).Related(&servlet.Forms, "ServletId")

	codeArr := []string{}
	for _, form := range servlet.Forms {

		if form.IsStore {
			codeArr = append(codeArr, "&"+form.Name+"{}")
		}
	}
	return strings.Join(codeArr, ",")
}
