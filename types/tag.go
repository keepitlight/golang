package types

import (
	"errors"
	"reflect"
	"strings"
)

// TagName represents the name of the struct field tag.
//
// 结构字段标签名称
type TagName string

// TagParseFunc is a tag parser.
//
// 结构字段标签解析器
type TagParseFunc[T any] func(tag string) (T, error)

type wrapper struct {
	parse any
	wrap  func(string) (any, error)
}

var (
	tagParsers = map[TagName]wrapper{
		JSON: wrapper{JsonTagParse, func(tag string) (any, error) {
			return JsonTagParse(tag)
		}},
	}
)

// RegisterTagParser registers a tag parser.
//
// 注册标签解析器
func RegisterTagParser[T any](tag TagName, parser TagParseFunc[T]) {
	tagParsers[tag] = wrapper{
		parser,
		func(tag string) (any, error) {
			return parser(tag)
		},
	}
}

// Tags represents the tags of the struct.
//
// 结构字段标签集合
type Tags map[TagName]any

// Tag returns the structured tag data from the tag set, note that the tag name matches the struct,
// and the non-matching return nil
//
// 从标签集中获取结构化的标签数据，注意标签名与结构体的匹配，不匹配返回 nil
func Tag[T any](tags Tags, name TagName) (result T, ok bool) {
	if s, f := tags[name]; f {
		if v, o := s.(T); o {
			return v, true
		}
	}
	return
}

// ParseTags returns all registered struct field tags, and uses the Tag method to extract the struct data
//
// 解析所有已注册的结构字段标签，使用 Tag 方法提取结构数据
func ParseTags(tag reflect.StructTag) (tags Tags, err error) {
	for k, w := range tagParsers {
		if s, ok := tag.Lookup(string(k)); ok {
			if v, e := w.wrap(s); e != nil {
				return nil, e
			} else {
				if tags == nil {
					tags = map[TagName]any{}
				}
				tags[k] = v
			}
		}
	}
	return
}

var (
	ErrorTagUndefined     = errors.New("tag undefined")
	ErrorTagNotFound      = errors.New("tag not found")
	ErrorTagParserInvalid = errors.New("tag parser type invalid")
)

// ParseTag returns the structured tag data from the tag set, note that the tag name matches the struct,
//
// 解析结构字段标签
func ParseTag[T any](tag reflect.StructTag, name TagName) (result T, err error) {
	if w, f := tagParsers[name]; f {
		var parser TagParseFunc[T]
		if c, ok := w.parse.(TagParseFunc[T]); !ok {
			err = ErrorTagParserInvalid
			return
		} else {
			parser = c
		}
		if s, ok := tag.Lookup(string(name)); ok {
			if t, e := parser(s); e != nil {
				err = e
			} else {
				result = t
			}
		} else {
			err = ErrorTagNotFound
		}
	} else {
		err = ErrorTagUndefined
	}
	return
}

// JsonTag represents the json tag.
//
// 结构体 JSON 标签
type JsonTag struct {
	Name      string `json:"name,omitempty"`      // json 字段名
	OmitEmpty bool   `json:"omitEmpty,omitempty"` // 是否忽略空值
}

const (
	JSON TagName = "json" // json 标签
)

// JsonTagParse returns the json tag data as a JsonTag struct
//
// 解析 json 标签为 JsonTag 结构，返回值为 JsonTag 结构
func JsonTagParse(tag string) (*JsonTag, error) {
	if n, o, f := strings.Cut(tag, ","); f {
		return &JsonTag{Name: n, OmitEmpty: o == "omitempty"}, nil
	}
	return nil, nil
}
