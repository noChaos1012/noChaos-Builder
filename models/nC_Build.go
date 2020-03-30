package models

import (
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"os/exec"
	"path"
	"strings"
)

//构建程序代码
type NC_Build struct {
	gorm.Model
	ServletId          uint `gorm:"not null"`
	Version            string
	VersionDescription string
}

//构建form代码
func (form NC_Form) BuildForm() {

	NCDB.Debug().First(&form)
	NCDB.Debug().Model(&form).Related(&form.Fields, "FormId")

	servlet := NC_Servlet{}
	servlet.ID = form.ServletId
	NCDB.Debug().First(&servlet)
	//绝对路径
	appPath := utils.DeployPath + "/" + servlet.GetName()
	//写入文件
	utils.WriteToFile(path.Join(appPath, "models", utils.GetPinYin(form.Name)+".go"), form.GetCode())
	gofmtCMD := "gofmt -w  {{.File}}"
	cmd := exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "models")+";"+strings.Replace(gofmtCMD, "{{.File}}", utils.GetPinYin(form.Name)+".go", -1))
	cmd.Run()
}

//构建model代码
func (servlet NC_Servlet) BuildModels() (bool, string) {

	modelsCode := `
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

	NCDB.Debug().First(&servlet)
	NCDB.Debug().Model(&servlet).Related(&servlet.Forms, "ServletId")
	if len(servlet.Forms) == 0 {
		return false, "未配置表单数据"
	}

	db := NC_Database{}
	db.ServletId = servlet.ID
	NCDB.Where(&db).Find(&db)
	if db.Host == "" {
		return false, "未配置数据库"
	}

	modelsCode = strings.Replace(modelsCode,
		"{{.DBConnection}}",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", db.UserName, db.Password, db.Host, db.Port, servlet.GetName(), db.Charset),
		-1)

	modelsCode = strings.Replace(modelsCode,
		"{{.MigrateForms}}",
		func(servlet NC_Servlet) string {
			NCDB.Debug().Model(&servlet).Related(&servlet.Forms, "ServletId")
			codeArr := []string{}
			for _, form := range servlet.Forms {
				if form.IsStore {
					codeArr = append(codeArr, "&"+form.GetName()+"{}")
				}
			}
			return strings.Join(codeArr, ",")
		}(servlet),
		-1)

	appPath := utils.DeployPath + "/" + servlet.GetName()
	utils.WriteToFile(path.Join(appPath, "models", "models.go"), modelsCode)

	gofmtCMD := "gofmt -w  {{.File}}"
	cmd := exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "models")+";"+strings.Replace(gofmtCMD, "{{.File}}", "models.go", -1))
	cmd.Run()
	return true, ""
}

//构建程序源码
func (servlet NC_Servlet) BuildRouter() {
	var routerCode = `
	// @APIVersion 1.0.0
	// @Title beego Test API
	// @Description beego has a very cool tools to autogenerate documents for your API
	// @Contact astaxie@gmail.com
	// @TermsOfServiceUrl http://beego.me/
	// @License Apache 2.0
	// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
	package routers
	
	{{.Package}}
	
	func init() {
		{{.Namespace}}
	}
	`

	NCDB.Debug().First(&servlet)
	NCDB.Debug().Model(&servlet).Related(&servlet.Logics, "ServletId")

	routerCode = strings.Replace(routerCode, "{{.Package}}", func(logics []NC_Logic) string {
		if len(logics) > 0 {
			pkgCode := `	import (
				"noChaos-Server_Data/{{.ServletName}}/controllers"
				"github.com/astaxie/beego"
			)`
			return strings.Replace(pkgCode, "{{.ServletName}}", servlet.GetName(), -1)
		} else {
			return ""
		}
	}(servlet.Logics), -1)

	routerCode = strings.Replace(routerCode, "{{.Namespace}}",
		func(logics []NC_Logic) string {
			//ns := beego.NewNamespace("/v1",
			//	beego.NSNamespace("/object",
			//		beego.NSInclude(
			//			&controllers.ObjectController{},
			//		),
			//	),
			//)
			//beego.AddNamespace(ns)
			if len(logics) > 0 {
				return ""
			} else {
				return ""
			}
		}(servlet.Logics), -1)

	appPath := utils.DeployPath + "/" + servlet.GetName()

	fmt.Println(routerCode)

	utils.WriteToFile(path.Join(appPath, "routers", "router.go"), routerCode)

	gofmtCMD := "gofmt -w  {{.File}}"
	cmd := exec.Command("/bin/bash", "-c", "cd "+path.Join(appPath, "routers")+";"+strings.Replace(gofmtCMD, "{{.File}}", "router.go", -1))
	cmd.Run()

}

//以控制器创建逻辑方法
func (logic NC_Logic) BuildLogic() {

	NCDB.Debug().First(&logic)
	//NCDB.Debug().Model(&logic).Related(&logic.Nodes, "LogicId")

	servlet := NC_Servlet{}
	servlet.ID = logic.ServletId
	NCDB.Debug().First(&servlet)

	controllerDir := path.Join(utils.DeployPath, servlet.GetName(), "controllers") //控制器文件夹
	utils.WriteToFile(path.Join(controllerDir, logic.GetName()+".go"), logic.GetCode())
	gofmtCMD := "gofmt -w  {{.logicName}}.go"
	gofmtCMD = strings.Replace(gofmtCMD, "{{.logicName}}", logic.GetName(), -1)
	cmd := exec.Command("/bin/bash", "-c", "cd "+controllerDir+";"+gofmtCMD)

	fmt.Println(cmd.Run())
}
