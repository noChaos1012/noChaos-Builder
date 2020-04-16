package controllers

import (
	"com.waschild/noChaos-Server/models"
	"encoding/json"
)

type BuildController struct {
	NCController
}

// @Title	Create
// @Description 创建
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildServlet [Post]
func (c *BuildController) BuildServlet() {

	build := models.NC_Build{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &build)

	if c.handlerErrOK(err) {
		models.NCDB.Create(&build)
		build.BuildServlet()
		c.responseSuccess(map[string]interface{}{"model": &build})
	}
}

// @Title	GetMany
// @Description 查询多个
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getMany [Post]
func (c *BuildController) GetMany() { c.GetManyWithModel(&models.NC_Build{}, &[]models.NC_Build{}) }

// @Title	GetDetail
// @Description 查询详情
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDetail [Post]
func (c *BuildController) GetDetail() { c.GetDetailWithModel(&models.NC_Build{}) }

// @Title	GetDemo
// @Description 获取示例
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /getDemo [Post]
func (c *BuildController) GetDemo() {
	model := models.NC_Build{}
	c.responseSuccess(map[string]interface{}{"model": model})
}

// @Title	BuildForm
// @Description 构建表单源码
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildForm [Post]
//func (c *BuildController) BuildForm() {
//	form := models.NC_Form{}
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
//	if c.handlerErrOK(err) {
//		servlet := models.NC_Servlet{}
//		servlet.ID = form.ServletId
//		models.NCDB.Debug().First(&servlet)
//		form.BuildForm(servlet.GetName())
//		c.responseSuccess(map[string]interface{}{"model": form})
//	}
//}

// @Title	BuildModelsGo
// @Description 构建服务models.go源码
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildModelsGo [Post]
//func (c *BuildController) BuildModelsGo() {
//	servlet := models.NC_Servlet{}
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &servlet)
//	if c.handlerErrOK(err) {
//		success, desc := servlet.BuildModels()
//		if success {
//			c.responseSuccess(map[string]interface{}{"model": servlet})
//		} else {
//			res := make(map[string]interface{})
//			res["status"] = "failure"
//			res["description"] = desc
//			c.Data["json"] = res
//			c.ServeJSON()
//		}
//	}
//}

// @Title	BuildRouterGo
// @Description 构建服务router.go源码
// @Param	name	string	true	"表单名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildRouterGo [Post]
//func (c *BuildController) BuildRouterGo() {
//	servlet := models.NC_Servlet{}
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &servlet)
//	if c.handlerErrOK(err) {
//		servlet.BuildRouter(strconv.Itoa(int(time.Now().Unix())),"专门发布")
//		c.responseSuccess(map[string]interface{}{"model": servlet})
//	}
//}

// @Title	BuildForms
// @Description 构建服务所有form源码
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildForms [Post]
func (c *BuildController) BuildForms() {
	servlet := models.NC_Servlet{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &servlet)
	if c.handlerErrOK(err) {
		models.NCDB.Debug().First(&servlet)
		models.NCDB.Debug().Model(&servlet).Related(&servlet.Forms, "ServletId")
		for _, form := range servlet.Forms {
			form.BuildForm(servlet.GetName())
		}
		c.responseSuccess(map[string]interface{}{"model": servlet})
	}
}

// @Title	BuildLogic
// @Description 构建逻辑
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /buildLogic [Post]
func (c *BuildController) BuildLogic() {
	logic := models.NC_Logic{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &logic)
	if c.handlerErrOK(err) {
		models.NCDB.Debug().First(&logic)
		servlet := models.NC_Servlet{}
		servlet.ID = logic.ServletId
		models.NCDB.Debug().First(&servlet)
		logic.BuildLogic(servlet.GetName())
		c.responseSuccess(map[string]interface{}{"model": logic})
	}
}

//// @Title	BuildServlet
//// @Description 构建服务源码
//// @Param	name	string	true	"表单名称"
//// @Success 200 编译成功
//// @Failure 403 body is empty
//// @router /buildServlet [Post]
//func (c *BuildController) BuildServlet() {
//	servlet := models.NC_Servlet{}
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &servlet)
//	if c.handlerErrOK(err) {
//		servlet.BuildServlet()
//		c.responseSuccess(map[string]interface{}{"models": servlet})
//	}
//}

// @Title	TestStruct
// @Description 测试结构
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /testStruct [Post]
func (ncc *BuildController) TestStruct() {

	//fmt.Println()

	//p2 := NC_Property{}
	//
	//err := json.Unmarshal(ncc.Ctx.Input.RequestBody, &p2)
	//
	//if ncc.handlerErrOK(err) {
	//
	//	var structs []string
	//
	//	fmt.Println(GetCode(p2, &structs))
	//
	//	ncc.responseSuccess(map[string]interface{}{"model": p2})
	//}
}

//func GetCode(property NC_Property, structs *[]string) string {
//
//	code := property.Name + " "
//	if property.Multiple {
//		code += "[]"
//	}
//	if property.Type == "custom" {
//		code += getStructName(property.Name)
//
//	} else {
//		code += property.Type
//	}
//
//	if property.Type == "custom" {
//
//		var typeCode = `
//type {{.TypeName}} struct{
//	{{.TypeProperties}}
//}
//`
//		typeCode = strings.Replace(typeCode, "{{.TypeName}}", property.Name, -1)
//		typeCode = strings.Replace(typeCode, "{{.TypeProperties}}", buildStruct(property.Properties, structs), -1)
//
//		*structs = append(*structs, typeCode)
//
//	}
//
//	return code
//}
