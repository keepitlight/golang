package strings

import "strings"

// Unfold to split the tags string into a map[string]string using "separator" and "delimiter" as delimiters
//
// 使用 separator 和 delimiter 作为分隔符对 tags 字符串进行两次拆分，转换为 map[string]string，注意，未考虑字符转义的情况。
func Unfold(tags, separator, delimiter string, cast func(string) string) map[string]string {
	re := map[string]string{}
	for _, tag := range strings.Split(tags, separator) {
		nv := strings.SplitN(tag, delimiter, 2)
		var k, v string
		if len(nv) > 0 {
			k = strings.TrimSpace(nv[0])
			if cast != nil {
				k = cast(k)
			}
		}
		if len(nv) > 1 {
			v = nv[1]
		}
		if len(k) > 0 {
			re[k] = v
		}
	}
	return re
}
