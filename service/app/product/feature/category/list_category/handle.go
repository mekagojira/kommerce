package listcategory

import (
	"komo/app/product/common/constant"
	categoryService "komo/app/product/service/category"
)

func Handle(ctx Ctx) Response {
	res := categoryService.FilterExistingCategories(categoryService.FilterCategoriesInput{
		Limit:    ctx.Req.Limit,
		LastSlug: ctx.Req.Position,
		State:    constant.CATEGORY_STATE_ACTIVE,
	})

	if res.Error != nil {
		return ctx.ServerError()
	}

	var categories []Category
	position := ""

	for _, category := range res.PureData() {
		categories = append(categories, Category{
			Slug:         category.Slug,
			CategoryName: category.CategoryName,
		})
		position = category.Slug
	}

	return ctx.Ok(&Output{
		Categories: categories,
		Position:   position,
	})
}
