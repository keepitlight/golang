package slices

import (
	"sync"
	"testing"
)

func TestNewStackWithLocker(t *testing.T) {
	s := Lock[int](&sync.RWMutex{})
	var wg sync.WaitGroup
	l := 1000
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			s.Push(v)
		}(i)
	}
	wg.Wait()
	if len(s.items) != l {
		t.Errorf("stack error, want %d, got %d", l, len(s.items))
	}
}

func BenchmarkNewStack(b *testing.B) {
	s := Push[int]()
	for i := 0; i < b.N; i++ {
		b.Run("StackWithLocker", func(b *testing.B) {
			s.Push(i)
		})
	}
}

func BenchmarkNewStackWithLocker(b *testing.B) {
	s := Lock[int](&sync.RWMutex{})
	for i := 0; i < b.N; i++ {
		b.Run("StackWithLocker", func(b *testing.B) {
			s.Push(i)
		})
	}
}
