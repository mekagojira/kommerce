package product

import (
	"komo/app/product/common/constant"
	repo "komo/app/product/repo/pg/product"
	"komo/lib/engine"
	"komo/lib/util"
)

type CreateProductInput struct {
	Slug  string
	Name  string
	Price string
	State string
}

func CreateProduct(data CreateProductInput) *engine.Result[bool] {
	newProduct := repo.ProductData{
		Slug:  data.Slug,
		Name:  data.Name,
		Price: data.Price,
		State: data.State,
	}
	if util.IsEmptyOrWhitespace(newProduct.State) {
		newProduct.State = constant.PRODUCT_STATE_ACTIVE
	}

	res := repo.CreateProduct(newProduct)
	if !res.IsOk() {
		return res
	}

	return res
}
