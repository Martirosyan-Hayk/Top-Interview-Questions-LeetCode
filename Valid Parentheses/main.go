package main

func isValid(s string) bool {
	index := 0
	var stack []rune
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			index++
		} else if ch == ')' || ch == ']' || ch == '}' {
			if index == 0 {
				return false
			}
			top := stack[index-1]
			if (ch == ')' && top != '(') ||
				(ch == ']' && top != '[') ||
				(ch == '}' && top != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
			index--
		}
	}

	return index == 0
}
