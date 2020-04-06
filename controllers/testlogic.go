package controllers

import "encoding/json"

type t_Param_2 struct {
	Param_7 string
	Param_8 string
}

type T_In20 struct {
	Param_1 string
	Param_2 t_Param_2
}

type T_Out20 struct {
	Param_3 string
	Param_4 string
}

type t_Node_1 struct {
	Param_5 string
	Param_6 string
}

type Ceshiluoji20 struct {
	NCController
}

// @router /Ceshiluoji20 [Post]
func (c *Ceshiluoji20) Ceshiluoji20() {
	var In T_In20
	var Out T_Out20
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &In)
	if c.handlerErrOK(err) {
		c.LogicBody(&In, &Out)
		c.responseSuccess(map[string]interface{}{"output": Out})
	}
}

func (c *Ceshiluoji20) LogicBody(In *T_In20, Out *T_Out20) {
	var Node_1 t_Node_1
	Node_1.Param_5 = "aa"
	Node_1.Param_6 = "bb"
	Out.Param_3 = Node_1.Param_5 + In.Param_1
	Out.Param_4 = Node_1.Param_6 + In.Param_2.Param_7 + In.Param_2.Param_8
}
