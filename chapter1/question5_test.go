package chapter1

import (
	"fmt"
	"testing"
)

func Test_isOneEditAway(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"a", "a", true},
		{"ab", "ab", true},
		{"ab", "aa", true},
		{"snail", "smail", true},
		{"rail", "mail", true},
		{"gmail", "mail", true},
		{"google", "googlemaps", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("a: %s b: %s", tt.a, tt.b), func(t *testing.T) {
			if got := isOneEditAway(tt.a, tt.b); got != tt.want {
				t.Errorf("isOneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
