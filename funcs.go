package dot

import "fmt"

// Iif - inline "if".
func Iif[T any](condition bool, resultOnTrue T, resultOnFalse T) T {
	if condition {
		return resultOnTrue
	}
	return resultOnFalse
}

// MustMake - checks that the second argument is not an error then return first argument, otherwise it panics
// Usually, the arguments is passed directly as the result of calling another method.
func MustMake[T any](val T, err error) T {
	if err == nil {
		return val
	}

	file, line := GetCallPlace(2)
	err = fmt.Errorf("unexpected error at %s:%d: %w", file, line, err)
	panic(err)
}

// Must - equivalent to MustMake.
func Must[T any](val T, err error) T {
	return MustMake(val, err)
}

// MustDo - checks that the argument is not an error, otherwise it panics.
// Usually, the argument is passed directly as the result of calling another method.
func MustDo(err error) {
	if err == nil {
		return
	}

	file, line := GetCallPlace(2)
	err = fmt.Errorf("unexpected error at %s:%d: %w", file, line, err)
	panic(err)
}

// GetIf - returns the second argument if the condition is true, default value otherwise.
// Designed to replace the "var+if" block.
func GetIf[T any](condition bool, resultOnTrue T) (result T) {
	if condition {
		return resultOnTrue
	}

	return result
}

func FirstOfTwo[T1, T2 any](arg1 T1, _ T2) T1 {
	return arg1
}

func SecondOfTwo[T1, T2 any](_ T1, arg2 T2) T2 {
	return arg2
}
