package main

import (
	"fmt"
	"sort"
)

func sortDoubleArray(meetings [][]int) [][]int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][1] < meetings[j][1]
	})

	slc := []int{1, 2, 5, 3, 6}
	sort.Slice(slc, func(i, j int) bool {
		return slc[i] > slc[j]
	})

	fmt.Println("Slice desc sort", slc)

	// n := len(meetings)
	// for i := 0; i < n-1; i++ {
	// 	for j := 0; j < n-i-1; j++ {
	// 		if meetings[j][1] > meetings[j+1][1] {
	// 			meetings[j], meetings[j+1] = meetings[j+1], meetings[j]
	// 		}

	// 	}
	// }

	return meetings

}

func printNonOverlapping(meetings [][]int) {
	meetings = sortDoubleArray(meetings)
	count := 0
	last := -1

	for _, meeting := range meetings {
		if meeting[1] > last {
			count++
			last = meeting[1]
		}
	}

	fmt.Println("COUNT:", count)
}

func main() {
	meetings := [][]int{
		{1, 2}, {3, 4}, {1, 3}, {0, 6},

		{2, 3},
	}
	printNonOverlapping(meetings)
}
