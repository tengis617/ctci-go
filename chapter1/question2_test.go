package chapter1

import (
	"fmt"
	"testing"
)

func Test_isPermutation(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"abcd", "abcd", true},
		{"abaa", "baaa", true},
		{"abaa", "baba", false},
		{"a", "b", false},
		{"a", "a", true},
		{"aaaaa", "aaa", false},
		{"a ", " a", true},
		{"Abc", "abc", false}, // we assume its case sensitive!
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("a: %s b: %s", tt.a, tt.b), func(t *testing.T) {
			if got := isPermutation(tt.a, tt.b); got != tt.want {
				t.Errorf("isPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
