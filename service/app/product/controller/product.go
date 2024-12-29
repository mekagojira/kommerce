package controller

import (
	createProduct "komo/app/product/feature/product/create_product"
	"komo/lib/engine"
)

func init() {
	// ADMIN
	engine.RegisterEndpoint("/admin/product/create-product", createProduct.Handle, "admin")
}
