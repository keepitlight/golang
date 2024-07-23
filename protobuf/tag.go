package protobuf

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/keepitlight/golang/types"
)

const (
	TagName    types.TagName = "protobuf"
	KeyTagName types.TagName = "protobuf_key"
	ValTagName types.TagName = "protobuf_val"
)

func init() {
	types.RegisterTagParser(TagName, TagParse)
	types.RegisterTagParser(KeyTagName, TagParse)
	types.RegisterTagParser(ValTagName, TagParse)
}

type (
	Version int
	Type    int
)

const (
	VersionUndefined Version = iota // 版本未定义
	V1                              // 一般不使用，仅作占位符
	V2                              // proto2 语法
	V3                              // proto3 语法
)

const (
	Bytes Type = iota
	VarInt
)

type Tag struct {
	Type     Type    `json:"type,omitempty"`     // 1. 编码类型 "varint" "bytes"
	Index    int     `json:"index,omitempty"`    // 2. 位置索引
	Optional bool    `json:"optional,omitempty"` // 3. 是否可选 "opt"
	Repeat   bool    `json:"repeat,omitempty"`   // 3. 是否重复字段，即数组/片段 "rep"
	Name     string  `json:"name,omitempty"`     // 4. 名字 "name="
	Json     string  `json:"json,omitempty"`     // 5. json 名字，编码 JSON 数据时使用 "json="
	Version  Version `json:"version,omitempty"`  // 6. 版本，通常为 proto3，proto2
	Enum     string  `json:"enum,omitempty"`     // 7. 枚举 "enum="
}

const (
	tokenVarInt   = "varint"
	tokenBytes    = "bytes"
	tokenRepeat   = "rep"
	tokenOptional = "opt"
	tokenProto3   = "proto3"
	tokenProto2   = "proto2"

	namePrefix = "name="
	jsonPrefix = "json="
	enumPrefix = "enum="
)

// TagParse 解析 protobuf 标签 tag（为 reflect.StructTag 对象 Lookup 方法的返回值），如果没找到，返回 nil
func TagParse(tag string) (*Tag, error) {
	ts := strings.Split(tag, ",")
	if len(ts) < 1 {
		return nil, nil
	}
	r := &Tag{Version: VersionUndefined}
	for i, t := range ts {
		if i == 1 {
			if v, e := strconv.Atoi(t); e == nil {
				r.Index = v
				continue
			}
		}
		switch t {
		case tokenBytes:
			r.Type = Bytes
		case tokenVarInt:
			r.Type = VarInt
		case tokenOptional:
			r.Optional = true
		case tokenRepeat:
			r.Repeat = true
		case tokenProto3:
			r.Version = V3
		case tokenProto2:
			r.Version = V2
		default:
			if v, ok := strings.CutPrefix(t, namePrefix); ok {
				r.Name = v
			} else if v, ok = strings.CutPrefix(t, jsonPrefix); ok {
				r.Json = v
			} else if v, ok = strings.CutPrefix(t, enumPrefix); ok {
				r.Enum = v
			}
		}
	}
	return r, nil
}

func Parse(tag reflect.StructTag) (main, key, value *Tag, err error) {
	if t, ok := tag.Lookup(string(TagName)); ok {
		main, err = TagParse(t)
	}
	if err != nil {
		return
	}
	if t, ok := tag.Lookup(string(KeyTagName)); ok {
		key, err = TagParse(t)
	}
	if err != nil {
		return
	}
	if t, ok := tag.Lookup(string(ValTagName)); ok {
		value, err = TagParse(t)
	}
	return
}
