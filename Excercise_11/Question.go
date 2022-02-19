package main

import (
	"fmt"
	"sort"
)

func main() {
	var inputList = []int{-1, 0, 1, 2, -1, -4}

	var inputList2 []int

	var inputList3 = []int{0}

	CheckAndSort(inputList)

	fmt.Println("******************")

	CheckAndSort(inputList2)

	fmt.Println("******************")

	CheckAndSort(inputList3)
}

func CheckAndSort(inputList []int) {
	var inputListSorted []int
	var output [][]int
	if len(inputList) < 3 {
		PrintAnswer(inputList, output)
		return
	}

	for _, element := range inputList {
		inputListSorted = append(inputListSorted, element)
	}

	sort.Ints(inputListSorted)

	for index, element := range inputListSorted {

		if index != 0 && inputListSorted[index-1] == element {
			continue
		}

		temp := twoSum(inputListSorted[index+1:], 0-element)
		if len(temp) > 0 {
			for index2 := range temp {
				output = append(output, append(temp[index2], element))
			}
		}
	}

	PrintAnswer(inputList, output)
}

//I searched for this function on internet
func twoSum(n []int, t int) [][]int {
	var output [][]int
	left := 0
	right := len(n) - 1

	for left < right {
		tt := n[left] + n[right]
		if tt == t {
			output = append(output, []int{n[left], n[right]})
			left++
			for left < len(n) && n[left] == n[left-1] {
				left++
			}
			right--

			for right >= 0 && n[right] == n[right+1] {
				right--
			}
			continue
		}
		if tt < t {
			left++
		} else {
			right--
		}
	}
	return output
}

func PrintAnswer(inputList []int, output [][]int) {
	fmt.Println("Input: nums=", inputList)
	for index := range output {
		sort.Ints(output[index])
	}
	fmt.Println("Output:", output)
}
