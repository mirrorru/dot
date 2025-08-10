package pinerr

import (
	"fmt"
	"sync"
)

type WrappingError struct {
	once  sync.Once
	place callPlace // Calling place, inits once
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
	if p.place.empty() {
		p.once.Do(func() {
			p.place = getCallPlace(getCallPlaceSkip)
		})
	}

	return &wrappingError{
		message:    fmt.Sprintf("%s @%s:%d", errToWrap.Error(), p.place.fileName, p.place.line),
		innerError: errToWrap,
	}
}
