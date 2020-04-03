package controllers

import "encoding/json"

type t_inputNode struct {
	in_a string
	in_b string
}

type t_outputNode struct {
	out_a string
	out_b string
}

type t_Node1 struct {
	a string
	b string
}

type Jinriwangong struct {
	NCController
}

func (c *Jinriwangong) Jinriwangong() {
	var inputNode t_inputNode
	var outputNode t_outputNode
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &inputNode)
	if c.handlerErrOK(err) {
		var node1 t_Node1
		node1.a = "aa"
		node1.b = "bb"
		outputNode.out_a = node1.a + inputNode.in_a
		outputNode.out_b = node1.b + inputNode.in_b
		c.responseSuccess(map[string]interface{}{"output": outputNode})
	}
}
