package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"a", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"b", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"c", args{[]int{3, 2, 4}, 5}, []int{0, 1}},
		{"d", args{[]int{0, 8, 7, 3, 3, 4, 2}, 11}, []int{1, 3}},
		{"e", args{[]int{0, 1}, 1}, []int{0, 1}},
		{"f", args{[]int{0, 3}, 5}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
