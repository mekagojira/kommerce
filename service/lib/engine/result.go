package engine

type Result[T any] struct {
	Data    *T
	Error   error
	Handled bool
}

func NewResult[T any]() *Result[T] {
	return &Result[T]{}
}

func (r *Result[T]) IsOk() bool {
	return r.Error == nil
}

func (r *Result[T]) WithData(data *T) *Result[T] {
	r.Data = data
	return r
}

func (r *Result[T]) WithPureData(data T) *Result[T] {
	r.Data = &data
	return r
}

func (r *Result[T]) PureData() T {
	return *r.Data
}

func (r *Result[T]) WithError(err error) *Result[T] {
	r.Error = err

	if GetEnv("DEBUG", "0") == "1" {
		Logger.Error(err.Error())
	}

	return r
}
