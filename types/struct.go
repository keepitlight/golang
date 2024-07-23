package types

import "reflect"

// Struct returns the name, fields and tags of the struct, nil if v is nil or not struct
//
// 获取结构体的名字和字段值，v 是 nil 或非结构体返回 nil
func Struct(v any) (name string, fields map[string]any, tags map[string]reflect.StructTag) {
	rv := reflect.ValueOf(v)
	n, fs := of(rv)
	if fs == nil {
		return "", nil, nil
	}
	ts := declare(rv)
	return n, fs, ts
}

// IsStruct check whether v is struct
//
// 判断 v 是否是结构体
func IsStruct(v any) bool {
	t := reflect.TypeOf(v)
	t, _, _ = stripType(t)
	return t.Kind() != reflect.Struct
}

// Of returns the name and fields of the struct, nil if v is nil or not struct
//
// 获取结构体的名字和字段值，非结构体返回 nil
func Of(v any) (name string, fields map[string]any) {
	return of(reflect.ValueOf(v))
}

// TagsOf returns the tags of the struct, nil if v is nil or not struct
//
// 获取结构体的 tag，如果 v 是 nil 或者不是结构体，返回 nil
func TagsOf(v any) (tags map[string]reflect.StructTag) {
	if v == nil {
		return nil
	}
	return declare(reflect.ValueOf(v))
}

// TagOf returns the tag of the struct field, false if v is not struct or field not found
//
// 获取结构体指定字段的 tag，不是结构体，或者字段不存在，found 返回 false
func TagOf(v any, name string) (tag reflect.StructTag, found bool) {
	t := reflect.TypeOf(v)
	t, _, _ = stripType(t)
	if t.Kind() != reflect.Struct {
		// 不是结构体
		return "", false
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		if f.Name == name {
			return f.Tag, true
		}
	}
	return "", false
}
