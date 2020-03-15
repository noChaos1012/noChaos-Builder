package build

import "fmt"

//创建逻辑代码
func BuildLogic() {

	fmt.Println("buildLogic")

	//appPath := noChaos.DeployPath + "/" + appName

	//os.MkdirAll(appPath, 0755)
	//
	////cmd := exec.Command("/bin/bash", "-c", "cd "+ appPath )
	////fmt.Println(cmd.Run())
	//utils.WriteToFile(path.Join(appPath, "main.go"),strings.Replace(main_Code, "{{.Appname}}", appName, -1))
	//
	//os.Mkdir(path.Join(appPath, "conf"), 0755)
	//utils.WriteToFile(path.Join(appPath, "conf", "app.conf"), strings.Replace(conf_Code, "{{.Appname}}", appName, -1))
	//os.Mkdir(path.Join(appPath, "controllers"), 0755)
	//os.Mkdir(path.Join(appPath, "models"), 0755)
	//os.Mkdir(path.Join(appPath, "tests"), 0755)
}
