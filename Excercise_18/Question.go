package main

import (
	"fmt"
	"strconv"
)

func main() {
	var array []int
	var number int
	fmt.Println("Enter the positive integer number between 0 to 110 (-1 to stop): ")
	for {
		_, err := fmt.Scanln(&number)

		if number < 0 || number > 110 || err != nil {
			fmt.Println("Array Creation stoped...")
			break
		} else {
			array = append(array, number)
		}
	}

	var result string

	fmt.Println("Here is the array: ", array)

	array = reverseArray(array)
	result = addNumbers(array)
	fmt.Println("Here is the arranged string: ", result)

}

func reverseArray(array []int) []int {
	for i, x := 0, len(array)-1; i < x; i, x = i+1, x-1 {
		array[i], array[x] = array[x], array[i]
	}
	return array
}

func addNumbers(array []int) string {
	var y string

	for i := 0; i < len(array); i++ {
		y += strconv.Itoa(array[i])
	}
	return y
}
