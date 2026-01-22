package main

func main() {
	// SquareOfSortedArray finds the squares of each number in a sorted array
	// and returns a new array with the squares sorted in non-decreasing order.
	nums := []int{-4, -1, 0, 3, 10}
	result := SquareOfSortedArray(nums)
	for _, v := range result {
		println(v)
	}
}

func SquareOfSortedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
