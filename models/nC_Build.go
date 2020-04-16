// @Title  nC_Build
// @Description  构建代码的相关处理
package models

import (
	"com.waschild/noChaos-Server/utils"
	"fmt"
	"github.com/jinzhu/gorm"
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

//TODO NC_Servlet-构建models.go代码
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
	modelsCode = strings.Replace(modelsCode,
		"{{.DBConnection}}",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", servlet.Database.UserName, servlet.Database.Password, servlet.Database.Host, servlet.Database.Port, servlet.GetName(), servlet.Database.Charset),
		-1)
	modelsCode = strings.Replace(modelsCode,
		"{{.MigrateForms}}",
		func(servlet NC_Servlet) string {
			NCDB.Model(&servlet).Related(&servlet.Forms, "ServletId")
			codeArr := []string{}
			for _, form := range servlet.Forms {
				if form.IsStore {
					codeArr = append(codeArr, "&"+form.GetName()+"{}")
				}
			}
			return strings.Join(codeArr, ",")
		}(servlet),
		-1)

	WriteAndFormat(path.Join(utils.DeployPath, servlet.GetName(), "models"), "models.go", modelsCode)
	return true, ""
}

//TODO NC_Servlet-构建router.go代码
func (servlet NC_Servlet) BuildRouter(version, description string) {
	var routerCode = `
	// @APIVersion {{.version}}
	// @Description {{.description}}
	// @Title noChaos API Service
	package routers
	
	{{.Package}}

	func init() {
		{{.Namespace}}
	}
	`
	routerCode = strings.Replace(routerCode, "{{.version}}", version, -1)
	routerCode = strings.Replace(routerCode, "{{.description}}", description, -1)

	routerCode = strings.Replace(routerCode, "{{.Package}}", func(servlet NC_Servlet) string {
		pkgCode := `	import (
				"noChaos-Server_Data/{{.ServletName}}/controllers"
				"github.com/astaxie/beego"
			)`
		return strings.Replace(pkgCode, "{{.ServletName}}", servlet.GetName(), -1)
	}(servlet), -1)

	routerCode = strings.Replace(routerCode, "{{.Namespace}}",
		func(s NC_Servlet) string {
			nsCode := `
			ns := beego.NewNamespace("{{.servletName}}",
				{{.namespace}}
			)
			beego.AddNamespace(ns)
			`
			nsCode = strings.Replace(nsCode, "{{.servletName}}", s.GetName(), -1)
			rootDir := NC_Directory{}
			rootDir.ServletId = s.ID
			rootDir.ID = 0
			nsCode = strings.Replace(nsCode, "{{.namespace}}", rootDir.GetNamespaceCode(), -1)
			return nsCode
		}(servlet), -1)
	WriteAndFormat(path.Join(utils.DeployPath, servlet.GetName(), "routers"), "router.go", routerCode)
}

//TODO NC_Build-构建服务代码
func (build *NC_Build) BuildServlet() (bool, string) {

	servlet := NC_Servlet{}
	servlet.ID = build.ServletId
	NCDB.First(&servlet)
	NCDB.Where(&NC_Database{ID: servlet.DatabaseID}).First(&servlet.Database)
	NCDB.Model(&servlet).Related(&servlet.Forms, "ServletId")
	if servlet.DatabaseID > 0 && len(servlet.Forms) > 0 {
		for _, form := range servlet.Forms {
			form.BuildForm(servlet.GetName())
		}
		servlet.BuildModels()
	}
	NCDB.Model(&servlet).Related(&servlet.Logics, "ServletId")
	if len(servlet.Logics) > 0 {
		for _, logic := range servlet.Logics {
			logic.BuildLogic(servlet.GetName())
		}
		servlet.BuildRouter(build.Version, build.VersionDescription)
	}
	return true, "构建成功"
}

//TODO NC_Logic-构建逻辑logic代码
func (logic NC_Logic) BuildLogic(servletName string) {
	logic.CompileProperties() //组装属性
	fmt.Println(logic.GetCode())
	WriteAndFormat(path.Join(utils.DeployPath, servletName, "controllers"), logic.GetName()+".go", logic.GetCode())
}

//TODO NC_Form-构建表单form代码
func (form NC_Form) BuildForm(servletName string) {
	NCDB.Model(&form).Related(&form.Fields, "FormId")
	WriteAndFormat(path.Join(utils.DeployPath, servletName, "models"), utils.GetPinYin(form.Name)+".go", form.GetCode())
}
