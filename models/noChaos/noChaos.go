package noChaos

import (
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
)

var (
	DeployPath string // 部署环境地址
)

func init() {

	currentPath := `\`

	switch beego.BConfig.RunMode {
	case "dev":
		currentPath, _ = os.Getwd()
	case "test":
		dir, _ := os.Executable()
		currentPath = filepath.Dir(dir) + "/noChaos-Server"
	default:
		currentPath, _ = os.Getwd()
	}
	DeployPath = currentPath + "_Data"
}
