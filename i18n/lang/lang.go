package lang

import (
	"strings"
)

// Separators is the valid language code separator. "-" is the default separator and always works.
//
// 指定有效的语言代码分隔符，减号是默认的分隔符，并且总是有效的
var Separators = []byte{
	'_',
	'.',
	' ',
}

// Tag represents a language code of ISO639.
// 代表语言代码
type Tag string

const (
	// Unknown is the unknown language code.
	//
	// 未知语言代码
	Unknown Tag = ""

	// En is the default language code.
	//
	// 默认语言代码
	En Tag = "en"
	// Zh is the Chinese language code.
	//
	// 中文语言代码
	Zh Tag = "zh"

	EnUS Tag = "en-US" // 英语（美国）
	ZhCN Tag = "zh-CN" // 中文（中国）
)

func (t Tag) String() string {
	return strings.ToLower(string(t))
}

func (t Tag) Primary() Tag {
	parts := t.Tags()
	if len(parts) > 0 {
		return Tag(parts[0])
	}
	return Unknown
}

func (t Tag) Tags() (tags []string) {
	if len(t) < 1 {
		return
	}
	v := strings.ToLower(string(t))
	for _, sep := range Separators {
		if sep != '-' {
			v = strings.ReplaceAll(v, (string)([]byte{sep}), "-")
		}
	}
	return strings.Split(v, "-")
}
