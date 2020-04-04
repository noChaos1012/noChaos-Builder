package controllers

import "encoding/json"

type T_inputNode14 struct {
	In_a string
	In_b string
}

type T_outputNode14 struct {
	Out_a string
	Out_b string
}

type t_Node1 struct {
	A string
	B string
}

type Ceshiluoji14 struct {
	NCController
}

func (c *Ceshiluoji14) Ceshiluoji14() {
	var inputNode T_inputNode14
	var outputNode T_outputNode14
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &inputNode)
	if c.handlerErrOK(err) {
		c.LogicBody(&inputNode, &outputNode)
		c.responseSuccess(map[string]interface{}{"output": outputNode})
	}
}

func (c *Ceshiluoji14) LogicBody(inputNode *T_inputNode14, outputNode *T_outputNode14) {
	var node1 t_Node1
	node1.A = "aa"
	node1.B = "bb"
	outputNode.Out_a = node1.A + inputNode.In_a
	outputNode.Out_b = node1.B + inputNode.In_b
}
