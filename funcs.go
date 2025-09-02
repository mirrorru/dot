package dot

import "fmt"

// Iif - inline "if".
func Iif[T any](condition bool, resultOnTrue T, resultOnFalse T) T {
	if condition {
		return resultOnTrue
	}
	return resultOnFalse
}

func MustMake[T any](val T, err error) T {
	if err == nil {
		return val
	}

	file, line := GetCallPlace(2)
	err = fmt.Errorf("unexpected error at %s:%d: %w", file, line, err)
	panic(err)
}

func MustDo(err error) {
	if err == nil {
		return
	}

	file, line := GetCallPlace(2)
	err = fmt.Errorf("unexpected error at %s:%d: %w", file, line, err)
	panic(err)
}
