package chapter1

import (
	"fmt"
	"testing"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"aaa", "a3"},
		{"aabcccccaaa", "a2b1c5a3"},
		{"aaron", "aaron"},
		{"aaaaron", "aaaaron"},
		{"aaaaaron", "aaaaaron"},  // a4r1o1n1 is same as original string
		{"aaaaaaron", "a6r1o1n1"}, // a6r1o1n1 is shorter than original
		{"aa", "aa"},
		{"abcdefgh", "abcdefgh"},
		{"a", "a"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("s: %s", tt.s), func(t *testing.T) {
			if got := compress(tt.s); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCompressionLen(t *testing.T) {
	tests := []struct {
		s       string
		wantLen int
	}{
		{"a", 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("s: %s", tt.s), func(t *testing.T) {
			if gotLen := getCompressionLen(tt.s); gotLen != tt.wantLen {
				t.Errorf("getCompressionLen() = %v, want %v", gotLen, tt.wantLen)
			}
		})
	}
}
