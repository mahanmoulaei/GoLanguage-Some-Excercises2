package main

import "fmt"

func main() {
	var inputArray []int
	var number int

	fmt.Println("Enter an integer number between -231 to 230 to insert into the array. Hit 'c' to stop: ")
	for {
		_, err := fmt.Scanln(&number)

		if number < -231 || number > 230 || err != nil {
			fmt.Println("Array insertion stoped...")
			break
		} else {
			inputArray = append(inputArray, number)
		}
	}

	if len(inputArray) < 1 || len(inputArray) > 105 {
		fmt.Println("The lenght of the input array should be between 1 to 105 all together")
		return
	}

	var x int
	fmt.Println("Enter the number steps: ")
	_, err := fmt.Scan(&x)
	if x < 0 || err != nil {
		fmt.Println("An error has occured... ", err)
		return
	}
	newArray := rotate(inputArray, x)
	fmt.Println("The rotated array is: ")
	fmt.Println(newArray)
}

func rotate(array []int, y int) []int {
	var newArray []int

	for i := len(array) - y; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}

	for j := 0; j < len(array)-y; j++ {
		newArray = append(newArray, array[j])
	}

	return newArray
}
