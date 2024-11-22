package strings

import "fmt"

func ExamplePascalCase() {
	fmt.Println(PascalCase("primary_key"))
	// output:
	// PrimaryKey
}
func ExampleCamelCase() {
	fmt.Println(CamelCase("primary_key"))
	// output:
	// primaryKey
}
