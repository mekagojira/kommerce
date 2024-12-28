package adminlistcategory

import (
	categoryService "komo/app/product/service/category"
)

func Handle(ctx Ctx) Response {
	if err := validate(ctx); !err.IsOk() {
		return ctx.BadRequest()
	}

	filter := categoryService.FilterCategoriesInput{
		Limit:    ctx.Req.Limit,
		LastSlug: ctx.Req.Position,
		State:    ctx.Req.State,
	}
	res := categoryService.FilterExistingCategories(filter)

	if res.Error != nil {
		return ctx.ServerError()
	}

	var categories []Category
	position := ""

	for _, category := range res.PureData() {
		categories = append(categories, Category{
			Slug:         category.Slug,
			CategoryName: category.CategoryName,
			State:        category.State,
		})
		position = category.Slug
	}

	return ctx.Ok(&Output{
		Categories: categories,
		Position:   position,
	})
}
