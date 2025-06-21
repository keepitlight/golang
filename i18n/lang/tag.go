package lang

import (
	"slices"
)

// Separators is the valid language code separator. "-" is the default separator and always works.
//
// 指定有效的语言代码分隔符，减号是默认的分隔符，并且总是有效的
var Separators = []rune{
	'_',
}

// Tag represents a language code of ISO639.
// 代表语言代码
type Tag string

const (
	// Unknown is the unknown language code.
	//
	// 未知语言代码
	Unknown Tag = ""

	// EN is the default language code.
	//
	// 默认语言代码
	EN Tag = "en"
	// ZH is the Chinese language code.
	//
	// 中文语言代码
	ZH Tag = "zh"

	US Tag = "en-us" // 英语（美国）
	CN Tag = "zh-cn" // 中文（中国）
)

type char int

const (
	unknown   char = iota // 未知类型
	lower                 // 小写字母类型
	upper                 // 大写字母类型
	digit                 // 数字类型
	separator             // 分隔符类型
)

func kindOf(c rune) char {
	if c >= 'a' && c <= 'z' {
		return lower // 字母
	}
	if c >= 'A' && c <= 'Z' {
		return upper
	}
	if c >= '0' && c <= '9' {
		return digit // 数字
	}
	if c == '-' || slices.Contains(Separators, c) {
		return separator // 分隔符
	}
	return unknown // 未知类型
}

// Ensure to validate the syntax of this tag in ABNF [RFC 2234]
// Language-Tag = Primary-subtag *( "-" Subtag )
// Primary-subtag = 1*8ALPHA
//
// Subtag = 1*8(ALPHA / DIGIT)
//
// The tag must be at least 1 character long and can contain letters, digits, and separators.
//
// 确保语言代码的语法符合 ABNF [RFC 2234] 的规范
func Ensure(input string) (tag Tag, ok bool) {
	if len(input) < 1 {
		return Unknown, false
	}

	v := []rune(input)
	l := len(v)
	var p char
	var w byte

	for i, c := range v {
		t := kindOf(c)
		switch t {
		case separator:
			// 分隔符出现在首位，或者最后，或者多个连续的分隔符
			if i == 0 || i == l-1 || p == separator {
				return
			}
			v[i] = '-' // 转换为统一的分隔符
		case upper:
			v[i] = c + 32 // 转换为小写字母
		case digit:
			// 数字不能出现在首位
			if i == 0 {
				return
			}
		case lower:
		default:
			return
		}
		if t == separator || i == l-1 {
			if w > 8 {
				return // 分隔符前的子标签长度不能超过8
			}
			w = 0
		} else {
			w++
		}
		p = t
	}
	return Tag(v), true
}

// SubTags splits the input string into tags based on the defined separators.
// The tags must be at least 1 character long and can contain letters, digits, and separators.
// The function returns a lower slice of tags and a boolean indicating success.
//
// 将输入字符串按定义的分隔符拆分为标签，标签必须至少1个字符长，并且可以包含字母、数字和分隔符。
// 函数返回一个小写字母版本的标签切片和一个布尔值，指示是否成功
func SubTags(input string) (tags []string, ok bool) {
	if len(input) < 1 {
		return
	}
	v := []rune(input)
	l := len(v)

	var p char
	var s int

	for i, c := range v {
		t := kindOf(c)
		switch t {
		case separator:
			// 分隔符出现在首位，或者最后，或者多个连续的分隔符
			if i == 0 || i == l-1 || p == separator {
				return
			}
		case upper:
			v[i] = c + 32 // 转换为小写字母
		case digit:
			// 数字不能出现在首位
			if i == 0 {
				return
			}
		case lower:
		default:
			return
		}
		if t == separator || i == l-1 {
			// 遇到分隔符或者到达字符串末尾
			e := i
			if t == separator {
				e--
			}
			var tag = v[s : e+1] // 截取子标签
			if t == separator {
				s = i + 1 // 更新起始位置
			}
			if len(tag) > 8 {
				return // 分隔符前的子标签长度不能超过8
			}
			tags = append(tags, string(tag))
		}
		p = t
	}
	return tags, true
}

func (t Tag) Primary() Tag {
	tags, ok := t.SubTags()
	if !ok || len(tags) < 1 {
		return Unknown
	}
	return Tag(tags[0])
}

func (t Tag) SubTags() (tags []string, ok bool) {
	return SubTags((string)(t))
}

type Scope int

const (
	// Individual represents an individual language.
	//个体语言
	Individual Scope = iota
	// Macro represents a macro languages 宏语言
	Macro
)

type Type int

const (
	Undefined   Type = iota // 未定义的语言类型
	Living                  // 活跃使用的语言
	Historical              // 历史中存在过的语言
	Constructed             // 人工构造的语言
)

// Metadata represents the metadata of a language tag.
//
// 元数据代表语言标签的元数据
type Metadata struct {
	Type     Type     `json:"type,omitempty"`     // Type is the type of the language, such as living, historical, or constructed
	Scope    Scope    `json:"scope,omitempty"`    // Scope is the scope of the language, such as individual or macro
	Set1     string   `json:"set1,omitempty"`     // Set1 is the ISO 639-1 code, if available
	Set2     string   `json:"set2,omitempty"`     // Set2 is the ISO 639-2 code, if available
	Set3     string   `json:"set3,omitempty"`     // Set3 is the ISO 639-3 code, if available
	Names    []string `json:"names,omitempty"`    // Names is the name of the ISO language in English
	Endonyms []string `json:"endonyms,omitempty"` // Endonyms is the name of the language in its own language
}

// MetadataProvider is an interface for looking up language metadata by tag.
//
// 提供者是一个接口，用于通过标签查找语言元数据
type MetadataProvider interface {
	Lookup(tag Tag) (Metadata, bool)
}
