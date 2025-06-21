package data_test

import (
	"fmt"

	"github.com/keepitlight/golang/data"
)

func ExampleChunk() {
	c, e := data.Chunk(3, func(start, end int) (next bool, err error) {
		fmt.Println(start, end)
		next = start < 12
		// Here the boundary value is checked and stopped after the next loop call.
		// 此处检查边界值 < 12 可进入下一循环，所以下次调用后才退出
		//
		// Note: avoid dead loop
		// 注意: 避免死循环，仅在必要的时候为 next 赋值为 true
		return
	})
	fmt.Println(c, e)
	// Output:
	// 0 2
	// 3 5
	// 6 8
	// 9 11
	// 12 14
	// 5 <nil>
}

func ExampleOffset() {
	c, e := data.Offset(3, func(offset int) (next bool, err error) {
		fmt.Println(offset)
		next = offset < 12
		// Here the boundary value is checked and stopped after the next loop call.
		// 此处检查边界值 < 12 可进入下一循环，所以下次调用后才退出
		//
		// Note: avoid dead loop
		// 注意: 避免死循环，仅在必要的时候为 next 赋值为 true
		return
	})
	fmt.Println(c, e)
	// Output:
	// 0
	// 3
	// 6
	// 9
	// 12
	// 5 <nil>
}

func ExamplePaged() {
	c, e := data.Paged(3, func(number, start, end int) (next bool, err error) {
		fmt.Printf("p%d: [%d, %d]\n", number, start, end)
		next = number < 4
		// Here the boundary value is checked and stopped after the next loop call.
		// 此处检查边界值 < 4 可进入下一循环，所以下次调用后才退出
		//
		// Note: avoid dead loop
		// 注意: 避免死循环，仅在必要的时候为 next 赋值为 true
		return
	})
	fmt.Println(c, e)
	// Output:
	// p1: [0, 2]
	// p2: [3, 5]
	// p3: [6, 8]
	// p4: [9, 11]
	// 4 <nil>
}
