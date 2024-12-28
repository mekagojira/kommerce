package ping

import "komo/lib/engine"

type Input struct {
}
type Output struct {
	Message string `json:"message"`
}

type Ctx = *engine.Ctx[Input, Output]
type Response = *engine.Response[Output]
