package pinerr

import (
	"fmt"
	"sync"
)

type StaticError struct {
	once   sync.Once
	place  callPlace // Calling place, inits once
	format string    // Initial error text to convent in result message with calling place information
	err    error     // Result error
	args   []any
}

func NewStatic(format string, args ...any) *StaticError {
	newArgs := make([]any, len(args), len(args)+2)
	copy(newArgs, args)
	return &StaticError{
		format: format,
		args:   args,
	}
}

func (p *StaticError) Produce() error {
	const getCallPlaceSkip = 5
	p.once.Do(func() {
		p.place = getCallPlace(getCallPlaceSkip)
		p.args = append(p.args, p.place.fileName, p.place.line)
		p.err = fmt.Errorf(p.format+" @%s:%d", p.args...)
	})

	return p.err
}
