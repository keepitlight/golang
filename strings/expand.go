package strings

import "strings"

// Expand 将分隔符和定界符组织的字符串转换为 map[string]string，注意，未考虑字符转义的情况。
func Expand(tags, separator, delimiter string, cast func(string) string) map[string]string {
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
