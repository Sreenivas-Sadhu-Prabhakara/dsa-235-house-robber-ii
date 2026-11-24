package main

import "fmt"

// RobCircle is house robber on a circular street: the first and last houses
// are adjacent. Split into two linear problems — exclude the first house,
// or exclude the last — and take the better.
func RobCircle(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	return max(robLine(nums[1:]), robLine(nums[:n-1]))
}

func robLine(nums []int) int {
	robbed, skipped := 0, 0
	for _, v := range nums {
		robbed, skipped = skipped+v, max(robbed, skipped)
	}
	return max(robbed, skipped)
}

func main() {
	fmt.Println(RobCircle([]int{2, 3, 2}))
	fmt.Println(RobCircle([]int{1, 2, 3, 1}))
}
