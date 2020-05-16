package chapter1

import (
	"reflect"
	"testing"
)

func Test_zeroMatrix(t *testing.T) {
	tests := []struct {
		name string
		m    [][]int
		want [][]int
	}{
		{
			name: "2x2",
			m: [][]int{
				{1, 0},
				{1, 1},
			},
			want: [][]int{
				{0, 0},
				{1, 0},
			},
		},
		{
			name: "3x3",
			m: [][]int{
				{1, 1, 0},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
		},
		{
			name: "3x4",
			m: [][]int{
				{1, 1, 0, 0},
				{1, 1, 1, 0},
				{1, 1, 1, 0},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name: "3x4",
			m: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 1},
				{0, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroMatrix(tt.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
