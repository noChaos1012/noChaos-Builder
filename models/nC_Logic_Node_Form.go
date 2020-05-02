// @Title  nC_Logic_Node_Form
// @Description  表单节点的相关处理
package models

import (
	"com.waschild/noChaos-Server/utils"
	"strings"
)

type Form_Node struct {
	FormId        uint     //表单编号
	Type          string   //处理类型
	Conditions    []Judge  //判断条件
	Assigns       []Assign //赋值配置
	OutputAssigns []Assign //输出赋值
}

//TODO Form_Node-获取源码
func (f_node *Form_Node) GetCode(NodeMark string) string {

	form := NC_Form{}
	form.ID = f_node.FormId
	NCDB.First(&form)

	formName := form.GetName() + NodeMark
	code := formName + strings.Replace(" := models.{{form}}{}", "{{form}}", form.GetName(), 1) + "\n"

	assigns := []string{}
	for _, as := range f_node.Assigns {
		as.Key = formName + "." + utils.GetPinYin(as.Key)
		assigns = append(assigns, as.AnalyzeExpression())
	}
	code = code + strings.Join(assigns, "\n")

	//for _, as := range f_node.OutputAssigns {
	//	as.Value = utils.GetPinYin(form.Name + NodeMark + as.Key)
	//	assigns = append(assigns, as.AnalyzeExpression())
	//}

	switch f_node.Type {
	case "Insert":
		code += "\n" + strings.Replace("models.NCDB.Create(&{{formName}})", "{{formName}}", formName, 1)

	case "Update":

	case "Delete":

	case "Search":

	}
	return code
}

////TODO Form_Node-新增源码
//func (f_node *Form_Node) InsertCode(NodeMark string) string {
//
//	return code
//}
