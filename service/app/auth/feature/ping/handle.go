package ping

func Handle(ctx Ctx) Response {
	return ctx.Ok(&Output{
		Message: "pong",
	})

	// or error
	// return ctx.Error("Error", "Something went wrong")
}
