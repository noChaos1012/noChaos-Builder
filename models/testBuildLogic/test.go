package main

import (
	"com.waschild/noChaos-Server/models"
)

func main() {
	logic := models.NC_Logic{}
	logic.ID = 27
	models.NCDB.Debug().First(&logic)
	servlet := models.NC_Servlet{}
	servlet.ID = logic.ServletId
	models.NCDB.Debug().First(&servlet)
	logic.BuildLogic(servlet.GetName())

	//formNode := models.Form_Node{}
	//formNode.FormId = 8
	//formNode.Type = "Insert"
	//formNode.Conditions = []models.Judge{}
	//as := models.Assign{Key:"长文本",Value:`"长文本的value"`}
	//formNode.Assigns = []models.Assign{as}
	//formNode.OutputAssigns  = []models.Assign{}
	//fmt.Println(formNode.GetCode("Node_5"))
}
