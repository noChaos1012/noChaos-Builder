package logic

//控制器代码

var controller_Code = `
package controllers

{{.import_Code}}

{{.function_Code}}

`

//// @Title Create
//// @Description create object
//// @Param	body		body 	models.Object	true		"The object content"
//// @Success 200 {string} models.Object.Id
//// @Failure 403 body is empty
//// @router / [post]
//func (o *ObjectController) Post() {
//	var ob models.Object
//	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
//	objectid := models.AddOne(ob)
//	o.Data["json"] = map[string]string{"ObjectId": objectid}
//	o.ServeJSON()
//}

var package_Code = `"{{.PackageAddr}}"`

//const (
//	master = `Names:{{block "list" .}}{{"\n"}}  {{range .}}{{println "-" .}}{{end}}  {{end}}`
//
//	overlay = `{{define "list"}} {{join . ", "}}{{end}} `
//)

//获取载入包

func GetImports() []string {

	route1 := "aaa"

	return []string{route1, "route1"}
}

var (
//定义了一个功能
//funcs = template.FuncMap{"join": strings.Join}

)

////
//masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
//if err != nil {
//log.Fatal(err)
//}
//
//overlayTmpl, err := template.New("overlay").Funcs(funcs).Parse(overlay)
//
////overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
//if err != nil {
//	log.Fatal(err)
//}
//
//if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
//log.Fatal(err)
//}

func amain() {

}

//创建逻辑代码
//func BuildLogic() {
//
//	fmt.Println(GetImports)
//
//	masterTmpl, err := template.New("master").Parse(controller_Code)
//	//masterTmpl, err := template.New("master").Parse()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var doc bytes.Buffer
//
//	if err := masterTmpl.Execute(&doc, GetImports()); err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(doc.String())
//
//	utils.WriteToFile("aa", doc.String())
//
//	//appPath := noChaos.DeployPath + "/" + appName
//
//	//os.MkdirAll(appPath, 0755)
//	//
//	////cmd := exec.Command("/bin/bash", "-c", "cd "+ appPath )
//	////fmt.Println(cmd.Run())
//	//utils.WriteToFile(path.Join(appPath, "main.go"),strings.Replace(main_Code, "{{.ServletName}}", appName, -1))
//	//
//	//os.Mkdir(path.Join(appPath, "conf"), 0755)
//	//utils.WriteToFile(path.Join(appPath, "conf", "app.conf"), strings.Replace(conf_Code, "{{.ServletName}}", appName, -1))
//	//os.Mkdir(path.Join(appPath, "controllers"), 0755)
//	//os.Mkdir(path.Join(appPath, "models"), 0755)
//	//os.Mkdir(path.Join(appPath, "tests"), 0755)
//}
