package controller

import (
	createCategory "komo/app/product/feature/category/create_category"
	"komo/lib/engine"
)

func init() {
	engine.RegisterEndpoint("/product/category/create-category", createCategory.Handle)
}
