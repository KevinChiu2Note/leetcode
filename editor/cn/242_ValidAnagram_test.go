package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1: 
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
// 
//
// 示例 2: 
//
// 输入: s = "rat", t = "car"
//输出: false 
//
// 说明: 
//你可以假设字符串只包含小写字母。 
//
// 进阶: 
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？ 
// Related Topics 排序 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
// 使用map
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := [26]int{}
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
		table[t[i]-'a']--
	}
	for _, v := range table {
		if v != 0 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidAnagram(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))
	assert.Equal(t, false, isAnagram("anagrad", "nagaram"))
	assert.Equal(t, false, isAnagram("ana", "nagaram"))
}

func isAnagram_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range t {
		if _, ok := m[v]; ok {
			m[v]--
			if m[v] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func isAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := [26]int{}
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		table[t[i]-'a']--
		if table[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
