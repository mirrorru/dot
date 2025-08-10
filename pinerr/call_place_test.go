package pinerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCallPlace(t *testing.T) {
	t.Parallel()
	cp := getCallPlace(1)
	assert.Contains(t, cp.fileName, "pinerr.Test_getCallPlace")
}

func Test_callPlace_empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		src  callPlace
		want bool
	}{
		{name: "empty", want: true},
		{name: "not empty", src: callPlace{fileName: "-"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.src.empty())
		})
	}
}
