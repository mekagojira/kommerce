package category

import (
	constant "komo/app/product/common/constant"
	repo "komo/app/product/repo/pg/category"
	"komo/lib/engine"
	"time"
)

type CreateCategoryInput struct {
	Slug         string
	CategoryName string
}

func CreateCategory(input CreateCategoryInput) *engine.Result[bool] {
	return repo.CreateCategory(repo.CategoryData{
		Slug:         input.Slug,
		CategoryName: input.CategoryName,
		CreatedAt:    time.Now(),
		State:        constant.CATEGORY_STATE_INACTIVE,
	})
}
