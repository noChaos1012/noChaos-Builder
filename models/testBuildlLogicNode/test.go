package main

import (
	"com.waschild/noChaos-Server/models"
	"fmt"
)

func main() {
	logic := models.Logic_Node{}
	logic.Package = "self"
	logic.LogicId = 27
	inAs := models.Assign{}
	inAs.Key = "In.Param_1"
	inAs.Value = "1"
	logic.InputAssigns = append(logic.InputAssigns, inAs)

	outAs := models.Assign{}
	outAs.Key = "Out.Param_2"
	outAs.Value = "{{C.Out.Param_2}}"
	logic.OutputAssigns = append(logic.OutputAssigns, outAs)
	logic.GetCode("Node_3")

	fmt.Println(logic.GetCode("Node_3"))
}
