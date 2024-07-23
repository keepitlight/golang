package golang_test

import (
	"fmt"
	"time"

	"github.com/keepitlight/golang"
)

func ExampleAddTime() {
	s := time.Now()

	e1 := golang.AddTime(s, 1, 30, 0, 0)
	e2 := golang.AddTime(s, -1, -30, 0, 0)

	// There was an additional overhead in runtime time.
	// 有增加额外的运行时时间
	c1 := time.Now().Add(time.Hour)
	c2 := time.Now().Add(-time.Hour)

	fmt.Println(e1.Sub(s))
	fmt.Println(e2.Sub(s))

	fmt.Println(e1.Sub(s) == c1.Sub(s))
	fmt.Println(e2.Sub(s) == c2.Sub(s))

	// output:
	// 1h30m0s
	// -1h30m0s
	// false
	// false
}

func ExampleAddDays() {
	s := time.Now()

	d1 := golang.Day + 12*time.Hour + 40*time.Minute

	e1 := golang.AddDays(s, golang.Days(d1))
	e2 := golang.AddDays(s, -golang.Days(d1))

	fmt.Println(e1.Sub(s).Round(time.Second))
	fmt.Println(e2.Sub(s).Round(time.Second))

	// output:
	// 36h40m0s
	// -36h40m0s
}
