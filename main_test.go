package main

import (
	"testing"
)

var flagtests = []struct {
	in  string
	out string
}{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"45", ""},
	{`qwe\4\5`, "qwe45"},
	{`qwe\45`, "qwe44444"},
	{`qwe\\5`, `qwe\\\\\`},
}

func TestStringConvert(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := stringConvert(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}

}
