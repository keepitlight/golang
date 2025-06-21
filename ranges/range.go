package ranges

import (
	"cmp"

	"github.com/keepitlight/golang"
)

// In checks value v weather is in the range r.
//
// 检查值 v 是否在范围 r 内。
func In[T any](r golang.Range[T], v T) bool {
	c := r.Comparer()
	if c == nil {
		return false
	}
	l, u := r.Bounds()
	return c(v, l) >= 0 && c(v, u) <= 0
}

// Pick all value in range
//
// 挑选在范围内/区间内的值
func Pick[T any](r golang.Range[T], values ...T) (found []T) {
	c := r.Comparer()
	if c == nil {
		return
	}
	l, u := r.Bounds()
	for _, v := range values {
		if c(v, l) >= 0 && c(v, u) <= 0 {
			found = append(found, v)
		}
	}
	return
}

// Unpick picks all value not in range
//
// 挑选不在范围内/区间内的值
func Unpick[T any](r golang.Range[T], values ...T) (found []T) {
	c := r.Comparer()
	if c == nil {
		return
	}
	l, u := r.Bounds()
	for _, v := range values {
		if c(v, l) < 0 || c(v, u) > 0 {
			found = append(found, v)
		}
	}
	return
}

// Any return true if any value in range
//
// 检查是否有存在在范围内/区间内的值
func Any[T any](r golang.Range[T], values ...T) bool {
	c := r.Comparer()
	if c == nil {
		return false
	}
	l, u := r.Bounds()
	for _, v := range values {
		if c(v, l) >= 0 && c(v, u) <= 0 {
			return true
		}
	}
	return false
}

// All return true if all value in range
//
// 检查是否全部值均存在在范围内/区间内
func All[T any](r golang.Range[T], values ...T) bool {
	c := r.Comparer()
	if c == nil {
		return false
	}
	l, u := r.Bounds()
	for _, v := range values {
		if c(v, l) < 0 || c(v, u) > 0 {
			return false
		}
	}
	return true
}

// Intersect returns the intersection of two ranges, a and b, or nil if one of any is `nil` or non-overlapping.
//
// 返回 a 和 b 的交集，如果任一一个是 `nil` 或者不相交，则返回 nil。
func Intersect[T any](a, b golang.Range[T]) (golang.Range[T], bool) {
	if a == nil || b == nil {
		return nil, false
	}
	c := a.Comparer()
	al, au := a.Bounds()
	bl, bu := b.Bounds()
	// 断开的
	// 1: [al, au] [bl, bu]  => nil
	// 2: [bl, bu] [al, au]  => nil
	// 相交的
	// 1: [al, [bl, au], bu] => [bl: au]
	// 2: [al, [bl, bu], au] => [bl: bu]
	// 3: [bl, [al, au], bu] => [al: au]
	// 4: [bl, [al, bu], au] => [al: bu]
	x := c(al, bl)
	if x == 0 {
		// 起点相同
		if c(au, bu) < 0 {
			bu = au // 取终点最小的
		}
		return &rangeWrapper[T]{
			Lower: al,
			Upper: bu,

			comparer: c,
		}, true
	}
	if x < 0 {
		// a 的起点在 b 的左侧
		if c(au, bl) < 0 {
			// 不相连
			return nil, false
		}
		if c(au, bu) < 0 {
			bu = au
		}
		return &rangeWrapper[T]{
			Lower: bl,
			Upper: bu,

			comparer: c,
		}, true
	}
	if c(bu, al) < 0 {
		// 不相连
		return nil, false
	}
	if c(bu, au) < 0 {
		au = bu
	}
	return &rangeWrapper[T]{
		Lower: al,
		Upper: au,

		comparer: c,
	}, true
}

// Merge combines two ranges, a and b, returning nil if they're both `nil` or non-overlapping,
// and the non-nil range if one of them is nil.
//
// 合并两个范围 a 和 b，如果它们都是 `nil` 或者不相交，则返回 nil，否则返回非空的。
func Merge[T any](a, b golang.Range[T]) (golang.Range[T], bool) {
	if a == nil {
		return b, b != nil
	}
	if b == nil {
		return a, true
	}
	c := a.Comparer()
	al, au := a.Bounds()
	bl, bu := b.Bounds()
	// 断开的
	// 1: [al, au] [bl, bu]  => nil
	// 2: [bl, bu] [al, au]  => nil
	// 相交的
	// 1: [al, [bl, au], bu] => [al: bu]
	// 2: [al, [bl, bu], au] => [al: au]
	// 3: [bl, [al, au], bu] => [bl: bu]
	// 4: [bl, [al, bu], au] => [bl: au]
	x := c(al, bl)
	if x == 0 {
		// 起点相同
		if c(au, bu) > 0 {
			bu = au
		}
		return &rangeWrapper[T]{
			Lower: al,
			Upper: bu,

			comparer: c,
		}, true
	}
	if x < 0 {
		// a 的起点在 b 的左侧
		if c(au, bl) < 0 {
			// 不相连
			return nil, false
		}
		if c(au, bu) > 0 {
			bu = au
		}
		return &rangeWrapper[T]{
			Lower: al,
			Upper: bu,

			comparer: c,
		}, true
	}
	if c(bu, al) < 0 {
		// 不相连
		return nil, false
	}
	if c(bu, au) > 0 {
		au = bu
	}
	return &rangeWrapper[T]{
		Lower: bl,
		Upper: au,

		comparer: c,
	}, true
}

// Bounds get bounds of range
//
// 获取范围边界值
func Bounds[R golang.Range[T], T any](r R) (lower, upper T) {
	return r.Bounds()
}

type rangeWrapper[T any] struct {
	Lower T `json:"lower,omitempty"` // lower value
	Upper T `json:"upper,omitempty"` // upper value

	comparer func(a, b T) int
}

func (r *rangeWrapper[T]) Bounds() (lower, upper T) {
	return r.Lower, r.Upper
}

func (r *rangeWrapper[T]) Comparer() func(a, b T) int {
	return r.comparer
}

// New create a new range, comparer is a function to compare two value a and b,
// if a < b return -1, a == b return 0, a > b return 1, nil if comparer is nil.
//
// 使用比较函数创建一个范围，函数 comparer 比较两个值 a 和 b，如果 a < b 返回 -1，a == b 返回 0，a > b 返回 1。
// 如果 compare 为 nil，则返回 nil。
func New[T any](lower, upper T, comparer func(a, b T) int) golang.Range[T] {
	if comparer == nil {
		return nil
	}
	c := comparer(lower, upper)
	if c > 0 {
		lower, upper = upper, lower
	}
	return &rangeWrapper[T]{
		Lower: lower,
		Upper: upper,

		comparer: comparer,
	}
}

// Between create a new range of ordered value
//
// 创建一个有序值的区间
func Between[T cmp.Ordered](lower, upper T) golang.Range[T] {
	return New(lower, upper, cmp.Compare[T])
}
