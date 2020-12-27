package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{"abcabcbb"}, 3},
		{"b", args{"bbbbb"}, 1},
		{"c", args{"pwwkew"}, 3},
		{"d", args{""}, 0},
		{"e", args{" "}, 1},
		{"f", args{"aab"}, 2},
		{"g", args{"abcdefcba"}, 6},
		{"h", args{"bcaefgahm"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.args.s)
			assert.Equal(t, tt.want, got, "Test: %s", tt.name)
		})
	}
}
