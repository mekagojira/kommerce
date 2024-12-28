package err

func CategoryNotFound() (string, string) {
	return "CATEGORY_NOT_FOUND", "Category not found"
}

func CategoryAlreadyExists() (string, string) {
	return "CATEGORY_ALREADY_EXISTS", "Category already exists"
}
