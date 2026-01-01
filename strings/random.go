package strings

import (
	"crypto/rand"
	mr "math/rand/v2"
)

const (
	// printable characters with spaces 包含空格的可打印字符
	printable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ "
	// printable characters without spaces 不包含空格的可打印字符
	graphical = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	// printable characters without quotes 不包含引号的可打印字符
	noQuote = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#$%&()*+,-./:;<=>?@[\\]^_{|}~"

	// digits 数字
	digits = "0123456789"
	// alphabet and digits 数字和字母
	alphabetNumeric = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// all alphabet
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// only lowercase alphabet 小写字母
	alphabetLowercase = "abcdefghijklmnopqrstuvwxyz"
	// only uppercase alphabet 大写字母
	alphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// hexadecimal 16进制字符
	hexChar  = "0123456789abcdefABCDEF"
	hexLower = "0123456789abcdef"
	hexUpper = "0123456789ABCDEF"
	// base64 64进制字符
	base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

type RandomStringType int

const (
	Printable RandomStringType = iota
	Graphical
	NoQuote
	Digits
	AlphabetNumeric
	Alphabet
	AlphabetLowercase
	AlphabetUppercase
	Hex
	HexLower
	HexUpper
	Base64
)

var (
	// 映射表，元素均为 256 个字节
	mappings = map[RandomStringType]*[]byte{}
)

func getChars(t RandomStringType) string {
	switch t {
	case Printable:
		return printable
	case NoQuote:
		return noQuote
	case Digits:
		return digits
	case AlphabetNumeric:
		return alphabetNumeric
	case Alphabet:
		return alphabet
	case AlphabetLowercase:
		return alphabetLowercase
	case AlphabetUppercase:
		return alphabetUppercase
	case Hex:
		return hexChar
	case HexLower:
		return hexLower
	case HexUpper:
		return hexUpper
	case Base64:
		return base64
	default:
		return graphical
	}
}
func genMapping(chars string) []byte {
	cs := []byte(chars)
	// 随机打乱字符表
	mr.Shuffle(len(cs), func(i, j int) { cs[i], cs[j] = cs[j], cs[i] })
	result := make([]byte, 256)
	for i := 0; i < 256; i++ {
		// 循环填充
		result[i] = cs[i%len(cs)]
	}
	return result
}

func init() {
	// 生成映射表
	for i := Printable; i <= Base64; i++ {
		cs := genMapping(getChars(i))
		mappings[i] = &cs
	}
}

// RegisterRandomStringType to register a random string type
//
// 注册一个随机字符串类型
func RegisterRandomStringType(t RandomStringType, chars string) bool {
	if _, ok := mappings[t]; ok {
		return false
	}
	cs := genMapping(chars)
	mappings[t] = &cs
	return true
}

// RandomBytes to generate a random bytes, length < 0 returns nil
//
// 生成指定位数的随机字节，length < 0 时返回 nil
func RandomBytes(length int) []byte {
	if length <= 0 {
		return nil
	}
	result := make([]byte, length)
	_, _ = rand.Read(result)
	return result
}

// RandomString to generate a random string, length < 0 returns nil,
// notice only works for ASCII string
//
// 生成指定位数的安全随机字符串，length < 0 时返回 nil，注意仅适用于生成 ASCII 字符串
func RandomString(t RandomStringType, length int) []byte {
	if length <= 0 {
		return nil
	}

	var cs *[]byte
	if v, ok := mappings[t]; ok {
		cs = v
	} else {
		return nil
	}

	result := make([]byte, length)

	_, _ = rand.Read(result)

	for i := 0; i < len(result); i++ {
		result[i] = (*cs)[result[i]]
	}

	return result
}
