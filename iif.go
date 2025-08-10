package dot

// Iif - inline "if".
func Iif[T any](condition bool, resultOnTrue T, resultOnFalse T) T {
	if condition {
		return resultOnTrue
	}
	return resultOnFalse
}
