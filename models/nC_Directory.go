package models

import (
	"com.waschild/noChaos-Server/utils"
	"strings"
)

//文件夹结构 用来存储内容
type NC_Directory struct {
	ID        uint   `gorm:"primary_key"`
	ServletId uint   //服务ID
	Name      string //名称
	DirId     uint   //上级文件夹ID
	Type      string //类型为表单或逻辑
}

//TODO NC_Directory-获取命名空间源码
func (d *NC_Directory) GetNamespaceCode() string {
	code := `
	{{.includeCode}}
	{{.namespaceCode}}
	`
	includeCode := ""
	logics := []NC_Logic{}
	logic := NC_Logic{}
	logic.ServletId = d.ServletId
	logic.DirId = d.ID
	if d.ID == 0 {
		NCDB.Where("dir_id = ? AND servlet_id = ?", 0, d.ServletId).Find(&logics)
	} else {
		NCDB.Where(&logic).Find(&logics)
	}

	if len(logics) > 0 {
		for _, logic := range logics {
			includeCode += "beego.NSInclude(&controllers." + logic.GetName() + "{}),"
		}
	}
	namespaceCode := ""
	dirs := []NC_Directory{}
	dir := NC_Directory{}
	dir.ServletId = d.ServletId
	dir.DirId = d.ID
	dir.Type = "逻辑"
	if d.ID == 0 {
		NCDB.Where("dir_id = ? AND servlet_id = ? AND type = ?", 0, d.ServletId, "逻辑").Find(&dirs)
	} else {
		NCDB.Where(&dir).Find(&dirs)
	}
	if len(dirs) > 0 {
		for _, dir := range dirs {
			nsCode := `beego.NSNamespace("/{{.dirName}}",
				{{.namespaceCode}}`
			nsCode = strings.Replace(nsCode, "{{.dirName}}", utils.GetPinYin(dir.Name), -1)
			nsCode = strings.Replace(nsCode, "{{.namespaceCode}}", dir.GetNamespaceCode(), -1)
			namespaceCode += nsCode
		}
		namespaceCode += "),"
	}
	code = strings.Replace(code, "{{.namespaceCode}}", namespaceCode, -1)
	code = strings.Replace(code, "{{.includeCode}}", includeCode, -1)
	return code
}
