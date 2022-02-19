package main

import (
	"fmt"
	"strconv"
)

func main() {

	var array []int
	var number int

	fmt.Println("Enter the list of the money for each house (no more than 1000$, no less than 9$)")
	for {
		_, err := fmt.Scanln(&number)

		if number < 0 || number > 1000 || err != nil {
			fmt.Println("Insertion stoped...")
			break
		} else {
			array = append(array, number)
		}
	}

	if len(array) < 1 {
		fmt.Println("There is nothing to rob")
		return
	}

	fmt.Println("Tonight you have rob in total of $" + strconv.Itoa(rob(array)))
}

func rob(array []int) int {
	lastHM := array
	var robAmount int = 0
	var robcnt int = 0

	if len(lastHM)%2 == 0 {
		for i := 0; i < len(lastHM); i++ {
			robAmount += lastHM[i]
			robcnt++
			i++
		}
	} else {
		for i := 1; i < len(lastHM); i++ {
			robAmount += lastHM[i]
			robcnt++
			i++
		}
	}

	var C int = 0
	if robcnt == 1 {
		for i := 0; i < len(lastHM); i++ {
			if C < lastHM[i] {
				C = lastHM[i]
			}
		}
	}

	if C > robAmount {
		robAmount = C
	}

	return robAmount
}
