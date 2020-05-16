package chapter1

import (
	"reflect"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	tests := []struct {
		name string
		m    [][]int
		want [][]int
	}{
		{
			name: "2x2",
			m: [][]int{
				{1, 2},
				{4, 3},
			},
			want: [][]int{
				{4, 1},
				{3, 2},
			},
		},
		{
			name: "3x3",
			m: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			want: [][]int{
				{7, 8, 1},
				{6, 9, 2},
				{5, 4, 3},
			},
		},
		{
			name: "4x4",
			m: [][]int{
				{10, 11, 12, 13},
				{21, 22, 23, 14},
				{20, 25, 24, 15},
				{19, 18, 17, 16},
			},
			want: [][]int{
				{19, 20, 21, 10},
				{18, 25, 22, 11},
				{17, 24, 23, 12},
				{16, 15, 14, 13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateMatrix(tt.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
