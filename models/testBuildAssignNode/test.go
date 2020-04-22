package main

import (
	"com.waschild/noChaos-Server/models"
	"fmt"
)

func main() {
	as := models.Assign{}
	as.Key = "node1.B"
	as.Value = `((长度("{{node1.A}}")+12)*32)-64+((长度("{{node1.A}}")+12)*32)`
	fmt.Println(as.AnalyzeExpression())
	//fmt.Println(ex)
	//as.AssembleExpressions(ex)
	////替换参数
	//rex := regexp.MustCompile(`/(.*?/)`)
	//out := rex.FindAllStringSubmatch(ex, -1) //以空格区分
	//fmt.Println(out)
	//for _, i := range out {
	//	for _,j := range i{
	//		fmt.Println(j)
	//	}
	//}
}
