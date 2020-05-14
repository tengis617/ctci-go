package chapter1

import "testing"

func Test_isUnique(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"a", true},
		{"ab", true},
		{"abcd", true},
		{"abcda", false},
		{"aa", false},
		{"whatcouldgowrong", false},
		{` !"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_abcdefghijklmnopqrstuvwxyz{|}~`, true},
		{` !"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_abcdefghijklmnopqrstuvwxyz{|}~  !"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_abcdefghijklmnopqrstuvwxyz{|}~`, false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isUniqueHashTable(tt.s); got != tt.want {
				t.Errorf("isUniqueHashTable() = %v, want %v", got, tt.want)
			}
			if got := isUniqueLoop(tt.s); got != tt.want {
				t.Errorf("isUniqueLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
