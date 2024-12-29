package createproduct

import "komo/lib/engine"

type Input struct {
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	State string `json:"state"`
}

type Output struct {
	Id string `json:"id"`
}

type Ctx = *engine.Ctx[Input, Output]
type Response = *engine.Response[Output]
