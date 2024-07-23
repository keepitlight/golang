package ranges_test

import (
	"fmt"
	"time"

	"github.com/keepitlight/golang"
	"github.com/keepitlight/golang/ranges"
)

func ExampleBounds() {
	n := time.Now()
	s := golang.AddTime(n, -1, 0, 0, 0)
	r := ranges.Time(&s, &n)
	l, u := ranges.Bounds(r)
	fmt.Println(u.Sub(*l))
	// output:
	// 1h0m0s
}

func ExampleBetween() {
	r := ranges.Between(90, 100)
	l, u := ranges.Bounds(r)
	fmt.Println(l, u)
	fmt.Println(ranges.In(r, 50), ranges.In(r, 95))
	fmt.Println(ranges.Pick(r, 1, 2, 3, 4, 5, 99, 101))
	fmt.Println(ranges.Unpick(r, 1, 2, 3, 4, 5, 99, 101))
	// output:
	// 90 100
	// false true
	// [99]
	// [1 2 3 4 5 101]
}

func ExampleDuration() {
	s, e := time.Now(), time.Now().Add(time.Hour)
	// 实际运行时有运行时时间差
	t := ranges.Time(&s, &e)
	fmt.Println(ranges.Duration(t).Round(time.Second))
	// 舍掉时间计算的时间差
	// output:
	// 1h0m0s
}

func ExampleSince() {
	s := time.Now().Add(-time.Hour)
	// 实际运行时有运行时时间差
	t := ranges.Since(&s)
	fmt.Println(ranges.Duration(t).Round(time.Second))
	// 舍掉时间计算的时间差
	// output:
	// 1h0m0s
}

func ExampleMerge() {
	s, _ := time.Parse(golang.DateTimeZone, "2023-07-07 08:10:00 UTC")
	e1 := s.Add(time.Hour)
	e2 := s.Add(-time.Hour)
	i := s.Add(10 * time.Minute)
	t1 := ranges.Time(&s, &e1, ranges.WithInitialValue(&i), ranges.WithInterval(time.Minute))
	t2 := ranges.Time(&s, &e2, ranges.WithInitialValue(&i), ranges.WithInterval(time.Minute))
	m, ok := ranges.Merge[*time.Time](t1, t2)
	l, u := m.Bounds()
	fmt.Println(l, u, ok)
	// output:
	// 2023-07-07 07:10:00 +0000 UTC 2023-07-07 09:10:00 +0000 UTC true
}
