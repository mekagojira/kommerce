package listcategory

import "komo/lib/engine"

type Input struct {
	Limit    int    `json:"limit" validate:"min=1,max=100"`
	Position string `json:"lastSlug" validate:"min=0,max=255"`
}

type Category struct {
	Slug         string `json:"slug"`
	CategoryName string `json:"categoryName"`
}

type Output struct {
	Categories []Category `json:"categories"`
	Position   string     `json:"position"`
}

type Ctx = *engine.Ctx[Input, Output]
type Response = *engine.Response[Output]
