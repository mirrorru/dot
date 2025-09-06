package dot

import (
	"fmt"
	"reflect"
)

type Result[T any] struct {
	val T
	err error
}

func MakeResult[T any](val T, err error) Result[T] {
	return Result[T]{val: val, err: err}
}

func (r Result[T]) SaveVal(dest any) Result[T] {
	if r.err != nil {
		return r
	}

	// Используем reflection для проверки и присвоения
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("expected pointer to value or interface, got %v", v.Kind()))
	}

	// Получаем значение, на которое указывает pointer
	v = v.Elem()
	switch v.Kind() {
	case reflect.Interface:
		valueVal := reflect.ValueOf(r.val)
		if !valueVal.Type().Implements(v.Type()) {
			panic(fmt.Errorf("value of type %v does not implement interface %v",
				valueVal.Type(), v.Type()))
		}

		v.Set(valueVal)
	default:
		valueVal := reflect.ValueOf(r.val)
		v.Set(valueVal)
	}

	return r
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) Val() T {
	return r.val
}

func (r Result[T]) OrEmpty() (empty T) {
	if r.err == nil {
		return r.val
	}

	return empty
}

func (r Result[T]) OrElse(anotherVal T) T {
	if r.err == nil {
		return r.val
	}

	return anotherVal
}

func (r Result[T]) Unwarp() (T, error) {
	return r.val, r.err
}

func (r Result[T]) ToOption() Option[T] {
	if r.err != nil {
		return Option[T]{}
	}

	return Option[T]{Val: r.val, Ok: true}
}

func ConvertResult[T1, T2 any](res Result[T1], converter func(src T1) (T2, error)) Result[T2] {
	if res.err != nil {
		return Result[T2]{err: res.err}
	}

	return MakeResult(converter(res.val))
}
