package dot_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mirrorru/dot"

	"github.com/stretchr/testify/assert"
)

func TestIif(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		condition bool
		onTrue    any
		onFalse   any
		expect    any
	}{
		{name: "string true", condition: true, onTrue: "a", onFalse: "b", expect: "a"},
		{name: "string false", condition: false, onTrue: "aa", onFalse: "bb", expect: "bb"},
		{name: "int true", condition: true, onTrue: 1, onFalse: 2, expect: 1},

		{name: "int false", condition: false, onTrue: 11, onFalse: 22, expect: 22},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expect, dot.Iif(tc.condition, tc.onTrue, tc.onFalse))
		})
	}
}

func ExampleIif() {
	for i := range 4 {
		fmt.Println(dot.Iif(i%2 == 0, "Even", "Odd"))
	}

	// Output:
	// Even
	// Odd
	// Even
	// Odd
}

func TestMustMake(t *testing.T) {
	t.Parallel()

	const mustVal = "Abc"
	assert.NotPanics(t, func() {
		x := dot.MustMake(mustVal, nil)
		assert.Equal(t, mustVal, x)
	})

	assert.Panics(t, func() {
		dot.MustMake(mustVal, assert.AnError)
	})
}

func ExampleMustMake() {
	s := dot.MustMake(func() (string, error) {
		return "The created string", nil
	}())
	fmt.Println(s)

	// Output:
	// The created string
}

func TestMust(t *testing.T) {
	t.Parallel()

	const mustVal = "Abc"
	assert.NotPanics(t, func() {
		x := dot.MustMake(mustVal, nil)
		assert.Equal(t, mustVal, x)
	})

	assert.Panics(t, func() {
		dot.Must(mustVal, assert.AnError)
	})
}

func TestMustDo(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		dot.MustDo(nil)
	})

	assert.Panics(t, func() {
		dot.MustDo(assert.AnError)
	})
}

func ExampleMustDo() {
	dot.MustDo(func() error {
		fmt.Println("No error inside call")

		return nil
	}())

	// Output:
	// No error inside call
}

func TestGetIf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		condition bool
		onTrue    string
		expect    string
	}{
		{name: "string true", condition: true, onTrue: "a", expect: "a"},
		{name: "string false", condition: false, onTrue: "aa", expect: ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expect, dot.GetIf(tc.condition, tc.onTrue))
		})
	}
}

func ExampleGetIf() {
	cond := time.Now().IsZero()

	// Long way
	var aLongWay, bLongWay int
	if cond {
		aLongWay = 1
	}
	if !cond {
		bLongWay = 1
	}

	// Short way
	aShortWay := dot.GetIf(cond, 2)
	bShortWay := dot.GetIf(!cond, 2)

	fmt.Println(aLongWay, bLongWay, aShortWay, bShortWay)

	// Output:
	// 0 1 0 2
}
