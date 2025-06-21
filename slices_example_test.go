package golang_test

import (
	"fmt"
	"github.com/keepitlight/golang"
)

func ExampleUniqueString() {
	v := golang.UniqueString(
		"a", "b", "c", "d", "a",
		"a", "b", "c", "d", "d",
		"e", "f", "g", "h", "i",
		"e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n",
		"j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s",
		"o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x",
		"t", "u", "v", "w", "x",
		"y", "z", "x", "y", "z",
	)
	fmt.Println(v)
	// Output:
	// [a b c d e f g h i j k l m n o p q r s t u v w x y z]
}

func ExampleFoldString() {
	v := golang.FoldString(
		"a", "b", "c", "d", "A",
		"A", "B", "C", "D", "d",
		"E", "F", "G", "H", "I",
		"e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n",
		"J", "K", "L", "M", "N",
		"o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x",
		"y", "z", "X", "Y", "Z",
	)
	fmt.Println(v)
	// Output:
	// [a b c d E F G H I j k l m n o p q r s t u v w x y z]
}

func ExamplePick() {
	v := golang.Pick(func(index int, ele int) bool { return ele%2 == 0 }, 1, 2, 3, 4, 5)
	fmt.Println(v)
	// Output:
	// [2 4]
}

func ExampleSlices() {
	v := golang.Slices(12, 1, 2, 3, 4, 5)
	fmt.Println(v)
	// Output:
	// [1 2 3 4 5 1 2 3 4 5 1 2]
}

func ExampleRepeat() {
	v := golang.Repeat([]int{1, 2, 3}, 3)
	fmt.Println(v)
	// Output:
	// [1 2 3 1 2 3 1 2 3]
}

func ExampleReduce() {
	sum := func(prev, next int) int { return prev + next }
	v := golang.Reduce([]int{1, 2, 3, 4, 5}, sum, 10)
	fmt.Println(v)
	// Output:
	// 25
}

func ExampleMap() {
	v, err := golang.Map([]int{1, 2, 3, 4, 5}, func(index int, ele int) (key string, value int) {
		return fmt.Sprintf("key%d", ele), ele * ele
	})
	fmt.Println(v, err)
	// Output:
	// map[key1:1 key2:4 key3:9 key4:16 key5:25] <nil>
}

func ExampleMapFunc() {
	type Preferred int
	const (
		higherIndexPreferred  Preferred = iota // higher index first
		smallerIndexPreferred                  // smaller index first
		greaterValuePreferred                  // higher value first
		smallerValuePreferred                  // smaller value first
	)
	type element struct {
		key   string
		value int
	}
	var preferred = map[string]Preferred{
		"key1": smallerIndexPreferred,
		"key2": higherIndexPreferred,
		"key3": smallerValuePreferred,
		"key4": greaterValuePreferred,
	}
	s := []*element{
		{"key1", 1},
		{"key2", 2},
		{"key3", 3},
		{"key4", 4},
		{"key1", 5},
		{"key2", 6},
		{"key3", 7},
		{"key4", 8},
	}
	v, err := golang.MapFunc(s, func(index int, ele *element) (key string, value int, priority int) {
		key, value = ele.key, ele.value
		switch preferred[ele.key] {
		case smallerIndexPreferred:
			priority = -index
		case smallerValuePreferred:
			priority = -ele.value
		case greaterValuePreferred:
			priority = ele.value
		default:
			priority = index
		}
		return
	})
	fmt.Println(v, err)
	// Output:
	// map[key1:1 key2:6 key3:3 key4:8] <nil>
}

func ExampleCast() {
	v := golang.Cast([]int{1, 2, 3, 4, 5}, func(index int, ele int) (v string, omitted bool) {
		return fmt.Sprintf("key%d", ele), ele%2 == 0
	})
	fmt.Println(v)
	// Output:
	// [key1 key3 key5]
}
func ExampleUniqueFunc() {
	v := golang.UniqueFunc([]int{1, 1, 2, 3, 4, 5, 4, 5}, func(a, b int) bool {
		return a == b
	})
	fmt.Println(v)
	// Output:
	// [1 2 3 4 5]
}

func ExampleReplace() {
	v := []int{1, 2, 3, 4, 5}
	golang.Replace(v, func(index int, ele int) (v int, omitted bool) {
		return ele * ele, ele%2 == 0 // omit even values
	})
	fmt.Println(v)
	// Output:
	// [1 2 9 4 25]
}
func ExampleLookup() {
	v, index := golang.Lookup([]int{1, 2, 3, 4, 5}, func(index int, element int) (found bool) {
		return element == 3
	})
	fmt.Println(v, index)
	// Output:
	// 3 2
}

func ExampleApply() {
	v := []int{1, 2, 3, 4, 5}
	golang.Apply(v, func(index int, ele *int) {
		*ele *= 2
	})
	fmt.Println(v)
	// Output:
	// [2 4 6 8 10]
}
