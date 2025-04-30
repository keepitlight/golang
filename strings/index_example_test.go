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

func ExampleSnakeCase() {
	fmt.Println(SnakeCase("primaryKey"))
	// output:
	// primary_key
}

func ExampleUpperCase() {
	fmt.Println(UpperCase("primaryKey"))
	// output:
	// PRIMARY_KEY
}

func ExampleUnquote() {
	fmt.Println(Unquote("'primaryKey'", '\''))
	fmt.Println(Unquote("`tag`", '`'))
	// output:
	// primaryKey true 39
	// tag true 96
}

func ExampleUnquoteRune() {
	fmt.Println(UnquoteRune("'primaryKey'", '\''))
	fmt.Println(UnquoteRune("\"primaryKey\"", '"'))
	// output:
	// primaryKey true 39
	// primaryKey true 34
}

func ExampleUnbracketed() {
	fmt.Println(Unbracketed("<primaryKey>", '<', '>'))
	fmt.Println(Unbracketed("/regular expression/", '/'))
	// output:
	// primaryKey true
	// regular expression true
}

func ExampleCase() {
	fmt.Println(Case("primaryKey", UpperCaseMode))
	// output:
	// PRIMARY_KEY
}
