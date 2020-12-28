package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
