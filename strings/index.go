package strings

import "unicode/utf8"

// Unbracketed attempt to remove the brackets from the name.
//
//	Unbracketed("/regexp/", '/') // output: regexp
//	Unbracketed("<div>", '<', '>') // output: div
//
// 去括号，参数 brackets 为括号字符，提供一个字符时，此字符同时作为首尾括号；提供两个字符时，分别作为首尾括号，更多的字符则被忽略
func Unbracketed(name string, brackets ...rune) (string, bool) {
	var sb, eb rune
	bl := len(brackets)
	if bl < 1 {
		return name, false
	}
	sb = brackets[0]
	if bl > 1 {
		eb = brackets[1]
	} else {
		eb = sb
	}
	p, pw := utf8.DecodeRuneInString(name)
	if p == utf8.RuneError {
		return name, false
	}
	s, sw := utf8.DecodeLastRuneInString(name)
	if s == utf8.RuneError {
		return name, false
	}
	if p == sb && s == eb {
		w := len(name)
		return name[pw : w-sw], true
	}
	return name, false
}

// Unquote attempt to remove the quotes from the name.
//
// 尝试去除名字前后的任意字符（一般用于反撇号，单引号，双引号等），
// 当名字前后指定字符非成对时，视为不正确，ok 返回 false。
// 指定多个时依次尝试，但仅执行一次。
// found 返回找到的引号字符值，0 为未找到。
func Unquote(name string, quotes ...byte) (id string, ok bool, found byte) {
	if len(quotes) < 1 {
		return name, false, 0
	}
	l := len(name)
	if l < 2 {
		return name, false, 0
	}
	for _, quote := range quotes {
		if name[0] == quote && name[l-1] == quote {
			return name[1 : l-1], true, quote
		}
	}
	return name, false, 0
}

// UnquoteRune 尝试去除名字前后的任意字符对（一般用于反撇号，单引号，双引号等），
// 当名字前后指定字符非成对时，视为不正确，ok 返回 false。
// 指定多个时依次尝试，但仅执行一次。
// found 返回找到的引号字符值，utf8.RuneError 为未找到。
func UnquoteRune(name string, quotes ...rune) (id string, ok bool, found rune) {
	if len(quotes) < 1 {
		return name, false, utf8.RuneError
	}
	p, pw := utf8.DecodeRuneInString(name)
	if p == utf8.RuneError {
		return name, false, p
	}
	s, sw := utf8.DecodeLastRuneInString(name)
	if s == utf8.RuneError {
		return name, false, s
	}
	c := utf8.RuneCountInString(name)
	if c < pw+sw {
		return name, false, utf8.RuneError
	}
	for _, quote := range quotes {
		if p == quote && s == quote {
			// 找到引号并且前后一致
			w := len(name)
			return name[pw : w-sw], true, quote
		}
	}
	return name, false, utf8.RuneError
}

type CaseMode = int

const (
	Ignored        CaseMode = iota // ignore case style
	SnakeCaseMode                  // snake case style, 蛇形，形如：id_user, id_account
	UpperCaseMode                  // upper snake case style, 大蛇形，形如：ID_USER，ID_ACCOUNT
	CamelCaseMode                  // camel case style, 驼峰，形如：idUser，idAccount
	PascalCaseMode                 // upper camel case style, 大驼峰/Pascal，形如：IdUser，IdAccount
)

func Case(s string, mode CaseMode) string {
	switch mode {
	case CamelCaseMode:
		return CamelCase(s)
	case UpperCaseMode:
		return UpperCase(s)
	case PascalCaseMode:
		return PascalCase(s)
	case SnakeCaseMode:
		return SnakeCase(s)
	default:
	}
	return s
}

// SnakeCase 返回单词的蛇形样式（将大写字母作为单词分隔符，前加 _ 并转为小写，首字符除外），其它字符不作额外的处理，
// 例如 idUser => id_user，IdAccount => id_account
func SnakeCase(s string) string {
	var re []byte
	for i, n := range []byte(s) {
		if n >= 'A' && n <= 'Z' {
			if i > 0 {
				re = append(re, '_')
			}
			re = append(re, n+('a'-'A')) // a:97 - A:65 = 32
		} else {
			re = append(re, n)
		}
	}
	return string(re)
}

// UpperCase 返回单词的大蛇形样式（将大写字母作为单词分隔符，前加 _，首字符除外），
// 小写字母转为大写字母，其它字符不作处理，例如 idUser => ID_USER，IdAccount => ID_ACCOUNT，idUser3 => ID_USER3
func UpperCase(s string) string {
	var re []byte
	for i, n := range []byte(s) {
		if n >= 'A' && n <= 'Z' {
			if i > 0 {
				re = append(re, '_')
			}
			re = append(re, n)
		} else if n >= 'a' && n <= 'z' {
			re = append(re, n-('a'-'A')) // a:97 - A:65 = 32
		} else {
			re = append(re, n)
		}
	}
	return string(re)
}

// PascalCase returns the CamelCased name with the first character capitalized.
//
// 返回单词的大驼峰样式（将大写字母作为单词分隔符），例如 idUser => IdUser，IdAccount => IdAccount，idUser3 => IdUser3
func PascalCase(s string) string {
	if len(s) < 1 {
		return ""
	}
	t := make([]byte, 0, 32)
	// i := 0
	//if s[0] == '_' {
	//	// Need a capital letter; drop the '_'.
	//	t = append(t, 'X')
	//	i++
	//}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// CamelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercase names, it's extremely unlikely to have two fields
// with different capitalization.
// For example, _my_field_name_2 => MyFieldName2, idUser => idUser，
// IdAccount => idAccount，idUser3 => idUser3
//
// 返回字符串的小驼峰样式，例如 _my_field_name_2 => MyFieldName2, idUser => idUser，
// IdAccount => idAccount，idUser3 => idUser3
func CamelCase(s string) string {
	if len(s) < 1 {
		return s
	}
	s = PascalCase(s)
	b := []byte(s)
	if b[0] >= 'A' || b[0] <= 'Z' {
		b[0] = b[0] + ('a' - 'A')
	}
	return string(b)
}

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
