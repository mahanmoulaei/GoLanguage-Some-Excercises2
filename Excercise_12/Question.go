package main

import (
	"fmt"
)

/*
1  2  3
4  5  6
7  8  9
*/

/*
1  2  3  4
5  6  7  8
9  10 11 12
*/
func main() {

	var inputList = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	PrintAnswer(inputList, getSpiralOrder(inputList))

	var inputList2 = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	PrintAnswer(inputList2, getSpiralOrder(inputList2))
}

func getSpiralOrder(inputList [][]int) []int {
	var output []int

	var rows int = len(inputList)

	var columns int = len(inputList[0])

	if !(rows >= 1) || !(columns <= 10) {
		return output
	}

	for index := range inputList {
		for index2 := range inputList[index] {
			if inputList[index][index2] >= -100 && inputList[index][index2] <= 100 {
				return output
			}
		}
	}

	var top int = 0
	var bottom int = rows - 1
	var left int = 0
	var right int = columns - 1

	for top <= bottom && left <= right {
		//Top
		for i := left; i <= right; i++ {
			output = append(output, inputList[top][i])
		}

		//Right
		for i := top + 1; i <= bottom; i++ {
			output = append(output, inputList[i][right])
		}

		//Bottom
		if top != bottom {
			for i := right - 1; i >= left; i-- {
				output = append(output, inputList[bottom][i])
			}
		}

		//Left
		if left != right {
			for i := bottom - 1; i > top; i-- {
				output = append(output, inputList[i][left])
			}
		}

		top++
		bottom--
		left++
		right--
	}

	return output
}

func PrintAnswer(inputList [][]int, output []int) {
	fmt.Println("Input: matrix =", inputList)
	fmt.Println("Output:", output)
	fmt.Println("******************")
}
