package controller

import (
	adminListCategory "komo/app/product/feature/category/admin_list_category"
	createCategory "komo/app/product/feature/category/create_category"
	listCategory "komo/app/product/feature/category/list_category"
	"komo/lib/engine"
)

func init() {
	// PUBLIC
	engine.RegisterEndpoint("/public/product/category/list-category", listCategory.Handle)

	// ADMIN
	engine.RegisterEndpoint("/admin/product/category/list-category", adminListCategory.Handle, "admin")

	engine.RegisterEndpoint("/admin/product/category/create-category", createCategory.Handle, "admin")
}
