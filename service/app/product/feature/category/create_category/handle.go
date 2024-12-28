package createcategory

import (
	"komo/app/product/common/err"
	categoryService "komo/app/product/service/category"
	"strings"
)

func Handle(ctx Ctx) Response {
	ctx.Req.Slug = strings.ToLower(ctx.Req.Slug)

	{
		res := categoryService.CategorySlugExists(ctx.Req.Slug)
		if res.Error != nil {
			return ctx.ServerError()
		}
		if res.PureData() {
			return ctx.Error(err.CategoryAlreadyExists())
		}
	}

	{
		res := categoryService.CreateCategory(categoryService.CreateCategoryInput{
			Slug:         ctx.Req.Slug,
			CategoryName: ctx.Req.CategoryName,
		})
		if res.Error != nil {
			return ctx.ServerError()
		}
	}

	return ctx.Ok(&Output{
		Slug: ctx.Req.Slug,
	})
}
