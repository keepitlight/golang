package golang

// Range declares a range of type T
//
// 范围，值域
type Range[T any] interface {
	// Bounds returns the lower and upper bounds of the range
	//
	// 返回范围的下限和上限
	Bounds() (lower, upper T)
	// Comparer returns a function that compares two values of type T
	//
	// 返回一个比较两个值类型的函数
	Comparer() func(a, b T) int
}

// Interval declares an interval of type T
//
// 区间，表示一组连续值的范围
type Interval[T any] interface {
	Range[T]
	// Next return current value and advances the interval to the next value
	//
	// 返回当前值并移动区间到下一个值
	Next() (current T, end bool)
	// Init initializes or resets the interval
	//
	// 初始化或重置区间
	Init()
}
