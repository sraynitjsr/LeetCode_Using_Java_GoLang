package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	current := intervals[0]

	for _, interval := range intervals[1:] {
		if interval[0] <= current[1] {
			if interval[1] > current[1] {
				current[1] = interval[1]
			}
		} else {
			merged = append(merged, current)
			current = interval
		}
	}

	merged = append(merged, current)
	return merged
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge(intervals)
	fmt.Println(result)
}
