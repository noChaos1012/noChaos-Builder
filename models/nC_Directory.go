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

//func (s NC_Servlet) getNamespaceCode(dirId uint) string {
//
//	code := `
//	{{.namespaceCode}},
//	{{.includeCode}}
//	`
//	dirs := []NC_Directory{}
//
//	dir := NC_Directory{}
//	dir.ServletId = s.ID
//	dir.DirId = 0
//	dir.Type = "逻辑"
//	fmt.Println("dir is ",dir)
//
//	NCDB.Debug().Where(&dir).Find(&dirs)
//	fmt.Println("dirs is ",dirs)
//
//	namespaceCodes := []string{}
//
//	//没有下一级返回空
//	has := 1
//	if len(dirs) > 0 {
//		has = 1
//		for _, dir := range dirs {
//			nsCode := `beego.NSNamespace("/{{.dirName}}",
//				{{.namespaceCode}},
//			)`
//			nsCode = strings.Replace(nsCode, "{{.dirName}}", dir.Name, -1)
//			nsCode = strings.Replace(nsCode, "{{.namespaceCode}}", s.getNamespaceCode(dir.ID), -1)
//			namespaceCodes = append(namespaceCodes, nsCode)
//		}
//	} else {
//		has = 0
//	}
//
//	includeCodes := []string{}
//	logics := []NC_Logic{}
//	NCDB.Where(&NC_Logic{DirId: dirId, ServletId: s.ID}).Find(&logics)
//	fmt.Println("logics is ",logics)
//	if len(logics) > 0 {
//		has = 1
//		for _, logic := range logics {
//			includeCodes = append(includeCodes, "beego.NSInclude(&controllers."+logic.GetName()+"{})")
//		}
//	} else {
//		has = 0
//	}
//
//	if has == 0 {
//		return ""
//	} else {
//		code = strings.Replace(code, "{{.namespaceCode}}", strings.Join(namespaceCodes, ","), -1)
//		code = strings.Replace(code, "{{.includeCode}}", strings.Join(includeCodes, ","), -1)
//	}
//
//	fmt.Println("code is ", code)
//
//	return code
//}
