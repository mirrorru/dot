package dot

import "fmt"

func Must[T any](val T, err error) T {
	if err != nil {
		return val
	}

	file, line := GetCallPlace(2)
	err = fmt.Errorf("unexpected error at %s:%d: %w", file, line, err)
	panic(err)
}
