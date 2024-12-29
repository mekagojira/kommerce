package product

import (
	repo "komo/app/product/repo/pg/product"
	"komo/lib/engine"
)

type CreateProductInput struct {
	Slug  string
	Name  string
	Price string
}

func CreateProduct(data CreateProductInput) *engine.Result[bool] {
	newProduct := repo.ProductData{
		Slug:  data.Slug,
		Name:  data.Name,
		Price: data.Price,
	}

	res := repo.CreateProduct(newProduct)
	if !res.IsOk() {
		return res
	}

	return res
}
