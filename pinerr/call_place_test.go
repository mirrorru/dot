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
