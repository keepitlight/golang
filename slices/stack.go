package slices

import (
	"sync"

	"github.com/keepitlight/golang"
)

// Stack 栈，可以使用 sync.Mutex 或 sync.RWMutex 作为锁（互斥锁或读写锁）
type Stack[T any] struct {
	l     sync.Locker
	items []T
}

// Push 创建栈并将多个元素压入
func Push[T any](items ...T) *Stack[T] {
	return &Stack[T]{items: items}
}

// Lock 使用锁创建栈
func Lock[T any](l sync.Locker, items ...T) *Stack[T] {
	return &Stack[T]{items: items, l: l}
}

func RWLock[T any](items ...T) *Stack[T] {
	return Lock[T](&sync.RWMutex{}, items...)
}

func Mutex[T any](items ...T) *Stack[T] {
	return Lock[T](&sync.Mutex{}, items...)
}

// Push 在栈顶压入元素
func (s *Stack[T]) Push(items ...T) *Stack[T] {
	if s.l != nil {
		s.l.Lock()
		defer s.l.Unlock()
	}
	s.items = append(s.items, items...)
	return s
}

// Pop 弹出栈顶元素
func (s *Stack[T]) Pop() (T, bool) {
	if s.l != nil {
		s.l.Lock()
		defer s.l.Unlock()
	}
	var null T
	if len(s.items) == 0 {
		return null, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek 查看栈顶数据元素
func (s *Stack[T]) Peek() (T, bool) {
	if rw, ok := s.l.(golang.RWLocker); ok && rw != nil {
		rw.RLock()
		defer rw.RUnlock()
	}
	var null T
	if len(s.items) == 0 {
		return null, false
	}
	return s.items[len(s.items)-1], true
}

// Each to iterate the stack elements from top to bottom
//
// 从栈顶向下遍历数据元素，f 返回 false 则终止遍历立即返回
func (s *Stack[T]) Each(f func(elem T, index int) (stop bool)) (stopped bool) {
	if rw, ok := s.l.(golang.RWLocker); ok && rw != nil {
		rw.RLock()
		defer rw.RUnlock()
	}
	if len(s.items) == 0 {
		return false
	}
	r := false
	for i := len(s.items) - 1; i >= 0; i-- {
		if f(s.items[i], i) {
			r = true
			break
		}
	}
	return r
}

// Len 返回栈中元素个数
func (s *Stack[T]) Len() int {
	if rw, ok := s.l.(golang.RWLocker); ok && rw != nil {
		rw.RLock()
		defer rw.RUnlock()
	}
	return len(s.items)
}

// Clear 清空栈
func (s *Stack[T]) Clear() {
	if s.l != nil {
		s.l.Lock()
		defer s.l.Unlock()
	}
	s.items = make([]T, 0)
}
