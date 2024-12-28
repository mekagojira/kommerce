package service

import (
	"komo/app/product/repo/pg"
	"komo/lib/engine"
	"time"
)

type CreateCategoryInput struct {
	Slug         string
	CategoryName string
}

func CreateCategory(input CreateCategoryInput) *engine.Result[bool] {
	return pg.CreateCategory(pg.CategoryData{
		Slug:         input.Slug,
		CategoryName: input.CategoryName,
		CreatedAt:    time.Now(),
		State:        pg.CATEGORY_STATE_INACTIVE,
	})
}

func CategorySlugExists(slug string) *engine.Result[bool] {
	return pg.CategorySlugExists(slug)
}
