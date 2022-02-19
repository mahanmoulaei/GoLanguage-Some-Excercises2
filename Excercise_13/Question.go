package main

import "fmt"

func jump(nums []int) int {
	ret := 0
	leng := len(nums)
	max := 0  // max Index
	curr := 0 // current Index

	for i := 0; i < leng-1; i++ {
		if curr < i {
			ret++
			curr = max
		}
		if max < nums[i]+i {
			max = nums[i] + i
		}
	}
	return ret
}

func main() {
	nums := []int{2, 3, 0, 1, 4}

	fmt.Println(jump(nums))
}
