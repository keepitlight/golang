package golang_test

import (
	"fmt"
	"github.com/keepitlight/golang"
	"slices"
)

func ExampleRand() {
	var b0 = make([]byte, 10)
	var b1 = make([]byte, 10)
	var b2 = make([]byte, 10)
	var b3 = make([]byte, 10)
	golang.Rand(b1, golang.UsePCG(0))
	golang.Rand(b2, golang.UseChaCha8())
	golang.Rand(b3, golang.CryptoRand())
	fmt.Println(len(b1), len(b2), len(b3))
	fmt.Println(slices.Equal(b1, b0), slices.Equal(b2, b0), slices.Equal(b3, b0))
	fmt.Println(slices.Equal(b1, b2), slices.Equal(b1, b3), slices.Equal(b2, b3))
	// Output:
	// 10 10 10
	// false false false
	// false false false
}

func ExampleShuffle() {
	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	golang.Shuffle(len(v1), func(i, j int) {
		v1[i], v1[j] = v1[j], v1[i]
	}, golang.UsePCG(0))

	golang.Shuffle(len(v2), func(i, j int) {
		v2[i], v2[j] = v2[j], v2[i]
	}, golang.UseChaCha8())

	golang.Shuffle(len(v3), func(i, j int) {
		v3[i], v3[j] = v3[j], v3[i]
	}, golang.CryptoRand())

	//fmt.Println(v1, v2, v3)
	fmt.Println(
		slices.Equal(v1, v2),
		slices.Equal(v1, v3),
		slices.Equal(v2, v3),
	)
	slices.Sort(v1)
	slices.Sort(v2)
	slices.Sort(v3)
	fmt.Println(
		slices.Equal(v1, v2),
		slices.Equal(v1, v3),
		slices.Equal(v2, v3),
	)
	// Output:
	// false false false
	// true true true
}

func ExampleShuffleSlice() {
	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	golang.ShuffleSlice(v1, golang.UsePCG(0))
	golang.ShuffleSlice(v2, golang.UseChaCha8())
	golang.ShuffleSlice(v3, golang.CryptoRand())

	//fmt.Println(v1, v2, v3)
	fmt.Println(
		slices.Equal(v1, v2),
		slices.Equal(v1, v3),
		slices.Equal(v2, v3),
	)
	slices.Sort(v1)
	slices.Sort(v2)
	slices.Sort(v3)
	fmt.Println(
		slices.Equal(v1, v2),
		slices.Equal(v1, v3),
		slices.Equal(v2, v3),
	)
	// Output:
	// false false false
	// true true true
}

func ExampleShuffledSlice() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v1 := golang.ShuffledSlice(s, golang.UsePCG(0))
	v2 := golang.ShuffledSlice(s, golang.UseChaCha8())

	fmt.Println(slices.Equal(s, v1), slices.Equal(s, v2), slices.Equal(v1, v2))
	slices.Sort(v1)
	slices.Sort(v2)
	fmt.Println(slices.Equal(s, v1), slices.Equal(s, v2), slices.Equal(v1, v2))
	// Output:
	// false false false
	// true true true
}
