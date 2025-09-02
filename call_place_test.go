package dot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCallPlace(t *testing.T) {
	t.Parallel()
	file, line := GetCallPlace(1)
	assert.Contains(t, file, "dot.Test_getCallPlace")
	assert.Equal(t, 11, line)
}
