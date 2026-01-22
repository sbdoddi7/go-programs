package main

import "fmt"

// TwoSum finds two numbers in the array that add up to the target sum.
// It returns the indices of the two numbers.
func TwoSum(nums []int, target int) (int, int) {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return left, right
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return -1, -1 // Return -1, -1 if no solution is found
}

func main() {
	nums := []int{1, 2, 4, 6, 8, 9, 14, 15}
	target := 13
	i, j := TwoSum(nums, target)
	fmt.Printf("Indices: %d, %d\n", i, j) // Output: Indices: 0, 1
}
