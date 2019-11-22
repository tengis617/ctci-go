package main

import (
	"fmt"
	"testing"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"a", true},
		{"ab", true},
		{"abb", false},
		{"abcde", true},
		{"01234", true},
		{"011", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d - %s", i, tt.s), func(t *testing.T) {
			got := isUniqueHashMap(tt.s)

			if got != tt.want {
				t.Errorf("got %t want %t", got, tt.want)
			}
		})
	}
}

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"a", "b", false},
		{"a", "bc", false},
		{"ab", "ba", true},
		{"abcd", "baaa", false},
		{"ab", "aaaabaaa", false},
		{"aabb", "abbb", false},
		{"Ab", "ba", false},
		{"space    ", "s p a c e", true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d - %s and %s", i, tt.a, tt.b), func(t *testing.T) {
			got := isPermutation(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %t want %t", got, tt.want)
			}
		})
	}
}

func TestURLify(t *testing.T) {
	tests := []struct {
		s    string
		l    int
		want string
	}{
		{"a b  ", 3, "a%20b"},
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d - %s with len %d", i, tt.s, tt.l), func(t *testing.T) {
			got := URLify(tt.s, tt.l)
			if got != tt.want {
				t.Errorf("got %s want %s", got, tt.want)
			}
		})
	}
}
