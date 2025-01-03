package engine

import "fmt"

type Result[T any] struct {
	Data    *T
	Error   error
	Handled bool
}

func IsError[T any](res *Result[T]) bool {
	if res == nil {
		return true
	}

	return !res.IsOk()
}

func NewResult[T any](defaultData ...T) *Result[T] {
	res := &Result[T]{}

	if len(defaultData) > 0 {
		res.Data = &defaultData[0]
	}

	return res
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

func (r *Result[T]) WithErrorString(err string) *Result[T] {
	return r.WithError(fmt.Errorf(err))
}

func (r *Result[T]) WithError(err error) *Result[T] {
	r.Error = err

	if GetEnv("DEBUG", "0") == "1" {
		Logger.Error(err.Error())
	}

	return r
}
