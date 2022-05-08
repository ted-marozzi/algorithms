package leetcode

import "testing"

func TestLengthOfLongestSubString(t *testing.T) {

	if LengthOfLongestSubstring("abcabcbb") != 3 {
		t.Error("1)")
	}

	if LengthOfLongestSubstring("bbbbb") != 1 {
		t.Error("2)")
	}

	if LengthOfLongestSubstring("pwwkew") != 3 {
		t.Error("3)")
	}
}
