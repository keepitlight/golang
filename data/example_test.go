package data_test

import (
	"fmt"

	"github.com/keepitlight/golang/data"
)

func ExampleChunk() {
	_, _ = data.Chunk(3, func(start, end int) (next bool, err error) {
		if start < 12 {
			// Note: pay attention to check the boundary value
			// 注意：边界值要注意
			fmt.Println(start, end)
			// Note: avoid dead loop
			// 注意: 避免死循环，仅在必要的时候为 next 赋值为 true
			next = true
		}
		return
	})
	// Output:
	// 0 3
	// 3 6
	// 6 9
	// 9 12
}

func ExampleOffset() {
	_, _ = data.Offset(3, func(offset int) (next bool, err error) {
		if offset < 12 {
			// Note: pay attention to check the boundary value
			// 注意：边界值要注意
			fmt.Println(offset)
			// Note: avoid dead loop
			// 注意: 避免死循环，仅在必要的时候为 next 赋值为 true
			next = true
		}
		return
	})
	// Output:
	// 0
	// 3
	// 6
	// 9
}

func ExamplePaged() {
	_, _ = data.Paged(3, func(number, start, end int) (next bool, err error) {
		if number < 4 {
			// Note: pay attention to check the boundary value
			// 注意：边界值要注意
			fmt.Println(number, start, end)
			// Note: avoid dead loop
			// 注意: 避免死循环，仅在必要的时候为 next 赋值为 true
			next = true
		}
		return
	})
	// Output:
	// 1 0 3
	// 2 3 6
	// 3 6 9
}
