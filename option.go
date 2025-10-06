package dot

type Option[T any] struct {
	Val T
	Ok  bool
}

func ToOption[T any](input T) Option[T] {
	return Option[T]{Val: input, Ok: true}
}
