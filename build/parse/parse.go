package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

//定义
type Parse_Define struct {
	Name  string
	Class string
}

type User struct {
	Username, Password string
	RegTime            time.Time
}

func ShowTime(t time.Time, format string) string {
	return t.Format(format)
}

//值的封装变化 Funcs
//

func main() {
	const (
		master = `Names:{{block "list" .}}{{"\n"}}  {{range .}}{{println "-" .}}{{end}}  {{end}}`

		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		//定义了一个功能
		funcs = template.FuncMap{"join": strings.Join}

		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)

	////Execute
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	//
	//overlayTmpl, err := template.New("overlay").Funcs(funcs).Parse(overlay)
	//
	////overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(masterTmpl.DefinedTemplates())
	//
	//if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(overlayTmpl.DefinedTemplates())
	////
	//newTmpl, err := masterTmpl.New("wasc").Option()
	//if err := newTmpl.Execute(os.Stdout, guardians); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(overlayTmpl.DefinedTemplates())

}

//
//func main() {
//Ò
//	var bbb int
//	bbb = 4
//	b, f, s = s, f, b
//
//	zhangSan = 3
//	sweaters := Parse_Define{"张三", "int"}
//	tmpl, err := template.New("add").Parse("var {{.Name}} {{.Class}}")
//
//	if err != nil { panic(err) }
//
//	buf := new(bytes.Buffer) //实现了读写方法的可变大小的字节缓冲
//	//err = tmp1.Execute(buf,name) //err = tmp1.Execute(os.Stdout,name) 表示标准输出写入到控制台
//
//	err = tmpl.Execute(os.Stdout, sweaters)
//
//	if err != nil { panic(err) }
//}
