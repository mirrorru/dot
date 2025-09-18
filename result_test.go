package dot

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeResult(t *testing.T) {
	t.Parallel()

	t.Run("successful result", func(t *testing.T) {
		t.Parallel()
		val := "test"
		res := MakeResult(val, nil)

		assert.Equal(t, val, res.val)
		assert.NoError(t, res.err)
	})

	t.Run("error result", func(t *testing.T) {
		t.Parallel()
		err := errors.New("test error")
		res := MakeResult("", err)

		assert.Empty(t, res.val)
		assert.ErrorIs(t, err, res.err)
	})
}

func TestResult_IsErr(t *testing.T) {
	t.Parallel()

	t.Run("with error", func(t *testing.T) {
		t.Parallel()
		err := errors.New("test error")
		res := Result[string]{err: err}

		assert.True(t, res.IsErr())
	})

	t.Run("without error", func(t *testing.T) {
		t.Parallel()
		res := Result[string]{val: "test", err: nil}

		assert.False(t, res.IsErr())
	})
}

func TestResult_Err(t *testing.T) {
	t.Parallel()

	t.Run("returns error", func(t *testing.T) {
		t.Parallel()
		err := errors.New("test error")
		res := Result[string]{err: err}

		assert.Equal(t, err, res.Err())
	})

	t.Run("returns nil when no error", func(t *testing.T) {
		t.Parallel()
		res := Result[string]{val: "test", err: nil}

		assert.NoError(t, res.Err())
	})
}

func TestResult_Val(t *testing.T) {
	t.Parallel()

	t.Run("returns value", func(t *testing.T) {
		t.Parallel()
		val := "test value"
		res := Result[string]{val: val, err: nil}

		assert.Equal(t, val, res.Val())
	})

	t.Run("returns zero value when error exists", func(t *testing.T) {
		t.Parallel()
		err := errors.New("test error")
		res := Result[string]{val: "original", err: err}

		// Val() should return the stored value even with error
		assert.Equal(t, "original", res.Val())
	})
}

func TestResult_OrEmpty(t *testing.T) {
	t.Parallel()

	t.Run("returns empty value", func(t *testing.T) {
		t.Parallel()
		res := Result[string]{val: "test", err: nil}
		empty := res.OrEmpty()

		assert.Equal(t, "test", empty) // empty string for string type
	})

	t.Run("returns empty value even with error", func(t *testing.T) {
		t.Parallel()
		err := errors.New("test error")
		res := Result[string]{val: "test", err: err}
		empty := res.OrEmpty()

		assert.Empty(t, empty)
	})
}

func TestResult_OrElse(t *testing.T) {
	t.Parallel()

	t.Run("returns original value when no error", func(t *testing.T) {
		t.Parallel()
		val := "original"
		fallback := "fallback"
		res := Result[string]{val: val, err: nil}

		result := res.OrElse(fallback)
		assert.Equal(t, val, result)
	})

	t.Run("returns fallback value when error exists", func(t *testing.T) {
		t.Parallel()
		err := errors.New("test error")
		fallback := "fallback"
		res := Result[string]{val: "original", err: err}

		result := res.OrElse(fallback)
		assert.Equal(t, fallback, result)
	})
}

func TestResult_Unwarp(t *testing.T) {
	t.Parallel()

	t.Run("returns value and nil error", func(t *testing.T) {
		t.Parallel()
		val := "test value"
		res := Result[string]{val: val, err: nil}

		resultVal, resultErr := res.Unwarp()
		assert.Equal(t, val, resultVal)
		assert.NoError(t, resultErr)
	})

	t.Run("returns value and error", func(t *testing.T) {
		t.Parallel()
		val := "test value"
		err := errors.New("test error")
		res := Result[string]{val: val, err: err}

		resultVal, resultErr := res.Unwarp()
		assert.Equal(t, val, resultVal)
		assert.Equal(t, err, resultErr)
	})
}

func TestResult_ToOption(t *testing.T) {
	t.Parallel()

	t.Run("converts successful result to Option with Ok=true", func(t *testing.T) {
		t.Parallel()
		val := "test value"
		res := Result[string]{val: val, err: nil}

		option := res.ToOption()
		assert.True(t, option.Ok)
		assert.Equal(t, val, option.Val)
	})

	t.Run("converts error result to Option with Ok=false", func(t *testing.T) {
		t.Parallel()
		err := errors.New("test error")
		res := Result[string]{val: "test", err: err}

		option := res.ToOption()
		assert.False(t, option.Ok)
		assert.Empty(t, option.Val) // zero value
	})
}

func TestConvertResult(t *testing.T) {
	t.Parallel()

	t.Run("propagates error from source result", func(t *testing.T) {
		t.Parallel()
		srcErr := errors.New("source error")
		srcRes := Result[int]{val: 42, err: srcErr}

		converter := func(_ int) (string, error) {
			return "converted", nil
		}

		result := ConvertResult(srcRes, converter)
		assert.True(t, result.IsErr())
		assert.Equal(t, srcErr, result.Err())
	})

	t.Run("converts successful result with successful converter", func(t *testing.T) {
		t.Parallel()
		srcRes := Result[int]{val: 42, err: nil}
		converter := func(_ int) (string, error) {
			return "converted", nil
		}

		result := ConvertResult(srcRes, converter)
		assert.False(t, result.IsErr())
		assert.Equal(t, "converted", result.Val())
	})

	t.Run("returns error from converter", func(t *testing.T) {
		t.Parallel()
		srcRes := Result[int]{val: 42, err: nil}
		converterErr := errors.New("converter error")
		converter := func(_ int) (string, error) {
			return "", converterErr
		}

		result := ConvertResult(srcRes, converter)
		assert.True(t, result.IsErr())
		assert.Equal(t, converterErr, result.Err())
	})
}

func TestOption_Struct(t *testing.T) {
	t.Parallel()

	t.Run("Option with Ok=true", func(t *testing.T) {
		t.Parallel()
		val := "test"
		option := Option[string]{Val: val, Ok: true}

		assert.True(t, option.Ok)
		assert.Equal(t, val, option.Val)
	})

	t.Run("Option with Ok=false", func(t *testing.T) {
		t.Parallel()
		option := Option[string]{}

		assert.False(t, option.Ok)
		assert.Empty(t, option.Val)
	})
}

// Дополнительные тесты для проверки различных типов данных
func TestResult_WithDifferentTypes(t *testing.T) {
	t.Parallel()

	t.Run("int type", func(t *testing.T) {
		t.Parallel()
		res := MakeResult(42, nil)

		assert.Equal(t, 42, res.Val())
		assert.False(t, res.IsErr())
	})

	t.Run("struct type", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		person := Person{Name: "John", Age: 30}
		res := MakeResult(person, nil)

		assert.Equal(t, person, res.Val())
		assert.False(t, res.IsErr())
	})

	t.Run("pointer type", func(t *testing.T) {
		t.Parallel()
		value := "test"
		res := MakeResult(&value, nil)

		assert.Equal(t, &value, res.Val())
		assert.Equal(t, "test", *res.Val())
		assert.False(t, res.IsErr())
	})
}

// Тест для проверки конкурентного доступа (демонстрация безопасности)
func TestResult_ConcurrentAccess(t *testing.T) {
	t.Parallel()

	res := MakeResult("safe", nil)

	t.Run("multiple goroutines reading", func(t *testing.T) {
		t.Parallel()

		// Запускаем несколько горутин для чтения
		done := make(chan bool, 10)
		for range 10 {
			go func() {
				assert.Equal(t, "safe", res.Val())
				assert.False(t, res.IsErr())
				assert.NoError(t, res.Err())
				done <- true
			}()
		}

		// Ждем завершения всех горутин
		for range 10 {
			<-done
		}
	})
}

func TestResult_SaveVal(t *testing.T) {
	t.Parallel()
	const testVal = int(42)

	t.Run("int type", func(t *testing.T) {
		t.Parallel()
		res := MakeResult(testVal, nil)
		var dest int
		require.NotPanics(t, func() {
			res.SaveVal(&dest)
		})
		assert.Equal(t, testVal, dest)
	})

	t.Run("struct type", func(t *testing.T) {
		t.Parallel()
		type MyStruct struct {
			val int
		}
		reference := MyStruct{val: testVal}
		res := MakeResult(reference, nil)
		var dest MyStruct
		require.NotPanics(t, func() {
			res.SaveVal(&dest)
		})
		assert.Equal(t, reference, dest)
	})

	t.Run("time stringer", func(t *testing.T) {
		t.Parallel()
		reference := time.Now()
		res := MakeResult(reference, nil)
		var dest fmt.Stringer
		require.NotPanics(t, func() {
			res.SaveVal(&dest)
		})
		assert.Equal(t, reference, dest)
	})

	t.Run("struct not stringer type", func(t *testing.T) {
		t.Parallel()
		type MyStruct struct {
			val int
		}
		reference := MyStruct{val: testVal}
		res := MakeResult(reference, nil)
		var dest fmt.Stringer
		assert.Panics(t, func() {
			res.SaveVal(&dest)
		})
	})
}

func TestCastResult(t *testing.T) {
	t.Parallel()

	t.Run("successful result", func(t *testing.T) {
		t.Parallel()
		val := "test"
		res := CastResult[string](MakeResult[any](val, nil))

		assert.Equal(t, val, res.val)
		assert.NoError(t, res.err)
	})

	t.Run("cast error", func(t *testing.T) {
		t.Parallel()
		val := "test"
		res := CastResult[int](MakeResult[any](val, nil))

		assert.Error(t, res.err)
	})
}

func TestFromResult(t *testing.T) {
	t.Parallel()

	t.Run("result without error", func(t *testing.T) {
		t.Parallel()
		val := 40
		res1 := MakeResult(val, nil)

		res2 := FromResult(res1, func(src int) (int, error) {
			return src + 2, nil
		})
		require.NoError(t, res2.err)
		assert.Equal(t, 42, res2.val)
	})

	t.Run("result without error", func(t *testing.T) {
		t.Parallel()
		val := 40
		res1 := MakeResult(val, assert.AnError)

		res2 := FromResult(res1, func(src int) (int, error) {
			return src + 2, nil
		})
		require.ErrorIs(t, res2.Err(), assert.AnError)
	})

	t.Run("processing error", func(t *testing.T) {
		t.Parallel()
		val := 40
		res1 := MakeResult(val, nil)

		res2 := FromResult(res1, func(_ int) (int, error) {
			return 0, assert.AnError
		})
		require.ErrorIs(t, res2.Err(), assert.AnError)
	})
}
