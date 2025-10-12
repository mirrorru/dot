package dot_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mirrorru/dot"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockScanner is a mock implementation of sql.Scanner for testing
type mockScanner int

func (m *mockScanner) Scan(src any) error {
	val, ok := src.(int)
	if !ok {
		return errors.New("value is not a int")
	}
	*m = mockScanner(val)
	return nil
}

func TestParseTypedVar(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name        string
		targetType  reflect.Type
		input       any
		expected    any
		expectedErr bool
	}{
		// String tests
		{
			name:       "string from string",
			targetType: reflect.TypeOf(""),
			input:      "hello",
			expected:   "hello",
		},
		{
			name:       "string from bytes",
			targetType: reflect.TypeOf(""),
			input:      []byte("hello"),
			expected:   "hello",
		},
		{
			name:       "string from int",
			targetType: reflect.TypeOf(""),
			input:      42,
			expected:   "42",
		},

		// Int tests
		{
			name:       "int from int",
			targetType: reflect.TypeOf(int(0)),
			input:      42,
			expected:   42,
		},
		{
			name:       "int from string",
			targetType: reflect.TypeOf(int(0)),
			input:      "42",
			expected:   42,
		},
		{
			name:        "int from invalid string",
			targetType:  reflect.TypeOf(int(0)),
			input:       "invalid",
			expectedErr: true,
		},
		{
			name:       "int8 from int",
			targetType: reflect.TypeOf(int8(0)),
			input:      42,
			expected:   int8(42),
		},
		{
			name:       "int16 from string",
			targetType: reflect.TypeOf(int16(0)),
			input:      "42",
			expected:   int16(42),
		},
		{
			name:       "int32 from string",
			targetType: reflect.TypeOf(int32(0)),
			input:      "42",
			expected:   int32(42),
		},
		{
			name:       "int64 from string",
			targetType: reflect.TypeOf(int64(0)),
			input:      "42",
			expected:   int64(42),
		},

		// Uint tests
		{
			name:       "uint from uint",
			targetType: reflect.TypeOf(uint(0)),
			input:      uint(42),
			expected:   uint(42),
		},
		{
			name:       "uint from string",
			targetType: reflect.TypeOf(uint(0)),
			input:      "42",
			expected:   uint(42),
		},
		{
			name:        "uint from invalid string",
			targetType:  reflect.TypeOf(uint(0)),
			input:       "invalid",
			expectedErr: true,
		},
		{
			name:       "uint8 from string",
			targetType: reflect.TypeOf(uint8(0)),
			input:      "42",
			expected:   uint8(42),
		},
		{
			name:       "uint16 from string",
			targetType: reflect.TypeOf(uint16(0)),
			input:      "42",
			expected:   uint16(42),
		},
		{
			name:       "uint32 from string",
			targetType: reflect.TypeOf(uint32(0)),
			input:      "42",
			expected:   uint32(42),
		},
		{
			name:       "uint64 from string",
			targetType: reflect.TypeOf(uint64(0)),
			input:      "42",
			expected:   uint64(42),
		},

		// Float tests
		{
			name:       "float32 from float32",
			targetType: reflect.TypeOf(float32(0)),
			input:      float32(42.5),
			expected:   float32(42.5),
		},
		{
			name:       "float64 from string",
			targetType: reflect.TypeOf(float64(0)),
			input:      "42.5",
			expected:   42.5,
		},
		{
			name:        "float64 from invalid string",
			targetType:  reflect.TypeOf(float64(0)),
			input:       "invalid",
			expectedErr: true,
		},

		// Bool tests
		{
			name:       "bool from bool",
			targetType: reflect.TypeOf(false),
			input:      true,
			expected:   true,
		},
		{
			name:       "bool from string",
			targetType: reflect.TypeOf(false),
			input:      "true",
			expected:   true,
		},
		{
			name:        "bool from invalid string",
			targetType:  reflect.TypeOf(false),
			input:       "invalid",
			expectedErr: true,
		},

		// Scanner tests
		{
			name:       "Scanner success int64",
			targetType: reflect.TypeOf(mockScanner(1)),
			input:      42,
			expected:   mockScanner(42),
		},
		{
			name:       "Scanner success string",
			targetType: reflect.TypeOf(mockScanner(1)),
			input:      "42",
			expected:   mockScanner(42),
		},
		{
			name:        "Scanner failure",
			targetType:  reflect.TypeOf(mockScanner(1)),
			input:       "test",
			expectedErr: true,
		},
		{
			name:        "invalid Scanner type",
			targetType:  reflect.TypeOf(&struct{}{}),
			input:       "test",
			expectedErr: true,
		},

		// Unsupported type
		{
			name:        "unsupported type",
			targetType:  reflect.TypeOf([]int{}),
			input:       "test",
			expectedErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := dot.ParseTypedVar(tt.targetType, tt.input)

			if tt.expectedErr {
				require.Error(t, err, "expected an error but got none")
				return
			}

			require.NoError(t, err, "unexpected error")
			assert.Equal(t, tt.expected, result, "result mismatch")
		})
	}
}
