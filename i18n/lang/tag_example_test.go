package lang

import (
	"fmt"
)

func ExampleEnsure() {
	z0, o0 := Ensure("zh")
	z1, o1 := Ensure("zh-CN")
	z2, o2 := Ensure("zh-TW")
	z3, o3 := Ensure("zh-HK")
	z4, o4 := Ensure("zh_SG")
	z5, o5 := Ensure("zh_MO")
	e0, k0 := Ensure("en")
	e1, k1 := Ensure("en_US")
	e2, k2 := Ensure("en_GB")

	fmt.Println(z0, o0, z1, o1, z2, o2, z3, o3, z4, o4, z5, o5)
	fmt.Println(e0, k0, e1, k1, e2, k2)
	fmt.Println(EN, ZH, US, CN)

	// Output:
	// zh true zh-cn true zh-tw true zh-hk true zh-sg true zh-mo true
	// en true en-us true en-gb true
	// en zh en-us zh-cn
}

func ExampleSubTags() {
	tags, ok := SubTags("zh-hans_CN")
	fmt.Println(ok)
	for _, tag := range tags {
		fmt.Println(tag)
	}

	// Output:
	// true
	// zh
	// hans
	// cn
}
