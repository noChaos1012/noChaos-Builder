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
	In  T_In27
	Out T_Out27
}

// @router /Panduanceshi20041301_27 [Post]
func (C *Panduanceshi20041301_27) Panduanceshi20041301_27() {
	err := json.Unmarshal(C.Ctx.Input.RequestBody, &C.In)
	if C.handlerErrOK(err) {
		C.LogicBody()
		C.responseSuccess(map[string]interface{}{"output": C.Out})
	}
}
func (C *Panduanceshi20041301_27) LogicBody() {

	if C.In.Param_1 == 1 {
		C.Out.Param_2 = "1"
	}
	if C.In.Param_1 == 2 {
		C.Out.Param_2 = "2"
	}
	if C.In.Param_1 == 3 {
		C.Out.Param_2 = "3"
	}

}
