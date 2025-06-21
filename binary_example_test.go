package golang_test

import (
	"fmt"
	"github.com/keepitlight/golang"
)

func ExampleBigEndianBytes() {
	b := golang.BigEndianBytes(0x1234567890ABCDEF)
	fmt.Printf("%x %x %x %x %x %x %x %x\n", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7])
	// Output:
	// 12 34 56 78 90 ab cd ef
}

func ExampleLittleEndianBytes() {
	b := golang.LittleEndianBytes(0x1234567890ABCDEF)
	fmt.Printf("%x %x %x %x %x %x %x %x\n", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7])
	// Output:
	// ef cd ab 90 78 56 34 12
}
