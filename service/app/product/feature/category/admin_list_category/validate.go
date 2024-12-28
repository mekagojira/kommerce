package adminlistcategory

import (
	"komo/app/product/common/constant"
	"komo/lib/engine"
)

func validate(ctx Ctx) *engine.Result[bool] {
	res := new(engine.Result[bool])

	input := ctx.Req

	if input.State != "" {
		if input.State != constant.CATEGORY_STATE_ACTIVE &&
			input.State != constant.CATEGORY_STATE_INACTIVE {
			return res.WithErrorString("Invalid state")
		}
	}

	return res
}
