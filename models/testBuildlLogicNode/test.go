package main

import "com.waschild/noChaos-Server/models"

func main() {
	logic := models.Logic_Node{}
	logic.Package = "self"
	logic.LogicId = 26
	inAs := models.Assign{}
	inAs.Key = "In.Param_1"
	inAs.Value = "12"
	logic.InputAssigns = append(logic.InputAssigns, inAs)

	outAs := models.Assign{}
	outAs.Key = "Out.Param_2"
	outAs.Value = "{{Out.Param_2}}"
	logic.OutputAssigns = append(logic.OutputAssigns, outAs)
	logic.GetCode("Node_3")
}
