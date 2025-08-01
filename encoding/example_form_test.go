package encoding

import "fmt"

func ExampleFormMarshal() {
	type Form struct {
		Key string `form:"key"`
	}
	f := Form{
		Key: "value",
	}
	data, _ := FormMarshal(f)
	fmt.Println(string(data))
	// Output:
	// key=value
}
func ExampleFormUnmarshal() {
	type Form struct {
		Key string `form:"key"`
	}
	data := []byte("key=value")
	f := Form{}
	_ = FormUnmarshal(data, &f)
	fmt.Println(f.Key)
	// Output:
	// value
}

func ExampleJsonMarshal() {
	type Json struct {
		Key string `json:"key"`
	}
	j := Json{
		Key: "value",
	}
	data, _ := JsonMarshal(j)

	fmt.Println(string(data))
	// Output:
	// {"key":"value"}
}
