package leetcode

/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func LengthOfLongestSubstring(s string) int {
	var result int
	for i := range s {
		var longestForChar int
		currentMap := make(map[rune]struct{})
		for _, charTwo := range s[i:] {
			if _, ok := currentMap[charTwo]; ok {
				// This is a repeat letter so break
				break
			}
			currentMap[charTwo] = struct{}{}
			longestForChar++
		}
		if longestForChar > result {
			result = longestForChar
		}
	}

	return result
}
