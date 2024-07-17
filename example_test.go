package golang_test

import (
	"fmt"
	"time"

	"github.com/keepitlight/golang"
)

func ExampleValue() {
	n := time.Now()
	s := golang.AddTime(n, -1, 0, 0, 0)
	r := golang.TimeRange(&s, &n)
	l, u := golang.Value[time.Time](r)
	fmt.Println(u.Sub(l))
	// output:
	// 1h0m0s
}

func ExampleDuration() {
	s, e := time.Now(), time.Now().Add(time.Hour) // 实际运行时有运行时时间差
	t := golang.TimeRange(&s, &e)
	fmt.Println(golang.Duration(t).Round(time.Second)) // 舍掉时间计算的时间差
	// output:
	// 1h0m0s
}

func ExampleAddTime() {
	s := time.Now()

	e1 := golang.AddTime(s, 1, 0, 0, 0)
	e2 := golang.AddTime(s, -1, 0, 0, 0)

	// There was an additional overhead in runtime time.
	// 有增加额外的运行时时间
	c1 := time.Now().Add(time.Hour)
	c2 := time.Now().Add(-time.Hour)

	fmt.Println(e1.Sub(s))
	fmt.Println(e2.Sub(s))

	fmt.Println(e1.Sub(s) == c1.Sub(s))
	fmt.Println(e2.Sub(s) == c2.Sub(s))
	// output:
	// 1h0m0s
	// -1h0m0s
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

func ExampleBetween() {
	r := golang.Between(90, 100)
	fmt.Println(r.GetLower())
	fmt.Println(r.GetUpper())
	fmt.Println(r.In(50), r.In(95))
	fmt.Println(r.Pick(1, 2, 3, 4, 5, 99, 101))
	fmt.Println(r.Unpick(1, 2, 3, 4, 5, 99, 101))
	// output:
	// 90
	// 100
	// false true
	// [99]
	// [1 2 3 4 5 101]
}
