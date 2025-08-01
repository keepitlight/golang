package golang

import (
	"strings"
)

// Unique returns a slice of elements that contains no duplicate elements, keeping the original order
//
// 返回不重复的元素切片，保持非重复切片元素的相对顺序，索引大的重复元素被扔掉。
// 注意，不同于 pie.Unique 方法与 slices.Compact 方法
func Unique[E comparable](ss []E) (re []E) {
	if len(ss) < 2 {
		return ss
	}
	for _, s := range ss {
		f := false
		for _, v := range re {
			if v == s {
				f = true
				break
			}
		}
		if !f {
			re = append(re, s)
		}
	}
	return
}

// UniqueFunc returns a slice of elements that contains no duplicate elements, keeping the original order
//
// 根据比较方法的结果返回不重复的元素切片，不改变切片顺序，检索到重复的元素时，索引大的重复元素被扔掉
func UniqueFunc[E any](ss []E, equal func(a, b E) bool) (result []E) {
	l := len(ss)
	if l < 2 {
		return ss
	}
	for i := 0; i < l; i++ {
		f := false
		for _, v := range result {
			if equal(v, ss[i]) {
				f = true
				break
			}
		}
		if !f {
			result = append(result, ss[i])
		}
	}
	return
}

// UniqueString returns a slice of strings that contains no duplicate strings, interpreted as UTF-8 strings.
//
// 返回不重复的字符串（大小写敏感），不改变切片顺序，重复元素被扔掉
func UniqueString(ss ...string) []string {
	return Unique(ss)
}

// FoldString returns a slice of strings that contains no duplicate strings, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general form of case-insensitivity.
//
// 返回不重复的字符串（大小写不敏感），不改变切片顺序，重复元素被扔掉
func FoldString(ss ...string) []string {
	return UniqueFunc(ss, strings.EqualFold)
}

// Slices to create a slice of argument length, and fill it with the given initial element, which may be empty.
//
// 创建一个指定长度的切片，并使用给定的元素 initial 按序对结果切片进行填充，
// 如果 initial 为空，则创建一个指定长度的切片，不填充任何元素，
// 如果 initial 的长度大于 length，则截取 initial 的前 length 个元素填充到结果切片中，
// 如果 initial 的长度小于 length，则重复 initial 的元素填充到结果切片中。
func Slices[T any](length int, initial ...T) (result []T) {
	if len(initial) == 0 {
		return make([]T, length)
	}
	for {
		l := len(result)
		if l >= length {
			break
		}
		r := len(initial)
		if l+r > length {
			result = append(result, initial[:length-l]...)
			break
		}
		result = append(result, initial...)
	}
	return
}

// Repeat returns a slice of elements that contains the given source repeated count times.
//
// 将切片 source 的元素重复 count 次。
func Repeat[T any](source []T, count int) (result []T) {
	if count <= 0 {
		return nil
	}
	for i := 0; i < count; i++ {
		result = append(result, source...)
	}
	return
}

// Reduce returns the result of reducing the slices using the function f.
//
// 收敛，迭代切片，使用函数 f 对每个元素进行计算，并将之前迭代的计算结果作为下一个迭代的输入参数，返回计算结果。
func Reduce[S ~[]E, E any, R any](ss S, f func(R, E) R, initial R) R {
	for _, s := range ss {
		initial = f(initial, s)
	}
	return initial
}

// Cast to convert the slice elements to a new type.
//
// 转换切片元素类型，使用函数 cast 对每个元素进行转换，返回新的元素值 v 和 omitted 标记, omitted 为 true，则跳过该元素。
// 新切片的长度小于（如果有元素被跳过）或等于原切片的长度。
func Cast[E, R any](ss []E, cast func(index int, ele E) (v R, omitted bool)) (result []R) {
	if ss == nil {
		return
	}
	l := len(ss)
	if l == 0 {
		return []R{}
	}
	result = make([]R, 0, l)
	for i, s := range ss {
		if v, o := cast(i, s); !o {
			result = append(result, v)
		}
	}
	return
}

// MapFunc to convert slice to map, use f to convert each element to a new key and value, existing
// key-value pairs be overwritten if priority is greater, the same priority, the existing key is higher.
//
// 切片转换为映射，使用函数 f 对每个元素进行转换，返回转换结果。
// 使用 cast 函数的返回值 priority 决定是否覆盖已存在的键值对，priority 越大，优先级越高，优先级高的覆盖低的，相同时已存在的优先。
func MapFunc[E any, K comparable, V any](ss []E, cast func(index int, ele E) (key K, value V, priority int)) (result map[K]V, err error) {
	if ss == nil {
		return
	}
	var ps = make(map[K]int, len(ss))
	for i, s := range ss {
		if result == nil {
			result = make(map[K]V)
		}
		k, v, p := cast(i, s)
		if z, o := ps[k]; !o || p > z {
			result[k] = v
			ps[k] = p
		}
	}
	return
}

// Map convert slice to map, use f to convert each element to a new key and value, existing key-value pairs will be overwritten.
//
// 以覆盖模式将切片转换为映射，使用函数 cast 对每个元素进行转换，以返回值 key 和 value 作为新的键值对。已存在的键值对将被覆盖。
func Map[E any, K comparable, V any](ss []E, f func(index int, ele E) (key K, value V)) (result map[K]V, err error) {
	if ss == nil {
		return
	}
	for i, s := range ss {
		if result == nil {
			result = make(map[K]V)
		}
		k, v := f(i, s)
		result[k] = v
	}
	return
}

// Apply to apply the function f to each element and f can modify the element value.
//
// 使用函数 f 对每个元素进行操作，允许 f 修改元素的值
func Apply[R ~*E, E any](ss []E, f func(index int, ele R)) {
	if len(ss) == 0 {
		return
	}
	for i := range ss {
		f(i, &ss[i])
	}
}

// Replace to replace the original element if f returns a new value and omitted is false.
//
// 替换元素，函数 f 返回值 omitted 为 false 则替换原元素。
func Replace[E any](ss []E, f func(index int, ele E) (v E, omitted bool)) {
	if len(ss) == 0 {
		return
	}
	for i, e := range ss {
		if v, omitted := f(i, e); !omitted {
			ss[i] = v
		}
	}
}

// Lookup to look up the element when the f return true.
//
// 迭代切片，使用函数 f 对每个元素进行判断，返回判断为 true 的第一个元素。
func Lookup[E any](ss []E, f func(index int, element E) (found bool)) (found E, index int) {
	index = -1
	for i, s := range ss {
		if f(i, s) {
			found = s
			index = i
			break
		}
	}
	return
}

// Pick to pick the elements when the pick return true.
//
// 挑选
func Pick[E any](pick func(index int, ele E) bool, vs ...E) (result []E) {
	for i, v := range vs {
		if pick(i, v) {
			result = append(result, v)
		}
	}
	return
}
