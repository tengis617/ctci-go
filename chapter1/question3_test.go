package chapter1

import "testing"

func Test_urlify(t *testing.T) {
	tests := []struct {
		s    string
		l    int
		want string
	}{
		{"hello world  ", 11, "hello%20world"},
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"M", 1, "M"},
		{"M M  ", 3, "M%20M"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := urlify(tt.s, tt.l); got != tt.want {
				t.Errorf("urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}
