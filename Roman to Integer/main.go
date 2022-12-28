package main

import "fmt"

func romanToInt(s string) int {
    romanNumbers := make(map[rune]int)

	romanNumbers['I'] = 1;
	romanNumbers['V'] = 5;
	romanNumbers['X'] = 10;
	romanNumbers['L'] = 50;
	romanNumbers['C'] = 100;
	romanNumbers['D'] = 500;
	romanNumbers['M'] = 1000;

	var number int = 0;

	length := len(s)

	for i := 0; i < length; i++ {
		if i + 1 < length && romanNumbers[rune(s[i])] < romanNumbers[rune(s[i + 1])] {
			number -= romanNumbers[rune(s[i])];
		} else {
			number += romanNumbers[rune(s[i])];
		}
	}

	return number;
}