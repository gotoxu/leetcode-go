package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"a", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		// {"b", args{[]int{1, 3}, []int{2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2)
			assert.Equal(t, float64(tt.want), got)
		})
	}
}
