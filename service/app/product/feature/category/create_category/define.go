package createcategory

import "komo/lib/engine"

type Input struct {
	Slug         string `json:"slug" validate:"min=1,max=255"`
	CategoryName string `json:"categoryName" validate:"min=1,max=255"`
}

type Output struct {
	Slug string `json:"slug"`
}

type Ctx = *engine.Ctx[Input, Output]
type Response = *engine.Response[Output]
