package strings

import "testing"

func TestUnbracketed(t *testing.T) {
	for _, tt := range []struct {
		input  string
		be     rune
		se     rune
		result string
		ok     bool
	}{
		{"(content)", '(', ')', "content", true},
		{"[content]", '[', ']', "content", true},
		{"{content}", '{', '}', "content", true},
		{"<content>", '<', '>', "content", true},
		{"【content】", '【', '】', "content", true},
		{"「中国的英文China」", '「', '」', "中国的英文China", true},
	} {
		if r, ok := Unbracketed(tt.input, tt.be, tt.se); r != tt.result || ok != tt.ok {
			t.Errorf("Unbracketed(%q) = %q, %t; want %q, %t", tt.input, r, ok, tt.result, tt.ok)
		}
	}
}
