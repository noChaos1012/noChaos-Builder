package controllers

import (
	"com.waschild/noChaos-Server/models"
	"encoding/json"
	"fmt"
)

type FormController struct {
	NCController
}

// @Title	Create
// @Description 创建表单
// @Param	name	string	true		"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (f *FormController) Create() {

	form := models.NC_Form{}

	err := json.Unmarshal(f.Ctx.Input.RequestBody, &form)

	if f.handlerErrOK(err) {
		IdField := models.NC_FormField{}
		IdField.Name = "id"
		IdField.Type = "整数"
		IdField.Size = 32
		IdField.Default = "0"
		IdField.AutoIncrement = true
		IdField.IsKey = true

		form.Fields = append(form.Fields, IdField)
		models.NCDB.Create(&form)
		//for _,field := range form.Fields{
		//	field.FormId = form.ID
		//	models.NCDB.Create(&field)
		//}
		f.responseSuccess(map[string]interface{}{"model": &form})
	}
}

// @Title	GetMany
// @Description 查询多个表单
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (f *FormController) GetMany() { f.GetManyWithModel(&models.NC_Form{}, &[]models.NC_Form{}) }

// @Title	GetDetail
// @Description 查询表单详情
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (f *FormController) GetDetail() {
	form := models.NC_Form{}
	err := json.Unmarshal(f.Ctx.Input.RequestBody, &form)

	if f.handlerErrOK(err) {
		models.NCDB.Debug().Where(&form).First(&form)
		models.NCDB.Debug().Model(&form).Related(&form.Fields, "FormId")
		f.responseSuccess(map[string]interface{}{"model": &form})
	}
}

// @Title	Update
// @Description 修改表单
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /update [Post]
func (f *FormController) Update() {

	form := models.NC_Form{}
	err := json.Unmarshal(f.Ctx.Input.RequestBody, &form)

	if f.handlerErrOK(err) {
		filed := models.NC_FormField{}
		filed.FormId = form.ID
		models.NCDB.Where(filed).Delete(models.NC_FormField{})
		models.NCDB.Save(&form)
		f.responseSuccess(map[string]interface{}{"model": &form})
	}

}

// @Title	Delete
// @Description 删除表单
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /delete [Post]
func (f *FormController) Delete() {
	form := models.NC_Form{}
	err := json.Unmarshal(f.Ctx.Input.RequestBody, &form)

	if f.handlerErrOK(err) {
		models.NCDB.Delete(&form)
		filed := models.NC_FormField{}
		filed.FormId = form.ID
		models.NCDB.Where(filed).Delete(models.NC_FormField{})
		f.responseSuccess(map[string]interface{}{"model": &form})
	}
}

// @Title	InitForm
// @Description 初始化表单，初始化表单字段
// @Param	name	string	true		"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /initForm [Post]
func (f *FormController) InitForm() {

	form := models.NC_Form{}
	form.Name = f.GetString("name")
	form.IsStore = true

	col0 := models.NC_FormField{}
	col0.Name = "id"
	col0.Type = "整数"
	col0.Size = 32
	col0.Default = "0"
	col0.AutoIncrement = true
	col0.IsKey = true

	col1 := models.NC_FormField{}
	col1.Name = "长文本"
	col1.Type = "文本"
	col1.Size = 6667652
	col1.NotNull = true
	col1.Default = "默认文本"

	col11 := models.NC_FormField{}
	col11.Name = "长文本2"
	col11.Type = "文本"
	col11.Size = 255
	col11.Default = "默认文本2"

	col2 := models.NC_FormField{}
	col2.Name = "整数32"
	col2.Type = "整数"
	col2.Size = 32
	col2.Default = "0"
	col2.AutoIncrement = true

	col3 := models.NC_FormField{}
	col3.Name = "整数64"
	col3.Type = "整数"
	col3.Size = 64
	col3.Index = true
	col3.Unique = true

	col4 := models.NC_FormField{}
	col4.Name = "小数32"
	col4.Type = "小数"
	col4.Size = 32
	col4.Decimal = 2
	col4.Maximum = 6

	col5 := models.NC_FormField{}
	col5.Name = "小数64"
	col5.Type = "小数"
	col5.Size = 64
	col5.Decimal = 2
	col5.Maximum = 5

	col6 := models.NC_FormField{}
	col6.Name = "布尔"
	col6.Type = "布尔"

	col7 := models.NC_FormField{}
	col7.Name = "时间"
	col7.Type = "时间"

	col8 := models.NC_FormField{}
	col8.Name = "日期"
	col8.Type = "日期"

	col9 := models.NC_FormField{}
	col9.Name = "日期时间"
	col9.Type = "日期时间"

	col10 := models.NC_FormField{}
	col10.Name = "时间戳"
	col10.Type = "时间戳"

	form.Fields = []models.NC_FormField{col0, col1, col2, col3, col4, col5, col6, col7, col8, col9, col10, col11}

	//models.NCDB.Create(&form)
	//for _, field := range form.Fields {
	//	field.FormId = form.ID
	//	models.NCDB.Create(&field)
	//}

	f.Data["json"] = form
	f.ServeJSON()
}

// @Title	CreatedServlet
// @Description 创建表单代码文件
// @Param	name	string	true		"服务名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /build [Post]
func (f *FormController) Build() {

	//一对多查询
	var form models.NC_Form
	id, _ := f.GetInt("id")
	form.ID = uint(id)
	models.NCDB.First(&form)
	models.NCDB.Debug().Model(&form).Related(&form.Fields, "FormId")
	form.GetCode()
	fmt.Println(form)

	f.Data["json"] = "create model success"
	f.ServeJSON()
}

// @Title	GetDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDemo [Post]
func (c *FormController) GetDemo() {
	model := models.NC_Form{}
	c.responseSuccess(map[string]interface{}{"model": model})
}

// @Title	GetFieldDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getFieldDemo [Post]
func (c *FormController) GetFieldDemo() {
	model := models.NC_FormField{}
	c.responseSuccess(map[string]interface{}{"model": model})
}
