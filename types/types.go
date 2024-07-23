package types

import (
	"reflect"
	"strings"
)

// PkgPath is a package path.
//
// GO 包路径
type PkgPath string

func (p PkgPath) Contain(path string) bool {
	return strings.Contains(string(p), path)
}
func (p PkgPath) StartsWith(path string) bool {
	return strings.HasPrefix(string(p), path)
}
func (p PkgPath) EndsWith(path string) bool {
	return strings.HasSuffix(string(p), path)
}
func (p PkgPath) Equal(path string) bool {
	return string(p) == path
}
func (p PkgPath) Empty() bool {
	return len(p) < 1
}

// Type is type information.
//
// 类型信息
type Type struct {
	Name  string            `json:"name,omitempty"`  // 变量名
	Value any               `json:"value,omitempty"` // 变量值
	Zero  any               `json:"zero,omitempty"`  // 变量同类型的零值
	Kind  reflect.Kind      `json:"kind,omitempty"`  // 变量类型
	Tag   reflect.StructTag `json:"tag,omitempty"`   // 字段的 tag，仅适用于结构的字段

	key    *Type            // map 的键类型
	elem   *Type            // 数组、切片或通道的元素类型
	fields map[string]*Type // 结构体的导出字段类型值
	t      reflect.Type
	v      reflect.Value
	p      bool // isPointer
	i      bool // isInterface
	n      bool // isNil
}

func (t *Type) IsPointer() bool {
	return t.p
}
func (t *Type) IsInterface() bool {
	return t.i
}
func (t *Type) IsNil() bool {
	return t.n
}
func (t *Type) IsScalar() bool {
	return t.Kind == reflect.Int ||
		t.Kind == reflect.Int8 ||
		t.Kind == reflect.Int16 ||
		t.Kind == reflect.Int32 ||
		t.Kind == reflect.Int64 ||
		t.Kind == reflect.Uint ||
		t.Kind == reflect.Uint8 ||
		t.Kind == reflect.Uint16 ||
		t.Kind == reflect.Uint32 ||
		t.Kind == reflect.Uint64 ||
		t.Kind == reflect.Float32 ||
		t.Kind == reflect.Float64 ||
		t.Kind == reflect.Bool ||
		t.Kind == reflect.String
}
func (t *Type) IsNumber() bool {
	return t.Kind == reflect.Int ||
		t.Kind == reflect.Int8 ||
		t.Kind == reflect.Int16 ||
		t.Kind == reflect.Int32 ||
		t.Kind == reflect.Int64 ||
		t.Kind == reflect.Uint ||
		t.Kind == reflect.Uint8 ||
		t.Kind == reflect.Uint16 ||
		t.Kind == reflect.Uint32 ||
		t.Kind == reflect.Uint64 ||
		t.Kind == reflect.Float32 ||
		t.Kind == reflect.Float64
}
func (t *Type) IsFloat() bool {
	return t.Kind == reflect.Float32 ||
		t.Kind == reflect.Float64
}
func (t *Type) IsComplex() bool {
	return t.Kind == reflect.Complex64 ||
		t.Kind == reflect.Complex128
}
func (t *Type) IsUnsigned() bool {
	return t.Kind == reflect.Uint ||
		t.Kind == reflect.Uint8 ||
		t.Kind == reflect.Uint16 ||
		t.Kind == reflect.Uint32 ||
		t.Kind == reflect.Uint64
}
func (t *Type) IsSigned() bool {
	return t.Kind == reflect.Int ||
		t.Kind == reflect.Int8 ||
		t.Kind == reflect.Int16 ||
		t.Kind == reflect.Int32 ||
		t.Kind == reflect.Int64
}
func (t *Type) IsString() bool {
	return t.Kind == reflect.String
}
func (t *Type) IsBool() bool {
	return t.Kind == reflect.Bool
}
func (t *Type) IsFunc() bool {
	return t.Kind == reflect.Func
}
func (t *Type) Comparable() bool {
	return t.t.Comparable()
}
func (t *Type) Bits() int {
	return t.t.Bits()
}
func (t *Type) PkgPath() PkgPath {
	p := t.t.PkgPath()
	return PkgPath(p)
}
func (t *Type) Type() reflect.Type {
	return t.t
}

func (t *Type) AsSlice() (elem *Type, ok bool) {
	if t.Kind == reflect.Slice {
		return t.elem, true
	}
	return nil, false
}
func (t *Type) AsArray() (elem *Type, ok bool) {
	if t.Kind == reflect.Array {
		return t.elem, true
	}
	return nil, false
}
func (t *Type) AsChan() (elem *Type, ok bool) {
	if t.Kind == reflect.Chan {
		return t.elem, true
	}
	return nil, false
}
func (t *Type) AsMap() (key, elem *Type, ok bool) {
	if t.Kind == reflect.Map {
		return t.key, t.elem, true
	}
	return nil, nil, false
}
func (t *Type) AsStruct() (fields map[string]*Type, ok bool) {
	if t.Kind == reflect.Struct {
		return t.fields, true
	}
	return nil, false
}

// TypeOf returns the type information of the variable.
//
// 获取变量的类型信息，如果 v 为 nil，返回 nil，可以直接传入 reflect.Type 或 reflect.Value
func TypeOf(v any) *Type {
	if v == nil {
		return nil
	}
	if f, ok := v.(reflect.Value); ok {
		return From(f)
	}
	if t, ok := v.(reflect.Type); ok {
		return parse(t, false)
	} else {
		return From(reflect.ValueOf(v))
	}
}

// Parse to get the type information from reflect.Type.
//
// 获取变量的类型信息
func Parse(t reflect.Type) *Type {
	return parse(t, false)
}

func stripType(v reflect.Type) (r reflect.Type, isPointer, isInterface bool) {
	r = v
	for r.Kind() == reflect.Pointer || r.Kind() == reflect.Interface {
		if r.Kind() == reflect.Pointer && !isPointer {
			isPointer = true
		}
		if r.Kind() == reflect.Interface && !isInterface {
			isInterface = true
		}
		r = r.Elem()
	}
	return
}
func stripValue(v reflect.Value) (r reflect.Value, isPointer, isInterface bool) {
	r = v
	for r.Kind() == reflect.Pointer || r.Kind() == reflect.Interface {
		if r.Kind() == reflect.Pointer && !isPointer {
			isPointer = true
		}
		if r.Kind() == reflect.Interface && isInterface {
			isInterface = true
		}
		r = r.Elem()
	}
	return
}

func parse(t reflect.Type, skipStruct bool) *Type {
	if t == nil {
		return nil
	}
	var ip, ii bool
	t, ip, ii = stripType(t)
	var k, e *Type
	var fs map[string]*Type
	switch {
	case t.Kind() == reflect.Slice, t.Kind() == reflect.Array, t.Kind() == reflect.Chan:
		e = parse(t.Elem(), false)
	case t.Kind() == reflect.Map:
		k = parse(t.Key(), true) // 不能为结构
		e = parse(t.Elem(), false)
	case !skipStruct && t.Kind() == reflect.Struct:
		fs = make(map[string]*Type)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if !f.IsExported() {
				continue
			}
			if f.Anonymous {
				ff := parse(f.Type, false)
				if ff != nil {
					for x, v := range ff.fields {
						fs[x] = v
					}
				}
			} else {
				ft := parse(f.Type, false)
				ft.Tag = f.Tag // 保留 tag
				fs[f.Name] = ft
			}
		}
	}
	z := reflect.Zero(t).Interface()
	return &Type{
		p:      ip,
		i:      ii,
		n:      false,
		Name:   t.Name(),
		Kind:   t.Kind(),
		Value:  z,
		Zero:   z,
		key:    k,
		elem:   e,
		fields: fs,
		t:      t,
	}
}

// From to get the type information from reflect.Value.
//
// 根据值信息获取变量的类型信息
func From(v reflect.Value) *Type {
	var ip, ii bool
	v, ip, ii = stripValue(v)
	r := parse(v.Type(), true)
	if ip {
		r.p = ip
	}
	if ii {
		r.i = ii
	}
	r.n = v.IsNil()
	r.Value = v.Interface()
	r.v = v
	if r.Kind == reflect.Struct {
		t := v.Type()
		var fs map[string]*Type
		for i := 0; i < v.NumField(); i++ {
			f := t.Field(i)
			if !f.IsExported() {
				continue
			}
			x := v.Field(i)
			if f.Anonymous {
				for k, z := range From(x).fields {
					fs[k] = z
				}
			} else {
				fs[f.Name] = From(x)
			}
		}
		r.fields = fs
	}
	return r
}

func of(v reflect.Value) (name string, fields map[string]any) {
	t := v.Type()
	t, _, _ = stripType(t)
	if t.Kind() != reflect.Struct {
		return "", nil
	}
	name = t.Name()
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		x := v.Field(i)
		if fields == nil {
			fields = make(map[string]any)
		}
		if f.Anonymous {
			if x.Kind() != reflect.Struct {
				continue
			}
			_, fs := of(x)
			for k, z := range fs {
				fields[k] = z
			}
		} else {
			fields[f.Name] = x.Interface()
		}
	}
	return
}

func declare(v reflect.Value) (tags map[string]reflect.StructTag) {
	t := v.Type()
	t, _, _ = stripType(t)
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		x := v.Field(i)
		if tags == nil {
			tags = make(map[string]reflect.StructTag)
		}
		if f.Anonymous {
			if x.Kind() != reflect.Struct {
				continue
			}
			fs := declare(x)
			for k, z := range fs {
				tags[k] = z
			}
		} else {
			tags[f.Name] = f.Tag
		}
	}
	return
}
