package dot

type Option[T any] struct {
	Val T
	Ok  bool
}

func ToOption[T any](input T) Option[T] {
	return Option[T]{Val: input, Ok: true}
}

func ToOptionEmpty[T any](input T) Option[T] {
	return Option[T]{}
}

func ToOptionPtr[T any, PT *T](input PT) Option[T] {
	if input == nil {
		return ToOptionEmpty(*input)
	}
	return ToOption(*input)

}
