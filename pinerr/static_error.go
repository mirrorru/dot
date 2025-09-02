package pinerr

import (
	"fmt"
	"sync"

	"github.com/mirrorru/dot"
)

type StaticError struct {
	once   sync.Once
	format string // Initial error text to convent in result message with calling place information
	err    error  // Result error
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
		fileName, line := dot.GetCallPlace(getCallPlaceSkip)
		p.args = append(p.args, fileName, line)
		p.err = fmt.Errorf(p.format+" @%s:%d", p.args...)
	})

	return p.err
}
