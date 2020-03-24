package controllers

import (
	"com.waschild/noChaos-Server/models"
	"fmt"
	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

// @Title	CreatedServlet
// @Description 创建服务
// @Param	name	string	true		"服务名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (f *FormController) Create() {

	form := models.NC_Form{}
	form.Name = f.GetString("name")

	field := models.NC_Field{}
	field.Name = form.Name + "col_1"
	field.Type = "string"
	form.Fields = append(form.Fields, field)

	models.NCDB.Create(&form)
	field.FormId = form.ID
	models.NCDB.Create(&field)

	//models.NCDB.Model(&form).Related(&[]models.NC_Field{field})

	//servlet.Create()
	f.Data["json"] = "create form success"
	f.ServeJSON()
}

//(invalid association [])
//[2020-03-25 00:01:29]

// @Title	CreatedServlet
// @Description 创建服务
// @Param	name	string	true		"服务名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /search [Post]
func (f *FormController) Search() {

	//一对多查询
	var form models.NC_Form
	models.NCDB.First(&form)
	models.NCDB.Debug().Model(&form).Related(&form.Fields, "FormId")

	fmt.Println(form)

	f.Data["json"] = "create model success"
	f.ServeJSON()
}
