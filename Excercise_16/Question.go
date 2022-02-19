package main

import (
	"fmt"
	"strings"
)

func wordsReverse(s string) string {
	// Slpit the string
	s1 := strings.Fields(s)

	reverse(&s1, 0, len(s1)-1)

	return strings.Join(s1, " ")
}

func reverse(s2 *[]string, i int, j int) {
	for i <= j {
		(*s2)[i], (*s2)[j] = (*s2)[j], (*s2)[i]
		i++
		j--
	}
}

func main() {
	var s string
	s = "the cake is sweet"
	fmt.Println(wordsReverse(s))
}
