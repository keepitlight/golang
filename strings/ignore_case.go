package strings

import (
	"strings"
	"unicode/utf8"
	"unsafe"
)

// IgnoreCase is case-insensitive string.
type IgnoreCase string

// MemoryUsage return the memory usage of IgnoreCase
func (s IgnoreCase) MemoryUsage() (sum int64) {
	if s == "" {
		return
	}
	return int64(unsafe.Sizeof(s))*2 + int64(len(s))
}

func (s IgnoreCase) Count(substr string) int {
	return strings.Count(strings.ToLower(string(s)), strings.ToLower(substr))
}
func (s IgnoreCase) Equal(t string) bool {
	return strings.EqualFold(string(s), t)
}
func (s IgnoreCase) ToLower() string {
	return strings.ToLower(string(s))
}
func (s IgnoreCase) ToUpper() string {
	return strings.ToUpper(string(s))
}

func (s IgnoreCase) HasPrefix(prefix string) bool {
	return strings.HasPrefix(strings.ToLower(string(s)), strings.ToLower(prefix))
}
func (s IgnoreCase) HasSuffix(suffix string) bool {
	return strings.HasSuffix(strings.ToLower(string(s)), strings.ToLower(suffix))
}
func (s IgnoreCase) Contains(substring string) bool {
	return strings.Contains(strings.ToLower(string(s)), strings.ToLower(substring))
}
func (s IgnoreCase) ContainsAny(chars string) bool {
	return strings.ContainsAny(strings.ToLower(string(s)), strings.ToLower(chars))
}
func (s IgnoreCase) ContainsRune(r rune) bool {
	return strings.ContainsRune(strings.ToLower(string(s)), r)
}
func (s IgnoreCase) Index(substring string) int {
	return strings.Index(strings.ToLower(string(s)), strings.ToLower(substring))
}
func (s IgnoreCase) IndexAny(chars string) int {
	return strings.IndexAny(strings.ToLower(string(s)), strings.ToLower(chars))
}
func (s IgnoreCase) IndexRune(r rune) int {
	return strings.IndexRune(strings.ToLower(string(s)), r)
}
func (s IgnoreCase) LastIndex(substring string) int {
	return strings.LastIndex(strings.ToLower(string(s)), strings.ToLower(substring))
}
func (s IgnoreCase) LastIndexAny(chars string) int {
	return strings.LastIndexAny(strings.ToLower(string(s)), strings.ToLower(chars))
}
func (s IgnoreCase) LastIndexRune(r rune) int {
	return strings.LastIndex(strings.ToLower(string(s)), string(r))
}
func (s IgnoreCase) Replace(old, new string, count int) string {
	z := string(s)

	old = strings.ToLower(old) // search ignoring case
	if old == new || count == 0 {
		return z // avoid allocation
	}

	dummy := strings.ToLower(z)

	// Compute number of replacements.
	if m := strings.Count(dummy, old); m == 0 {
		return z // avoid allocation
	} else if count < 0 || m < count {
		count = m
	}

	// Apply replacements to buffer.
	var b strings.Builder
	b.Grow(len(s) + count*(len(new)-len(old)))
	start := 0
	for i := 0; i < count; i++ {
		j := start
		if len(old) == 0 {
			if i > 0 {
				_, wid := utf8.DecodeRuneInString(dummy[start:])
				j += wid
			}
		} else {
			j += strings.Index(dummy[start:], old)
		}
		b.WriteString(z[start:j])
		b.WriteString(new)
		start = j + len(old)
	}
	b.WriteString(z[start:])
	return b.String()
}

func (s IgnoreCase) ReplaceAll(old, new string) string {
	return s.Replace(old, new, -1)
}

func explode(s string, n int) []string {
	l := utf8.RuneCountInString(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]string, n)
	for i := 0; i < n-1; i++ {
		_, size := utf8.DecodeRuneInString(s)
		a[i] = s[:size]
		s = s[size:]
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}

func genSplit(s, sep string, sepSave, n int) []string {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return explode(s, n)
	}
	sep = strings.ToLower(sep)
	dummy := strings.ToLower(s)
	if n < 0 {
		n = strings.Count(dummy, sep) + 1
	}

	if n > len(s)+1 {
		n = len(s) + 1
	}
	a := make([]string, n)
	n--
	i := 0
	for i < n {
		m := strings.Index(dummy, sep)
		if m < 0 {
			break
		}
		a[i] = s[:m+sepSave]
		c := m + len(sep)
		s = s[c:]
		dummy = dummy[c:]
		i++
	}
	a[i] = s
	return a[:i+1]
}

func (s IgnoreCase) Split(sep string) []string {
	return genSplit(string(s), sep, 0, -1)
}
func (s IgnoreCase) SplitN(sep string, n int) []string {
	return genSplit(string(s), sep, 0, n)
}
func (s IgnoreCase) SplitAfterN(sep string, n int) []string {
	return genSplit(string(s), sep, len(sep), n)
}
func (s IgnoreCase) SplitAfter(sep string) []string {
	return genSplit(string(s), sep, len(sep), -1)
}

func (s IgnoreCase) Cut(sep string) (before, after string, found bool) {
	z := string(s)
	if i := s.Index(sep); i >= 0 {
		return z[:i], z[i+len(sep):], true
	}
	return z, "", false
}
func (s IgnoreCase) CutPrefix(prefix string) (after string, found bool) {
	z := string(s)
	if strings.HasPrefix(strings.ToLower(z), strings.ToLower(prefix)) {
		return z[len(prefix):], true
	}
	return z, false
}
func (s IgnoreCase) CutSuffix(suffix string) (before string, found bool) {
	z := string(s)
	if strings.HasSuffix(strings.ToLower(z), strings.ToLower(suffix)) {
		return z[:len(z)-len(suffix)], true
	}
	return z, false
}

func (s IgnoreCase) TrimPrefix(prefix string) string {
	z := string(s)
	if strings.HasPrefix(strings.ToLower(z), strings.ToLower(prefix)) {
		return z[len(prefix):]
	}
	return z
}
func (s IgnoreCase) TrimSuffix(suffix string) string {
	z := string(s)
	if strings.HasSuffix(strings.ToLower(z), strings.ToLower(suffix)) {
		return z[:len(z)-len(suffix)]
	}
	return z
}
