package createproduct

func Handle(ctx Ctx) Response {
	return ctx.Ok(&Output{})
}
