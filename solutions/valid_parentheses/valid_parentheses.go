package leetcode

func IsValid(s string) bool {
	brackets := map[rune]rune{'[': ']', '(': ')', '{': '}'}

	list := []rune(s)
	stack := make([]rune, 0, len(list))

	for _, val := range list {
		if _, ok := brackets[val]; ok {
			stack = append(stack, val)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		last := stack[len(stack) - 1]
		if brackets[last] == val {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
