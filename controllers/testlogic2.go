package controllers

import (
	"encoding/json"
)

type T_In27 struct {
	Param_1 int
}

type T_Out27 struct {
	Param_2 string
}
type Panduanceshi20041301_27 struct {
	NCController
}

// @router /Panduanceshi20041301_27 [Post]
func (c *Panduanceshi20041301_27) Panduanceshi20041301_27() {
	var In T_In27
	var Out T_Out27
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &In)
	if c.handlerErrOK(err) {
		c.LogicBody(&In, &Out)
		c.responseSuccess(map[string]interface{}{"output": Out})
	}
}
func (c *Panduanceshi20041301_27) LogicBody(In *T_In27, Out *T_Out27) {
	LogicNode_3 := Panduanceshi20041301_26{}
	InNode_3 := T_In26{}
	InNode_3.Param_1 = 12
	OutNode_3 := T_Out26{}
	LogicNode_3.LogicBody(&InNode_3, &OutNode_3)
	Out.Param_2 = OutNode_3.Param_2
}
