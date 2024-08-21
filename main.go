package main

import (
	"fmt"
	"log"
)

func main() {
	// Test cases from the task
	assert(estimate([]int{33, 33, 5, 7}), 2, "Test case {33, 33, 5, 7} failed")
	assert(estimate([]int{1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6, "Test case {1,0,2,1,0,1,3,2,1,2,1} failed")
	assert(estimate([]int{0, 3, 0, 2, 0, 4}), 7, "Test case {0,3,0,2,0,4} failed")
	assert(estimate([]int{0, 2, 0, 0, 2, 0}), 4, "Test case {0,2,0,0,2,0} failed")

	// Additional test cases
	assert(estimate([]int{1, 2, 3, 4, 5}), 0, "Test case {1,2,3,4,5} failed")
	assert(estimate([]int{5, 4, 3, 2, 1}), 0, "Test case {5,4,3,2,1} failed")
	assert(estimate([]int{1, 2, 3, 2, 1}), 0, "Test case {1,2,3,2,1} failed")
	assert(estimate([]int{1, 2, 3, 2, 3, 2, 1}), 1, "Test case {1,2,3,2,3,2,1} failed")
	assert(estimate([]int{1, 2, 3, 2, 3, 2, 3, 2, 1}), 2, "Test case {1,2,3,2,3,2,3,2,1} failed")

	// Edge cases
	assert(estimate([]int{}), 0, "Test case {} failed")
	assert(estimate([]int{1}), 0, "Test case {1} failed")
	assert(estimate([]int{1, 2}), 0, "Test case {1,2} failed")
	assert(estimate([]int{1, 2, 3}), 0, "Test case {1,2,3} failed")

	log.Print("All test cases passed")
}

func assert(a, b int, err string) {
	if a != b {
		panic(fmt.Sprintf("%s: expected %d, but got %d", err, b, a))
	}
}

func estimate(levels []int) (capacity int) {
	if len(levels) < 3 {
		// could not store anything
		return 0
	}

	for pos := 0; pos+1 < len(levels); {
		if levels[pos+1] >= levels[pos] {
			pos++
			continue
		}

		add := 0
		pos, add = store(levels, pos)
		capacity += add
	}

	return capacity
}

func store(levels []int, startPos int) (pos int, totalCapacity int) {
	var currPeakPos = -1

	for pos = startPos; pos+1 < len(levels); pos++ {
		if levels[pos+1] >= levels[startPos] {
			// area closed
			currPeakPos = pos + 1
			break
		}

		if levels[pos+1] > levels[pos] {
			if currPeakPos < 0 || (levels[pos+1] > levels[currPeakPos]) {
				currPeakPos = pos + 1
			}
		}
	}

	if currPeakPos < 0 {
		return pos, 0
	}

	maxHeight := min(levels[startPos], levels[currPeakPos])

	for pos = startPos; pos <= currPeakPos; pos++ {
		if levels[pos] < maxHeight {
			totalCapacity += maxHeight - levels[pos]
		}
	}

	return currPeakPos, totalCapacity
}
