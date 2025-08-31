package dot

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSyncSlice(t *testing.T) {
	t.Parallel()

	t.Run("InitSize", func(t *testing.T) {
		t.Parallel()

		s := &SyncSlice[int]{}
		s.InitSize(3, 10)

		assert.Len(t, s.Values(), 3)
		assert.Equal(t, 10, cap(s.Values()))
	})

	t.Run("Append and Values", func(t *testing.T) {
		t.Parallel()

		s := &SyncSlice[string]{}

		// Добавляем элементы
		s.Append("first")
		s.Append("second")
		s.Append("third")

		// Проверяем значения
		values := s.Values()
		assert.Len(t, values, 3)
		assert.Equal(t, []string{"first", "second", "third"}, values)
	})

	t.Run("Get and Set", func(t *testing.T) {
		t.Parallel()

		s := &SyncSlice[int]{}
		s.InitSize(3, 5)

		// Устанавливаем значения
		s.Set(0, 100)
		s.Set(1, 200)
		s.Set(2, 300)

		// Получаем и проверяем значения
		assert.Equal(t, 100, s.Get(0))
		assert.Equal(t, 200, s.Get(1))
		assert.Equal(t, 300, s.Get(2))

		// Меняем значение и проверяем
		s.Set(1, 999)
		assert.Equal(t, 999, s.Get(1))
	})

	t.Run("Concurrent append", func(t *testing.T) {
		t.Parallel()

		s := &SyncSlice[int]{}
		const goroutines = 100
		const iterations = 100

		var wg sync.WaitGroup
		wg.Add(goroutines)

		for i := range goroutines {
			go func(start int) {
				defer wg.Done()
				for j := range iterations {
					s.Append(start + j)
				}
			}(i * iterations)
		}

		wg.Wait()

		// Проверяем, что все элементы добавлены
		assert.Len(t, s.Values(), goroutines*iterations)

		// Проверяем, что нет дубликатов (все значения уникальны)
		values := s.Values()
		valueSet := make(map[int]bool)
		for _, v := range values {
			valueSet[v] = true
		}
		assert.Len(t, valueSet, goroutines*iterations, "all values should be unique")
	})

	t.Run("Concurrent get and set", func(t *testing.T) {
		t.Parallel()

		s := &SyncSlice[int]{}
		s.InitSize(10, 10)

		const goroutines = 50

		var wg sync.WaitGroup
		wg.Add(goroutines * 2)

		// Горутины для записи
		for i := range goroutines {
			go func(index int) {
				defer wg.Done()
				for j := range 100 {
					s.Set(index%10, j)
				}
			}(i)
		}

		// Горутины для чтения
		for i := range goroutines {
			go func(index int) {
				defer wg.Done()
				for range 100 {
					val := s.Get(index % 10)
					require.IsType(t, 0, val)
				}
			}(i)
		}

		wg.Wait()
		// Тест проходит, если не было паники из-за гонки данных
	})

	t.Run("Empty slice operations", func(t *testing.T) {
		t.Parallel()

		s := &SyncSlice[float64]{}

		// Append в пустой слайс
		s.Append(3.14)
		assert.Equal(t, 3.14, s.Get(0))

		// Values пустого слайса
		assert.Empty(t, (&SyncSlice[string]{}).Values())
	})

	t.Run("Type compatibility", func(t *testing.T) {
		t.Parallel()

		// Тестируем с разными типами
		stringSlice := &SyncSlice[string]{}
		stringSlice.Append("test")
		assert.Equal(t, "test", stringSlice.Get(0))

		structSlice := &SyncSlice[struct{ name string }]{}
		testStruct := struct{ name string }{name: "test"}
		structSlice.Append(testStruct)
		assert.Equal(t, testStruct, structSlice.Get(0))
	})
}

func TestSyncSlice_EdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("InitSize on initialized slice", func(t *testing.T) {
		t.Parallel()

		s := &SyncSlice[int]{}
		s.values = []int{1, 2, 3}

		// InitSize не должен перезаписывать уже инициализированный слайс
		s.InitSize(5, 10)
		assert.Equal(t, []int{1, 2, 3}, s.Values())
	})

	t.Run("Zero value slice", func(t *testing.T) {
		t.Parallel()

		var s SyncSlice[int]

		// Методы должны работать на zero value
		s.Append(42)
		assert.Equal(t, 42, s.Get(0))
	})
}
