package golang_test

import (
	"fmt"
	"time"

	"github.com/keepitlight/golang"
)

func ExampleTuple() {
	t := golang.Tuple(0, "a")
	fmt.Println(t.Value())
	// Output:
	// 0 a
}

func ExampleTuple3() {
	t := golang.Tuple3(0, "a", true)
	fmt.Println(t.Value())
	// Output:
	// 0 a true
}

func ExampleTuple4() {
	t := golang.Tuple4(0, "a", true, 1.1)
	fmt.Println(t.Value())
	// Output:
	// 0 a true 1.1
}

func ExampleTuple5() {
	n := time.Date(2024, 8, 8, 8, 8, 8, 888, time.UTC)
	t := golang.Tuple5(0, "a", true, 1.1, n)
	fmt.Println(t.Value())
	// Output:
	// 0 a true 1.1 2024-08-08 08:08:08.000000888 +0000 UTC
}
