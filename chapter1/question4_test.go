package chapter1

import "testing"

func Test_isPalindromePermutation(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"Tact Coa", true},
		{"Tactic Coaic", true},
		{"aa bb cc", true},
		{"aba aba", true},
		{"aabbc", true},
		{"ab", false},
		{"ab ab", true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isPalindromePermutation(tt.s); got != tt.want {
				t.Errorf("isPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
