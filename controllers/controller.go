package controllers

import (
	"com.waschild/noChaos-Server/models"
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
		panic(err)
		res := make(map[string]interface{})
		res["status"] = "failure"
		res["description"] = "接收数据解析失败"
		ncc.Data["json"] = res
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

// 新增记录
func (ncc *NCController) CreateWithModel(model interface{}) {

	err := json.Unmarshal(ncc.Ctx.Input.RequestBody, model)

	if ncc.handlerErrOK(err) {
		models.NCDB.Create(model)
		ncc.responseSuccess(map[string]interface{}{"model": model})
	}
}

// 查询多个
func (ncc *NCController) GetManyWithModel(model, returnModels interface{}) {
	err := json.Unmarshal(ncc.Ctx.Input.RequestBody, model)
	if ncc.handlerErrOK(err) {
		models.NCDB.Where(model).Find(returnModels)
		ncc.responseSuccess(map[string]interface{}{"models": returnModels})
	}
}

// 查询详情
func (ncc *NCController) GetDetailWithModel(model interface{}) {

	err := json.Unmarshal(ncc.Ctx.Input.RequestBody, model)

	if ncc.handlerErrOK(err) {
		models.NCDB.Where(model).First(model)
		ncc.responseSuccess(map[string]interface{}{"model": model})
	}
}

// 修改
func (ncc *NCController) UpdateWithModel(model interface{}) {
	err := json.Unmarshal(ncc.Ctx.Input.RequestBody, model)

	if ncc.handlerErrOK(err) {
		models.NCDB.Debug().Save(model)
		ncc.responseSuccess(map[string]interface{}{"model": model})
	}
}

// 删除
func (ncc *NCController) DeleteWithModel(model interface{}) {

	err := json.Unmarshal(ncc.Ctx.Input.RequestBody, model)

	if ncc.handlerErrOK(err) {
		models.NCDB.Delete(model)
		ncc.responseSuccess(map[string]interface{}{"model": model})
	}
}
