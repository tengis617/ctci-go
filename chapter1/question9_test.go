package chapter1

import (
	"fmt"
	"testing"
)

func Test_isSubstring(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"abcd", "cdab", true},
		{"waterbottle", "erbottlewat", true},
		{"fitnesstrainer", "trainerfitness", true},
		{"worker", "erwork", true},
		{"abcd", "ab", false},
		{"a", "b", false},
		{"astring", "", false},
		{"", "s", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("s1: %s, s2: %s", tt.s1, tt.s2), func(t *testing.T) {
			if got := isSubstring(tt.s1, tt.s2); got != tt.want {
				t.Errorf("isSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
