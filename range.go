package golang

import (
	"cmp"
)

type RangeType[T any] interface {
	GetLower() T
	GetUpper() T
}

// Range define a range
//
// 范围，区间
type Range[T any] struct {
	Lower T `json:"lower,omitempty"` // lower value
	Upper T `json:"upper,omitempty"` // upper value

	comparer func(a, b T) int
	null     T
}

func (r *Range[T]) GetLower() T {
	return r.Lower
}

func (r *Range[T]) GetUpper() T {
	return r.Upper
}

func Value[T any](v RangeType[T]) (lower, upper T) {
	if v == nil {
		return
	}
	return v.GetLower(), v.GetUpper()
}

// NewRange create a new range, comparer is a function to compare two value a and b,
// if a < b return -1, a == b return 0, a > b return 1
//
// 使用比较函数创建一个范围，函数 comparer 比较两个值 a 和 b，如果 a < b 返回 -1，a == b 返回 0，a > b 返回 1
func NewRange[T any](lower, upper T, comparer func(a, b T) int) *Range[T] {
	c := comparer(lower, upper)
	if c > 0 {
		lower, upper = upper, lower
	}
	return &Range[T]{
		Lower: lower,
		Upper: upper,

		comparer: comparer,
	}
}

// In checks value weather is in the range.
//
// 检查值 value 是否在范围/区间内。
func (r *Range[T]) In(value T) bool {
	return r.comparer(value, r.Lower) >= 0 && r.comparer(value, r.Upper) <= 0
}

// Pick all value in range
//
// 获取所有在范围内/区间内的值
func (r *Range[T]) Pick(values ...T) (found []T) {
	for _, v := range values {
		if r.comparer(v, r.Lower) >= 0 && r.comparer(v, r.Upper) <= 0 {
			found = append(found, v)
		}
	}
	return
}

// Unpick picks all value not in range
//
// 获取所有不在范围内/区间内的值
func (r *Range[T]) Unpick(values ...T) (found []T) {
	for _, v := range values {
		if r.comparer(v, r.Lower) < 0 && r.comparer(v, r.Upper) > 0 {
			found = append(found, v)
		}
	}
	return
}

func (r *Range[T]) Value() (lower, upper T) {
	return r.Lower, r.Upper
}

// Between create a new range of ordered value
//
// 创建一个有序范围的区间
func Between[T cmp.Ordered](lower, upper T) *Range[T] {
	return NewRange(lower, upper, cmp.Compare[T])
}
