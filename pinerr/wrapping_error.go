package pinerr

import (
	"fmt"
	"sync"

	"github.com/mirrorru/dot"
)

type WrappingError struct {
	once  sync.Once
	place struct {
		fileName string
		line     int
	}
}

type wrappingError struct {
	message    string
	innerError error
}

func (e *wrappingError) Error() string {
	return e.message
}

func (e *wrappingError) Unwrap() error {
	return e.innerError
}

func (p *WrappingError) Produce(errToWrap error) error {
	const getCallPlaceSkip = 5
	p.once.Do(func() {
		p.place.fileName, p.place.line = dot.GetCallPlace(getCallPlaceSkip)
	})

	return &wrappingError{
		message:    fmt.Sprintf("%s @%s:%d", errToWrap.Error(), p.place.fileName, p.place.line),
		innerError: errToWrap,
	}
}
