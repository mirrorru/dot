package dot

type Option[T any] struct {
	Data  T
	Valid bool
}
