// @Title  nC_Servlet_SystemMethod
// @Description  系统逻辑源码
package models

var MethodMap map[string]string

//todo 通用-返回成功
var ResponseSuccess = `func (ncc *NCController) responseSuccess(result interface{}) {
		res := make(map[string]interface{})
		res["list"] = result
		res["status"] = "success"
		res["description"] = "操作成功"
		ncc.Data["json"] = res
		ncc.ServeJSON()
	}`

//todo 通用-处理是否解析成功
var HandlerErrOK = `func (ncc *NCController) handlerErrOK(err error) bool {
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
	}`
