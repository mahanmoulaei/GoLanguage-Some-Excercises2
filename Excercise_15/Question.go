package main

import "fmt"

func Decoding(s string) int {
	count := make([]int, len(s)+1)
	count[0], count[1] = 1, 1

	// if frist char is 0 (no letter represents for 0)
	if s[0] == '0' {
		return 0
	}

	for i := 2; i <= len(s); i++ {

		// 0 shouldn't be valid but if previous numb is greater than 0, show as valid (ex: 10, 20)
		// For numbs < 10
		if s[i-1] > '0' {
			count[i] = count[i-1]
		}

		// cannot be over 26, s i-2 check for 10-19, (s[i-2] == '2' && s[i-1] < '7') check for 20-27
		// For >= 10 numbs <=20
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			count[i] += count[i-2]
		}
	}
	return count[len(s)]
}
func main() {
	var numb string
	numb = "26"
	fmt.Println(Decoding(numb))
}
