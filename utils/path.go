package utils

import "os"

var (
	DeployPath      string // 部署环境地址
	PackageRootPath string // 部署环境地址
)

func init() {

	//currentPath := `\`
	//
	//switch beego.BConfig.RunMode {
	//case "dev":
	//	currentPath, _ = os.Getwd()
	//case "test":
	//	dir, _ := os.Executable()
	//	currentPath = filepath.Dir(dir) + "/noChaos-Server"
	//default:
	//	currentPath, _ = os.Getwd()
	//}

	currentPath := os.Getenv("GOPATH") + "/src/"
	DeployPath = currentPath + "noChaos-Server_Data"
	PackageRootPath = "noChaos-Server_Data"

}
