package cn

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{']': '[', '}': '{', ')': '('}
	for _, c := range s {
		switch c {
		case '[', '{', '(':
			stack = append(stack, c)
		case ']', '}', ')':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != m[c] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidParentheses(t *testing.T) {
	assert.Equal(t, false, isValid("[{]}"))
	assert.Equal(t, true, isValid("[]{}"))
	assert.Equal(t, false, isValid("[}"))
	assert.Equal(t, true, isValid("({[]})"))
	assert.Equal(t, false, isValid(")"))
}

func isValid_1(s string) bool {
	l := list.New()
	for _, i2 := range s {
		if string(i2) == "[" || string(i2) == "{" || string(i2) == "(" {
			l.PushBack(string(i2))
			continue
		}
		e := l.Back()
		if e == nil {
			return false
		}
		ls := e.Value.(string)
		if string(i2) == "]" {
			if ls == "[" {
				l.Remove(l.Back())
			} else {
				return false
			}
		} else if string(i2) == "}" {
			if ls == "{" {
				l.Remove(l.Back())
			} else {
				return false
			}
		} else if string(i2) == ")" {
			if ls == "(" {
				l.Remove(l.Back())
			} else {
				return false
			}
		}
	}
	return l.Len() == 0
}
