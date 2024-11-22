package strings

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

// MD5 is a shortcut of md5.Sum, it'll return a hex string
//
// 返回一组字符串联接以后的 MD5 值，将返回 32 位的十六进制字符串
func MD5(vs ...string) string {
	hash := md5.New()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha1 is a shortcut of sha1.Sum, it'll return a hex string
//
// 返回一组字符串联接以后的字符串的 SHA1 值，将返回 32 位的十六进制字符串
func Sha1(vs ...string) string {
	hash := sha1.New()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA224 is a shortcut of sha256.Sum224
//
// 返回一组字符串联接以后的 SHA224 值
func SHA224(vs ...string) string {
	hash := sha256.New224()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// SHA256 is a shortcut of sha256.Sum256
//
// 返回一组字符串联接以后的 SHA256 值
func SHA256(vs ...string) string {
	hash := sha256.New()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha512 is a shortcut of sha512.Sum
//
// 返回一组字符串联接以后的 SHA512 值
func Sha512(vs ...string) string {
	hash := sha512.New()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha384 is a shortcut of sha512.Sum384
//
// 返回一组字符串联接以后的 sha512.Sum384 值
func Sha384(vs ...string) string {
	hash := sha512.New384()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha224 is a shortcut of sha512.Sum512_224
//
// 返回一组字符串联接以后的 sha512.Sum512_224 值, 224 位
func Sha224(vs ...string) string {
	hash := sha512.New512_224()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha256 is a shortcut of sha512.Sum512_256
//
// 返回一组字符串联接以后的 sha512.Sum512_256 值, 256 位
func Sha256(vs ...string) string {
	hash := sha512.New512_256()
	for _, v := range vs {
		_, _ = io.WriteString(hash, v)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
