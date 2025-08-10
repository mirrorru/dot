package pinerr

import (
	"fmt"
	"sync"
)

type StaticError struct {
	once       sync.Once
	place      callPlace // Calling place, inits once
	messageSrc string    // Initial error text to convent in result message with calling place information
	err        error     // Result error
}

func NewStatic(staticMsg string) *StaticError {
	return &StaticError{
		messageSrc: staticMsg,
	}
}

func (p *StaticError) Produce() error {
	const getCallPlaceSkip = 5
	if p.err == nil {
		p.once.Do(func() {
			if p.place.empty() {
				p.place = getCallPlace(getCallPlaceSkip)
			}
			p.err = fmt.Errorf("%s @%s:%d", p.messageSrc, p.place.fileName, p.place.line)
		})
	}

	return p.err
}
