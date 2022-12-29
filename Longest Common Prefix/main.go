package main

func longestCommonPrefix(strs []string) string {
	strsLength := len(strs)

	if strsLength == 0 {
		return ""
	}

	if strsLength == 1 {
		return strs[0]
	}

	prefix := ""

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]

		for j := 1; j < strsLength; j++ {

			if i >= len(strs[j]) || strs[j][i] != c {
				return prefix
			}
		}

		prefix += string(c)
	}

	return prefix
}
