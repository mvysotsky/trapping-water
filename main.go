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
	size := len(levels)

	if size < 3 {
		return 0
	}

	last := size - 1
	maxL := levels[0]
	maxR := levels[last]
	posL := 1
	posR := last - 1

	for posL <= posR {
		opportunity := 0
		if maxL <= maxR {
			opportunity = maxL - levels[posL]
			if levels[posL] > maxL {
				maxL = levels[posL]
			}
			posL++
		} else {
			opportunity = maxR - levels[posR]
			if levels[posR] > maxR {
				maxR = levels[posR]
			}
			posR--
		}

		if opportunity > 0 {
			capacity += opportunity
		}
	}

	return capacity
}
