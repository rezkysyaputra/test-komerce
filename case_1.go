package main

import (
	"fmt"
	"strings"
)

func main() {
	ShortCharacters("Sample Case")
	ShortCharacters("Next Case")
}

func ShortCharacters(s string) {
	vowelCount := make(map[string]int)
	consonantCount := make(map[string]int)

	vowelOrder := []string{}
	consonantOrder := []string{}

	vowelChar := "aeiou"
	s = strings.ToLower(s)

	for _, v := range s {
		if v == ' ' {
			continue
		}

		if strings.Contains(vowelChar, string(v)) {
			if vowelCount[string(v)] == 0 {
				vowelOrder = append(vowelOrder, string(v))
			}
			vowelCount[string(v)]++
		} else {
			if consonantCount[string(v)] == 0 {
				consonantOrder = append(consonantOrder, string(v))
			}
			consonantCount[string(v)]++
		}
	}

	fmt.Print("Vowel Characters: ")
	for _, v := range vowelOrder {
		for i := 0; i < vowelCount[v]; i++ {
			print(v)
		}
	}
	fmt.Println()

	fmt.Print("Consonant Characters: ")
	for _, v := range consonantOrder {
		for i := 0; i < consonantCount[v]; i++ {
			print(v)
		}
	}
	fmt.Println()
}
