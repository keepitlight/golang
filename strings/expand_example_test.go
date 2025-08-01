package strings

import "fmt"

func ExampleExpand() {
	r1 := Expand("primaryKey;uniqueIndex:idx;size:20", ";", ":", nil)
	r2 := Expand("primaryKey;uniqueIndex:idx;size:20", ";", ":", SnakeCase)
	r3 := Expand("primaryKey;uniqueIndex:idx;size:20", ";", ":", UpperCase)
	r4 := Expand("primaryKey;uniqueIndex:idx;size:20", ";", ":", CamelCase)
	r5 := Expand("primaryKey;uniqueIndex:idx;size:20", ";", ":", PascalCase)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)
	// output:
	// map[primaryKey: size:20 uniqueIndex:idx]
	// map[primary_key: size:20 unique_index:idx]
	// map[PRIMARY_KEY: SIZE:20 UNIQUE_INDEX:idx]
	// map[primaryKey: size:20 uniqueIndex:idx]
	// map[PrimaryKey: Size:20 UniqueIndex:idx]
}
