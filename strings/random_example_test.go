package strings

import "fmt"

func ExampleRandomString() {
	for i := Printable; i < Base64; i++ {
		fmt.Printf(string(RandomString(i, 50)))
	}
	// Output:
	// 0123456789
	// 0123456789
}
