package build

import (
	"com.waschild/noChaos-Server/models"
)

//初始化服务框架
func InitServlet(name string) {

	//fmt.Println("buildServlet")
	//appPath := utils.DeployPath + "/" + name       //绝对路径
	//rootPath := utils.PackageRootPath + "/" + name //相对包内路径
	//
	//fmt.Println("appPath: " + appPath)
	//
	//os.MkdirAll(appPath, 0755)
	//
	////cmd := exec.Command("/bin/bash", "-c", "cd "+ appPath )
	////fmt.Println(cmd.Run())
	//
	//os.Mkdir(path.Join(appPath, "conf"), 0755)
	//os.Mkdir(path.Join(appPath, "controllers"), 0755)
	//os.Mkdir(path.Join(appPath, "models"), 0755)
	//os.Mkdir(path.Join(appPath, "routers"), 0755)
	//os.Mkdir(path.Join(appPath, "tests"), 0755)
	//
	//utils.WriteToFile(path.Join(appPath, "main.go"), strings.Replace(main_Code, "{{.ServletName}}", rootPath, -1))
	//utils.WriteToFile(path.Join(appPath, "conf", "app.conf"), strings.Replace(conf_Code, "{{.ServletName}}", name, -1))

	//objectLogic := Logic{"object", []Variable{}, []Variable{}, []Node{}}
	//utils.WriteToFile(path.Join(appPath, "controllers", "object.go"), objectLogic.GetCode())
	//utils.WriteToFile(path.Join(appPath, "models", "models.go"), model_Code)
	//utils.WriteToFile(path.Join(appPath, "routers", "router.go"), strings.Replace(router_Code, "{{.ServletName}}", rootPath, -1))

}

//打包服务框架
func PackageServlet(name, mode string) {

	//fmt.Println("buildServlet")
	//appPath := noChaos.DeployPath + "/" + name
	//
	//beePackCode := ""
	//switch mode {
	//case "mac":
	//	beePackCode = "bee pack"
	//case "linux":
	//	beePackCode = "bee pack -be GOOS=linux"
	//
	//default:
	//	beePackCode = ""
	//}
	//
	//fmt.Println(appPath + beePackCode)
	//
	//cmd := exec.Command("/bin/bash", "-c", "cd "+appPath+";"+beePackCode)
	//fmt.Println(cmd.Run())

}

//打包服务
func packageServlet(s models.NC_Servlet) {
	//appPath := noChaos.DeployPath + "/" + s.Name       //绝对路径

	//构建models文件
	//utils.WriteToFile(path.Join(appPath, "models", "models.go"), model_Code)

}
