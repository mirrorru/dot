package dot

import (
	"errors"
	"fmt"
	"reflect"
)

// Result holds operation result - some value and error
type Result[T any] struct {
	val T
	err error
}

// MakeResult maker Result from arguments
func MakeResult[T any](val T, err error) Result[T] {
	return Result[T]{val: val, err: err}
}

// SaveVal writes inner value to reference
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
	valueVal := reflect.ValueOf(r.val)
	if v.Kind() == reflect.Interface && !valueVal.Type().Implements(v.Type()) {
		panic(fmt.Errorf("value of type %v does not implement interface %v",
			valueVal.Type(), v.Type()))
	}

	v.Set(valueVal)

	return r
}

// IsErr reports about not nil inner error
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Err returns inner error
func (r Result[T]) Err() error {
	return r.err
}

// Val returns inner value
func (r Result[T]) Val() T {
	return r.val
}

// OrEmpty return default (empty) value, if result have error.
func (r Result[T]) OrEmpty() (empty T) {
	if r.err == nil {
		return r.val
	}

	return empty
}

// OrElse return argument value, if result have error
func (r Result[T]) OrElse(anotherVal T) T {
	if r.err == nil {
		return r.val
	}

	return anotherVal
}

// Unwarp extracts value and error
func (r Result[T]) Unwarp() (T, error) {
	return r.val, r.err
}

// ToOption - converts Result to Option type
func (r Result[T]) ToOption() Option[T] {
	if r.err != nil {
		return Option[T]{}
	}

	return Option[T]{Val: r.val, Ok: true}
}

func ConvertResult[T1, T2 any](res Result[T1], converter func(src T1) (T2, error)) Result[T2] {
	return FromResult(res, converter)
}

// FromResult - make new `Result[T2]` from  `Result[T1]` by `func(T1)Result[T2]`
func FromResult[T1, T2 any](res Result[T1], converter func(src T1) (T2, error)) Result[T2] {
	if res.err != nil {
		return Result[T2]{err: res.err}
	}

	return MakeResult(converter(res.val))
}

var errWrongCastingType = errors.New("wrong casting type")

// CastResult - casts any to specified type or return error
func CastResult[T any](src Result[any]) Result[T] {
	if src.IsErr() {
		return Result[T]{err: src.Err()}
	}
	val, ok := src.Val().(T)
	if !ok {
		return Result[T]{err: errWrongCastingType}
	}
	return MakeResult(val, nil)
}
