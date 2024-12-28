package category

import (
	"komo/app/product/repo/pg"
	repo "komo/app/product/repo/pg/category"
	"komo/lib/engine"
)

func CategorySlugExists(slug string) *engine.Result[bool] {
	return repo.CategorySlugExists(slug)
}

type FilterCategoriesInput struct {
	Limit    int
	LastSlug string
	State    string
}

func FilterExistingCategories(input FilterCategoriesInput) *engine.Result[[]repo.CategoryData] {
	return repo.ListCategories(repo.ListCategoriesInput{
		State: input.State}, pg.Paging{
		LastId: input.LastSlug,
		Limit:  input.Limit,
	})
}
