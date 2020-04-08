package cn

import (
	"sort"
)

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例: 
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//] 
//
// 说明： 
//
// 
// 所有输入均为小写字母。 
// 不考虑答案输出的顺序。 
// 
// Related Topics 哈希表 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	// 哈希表 key=字母数组 value=结果索引
	groups := make(map[[26]byte]int)
	// 返回的结果
	ret := make([][]string, 0)
	// 遍历输入
	for _, str := range strs {
		// 创建当前单词的哈希表的key
		var key [26]byte
		// 遍历单词每个字母出现的次数
		for _, c := range str {
			key[c-'a']++
		}
		// 查哈希表是否已经存在改key
		if i, ok := groups[key]; ok {
			ret[i] = append(ret[i], str)
		} else {
			// 新建一个字符串数组
			ret = append(ret, []string{str})
			groups[key] = len(ret) - 1
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func groupAnagrams_1(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		s := string(bs)
		groups[s] = append(groups[s], str)
	}
	ret := make([][]string, 0, len(groups))
	for _, str := range groups {
		ret = append(ret, str)
	}
	return ret
}
