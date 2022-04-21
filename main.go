package main

import "fmt"

func main() {
	var tests = []struct {
		arr       []int
		targetSum int
	}{
		{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10},
		{[]int{4, 6}, 10},
		{[]int{4, 6, 1}, 10},
		{[]int{4, 6, 1, -3}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 17},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18},
		{[]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5},
		{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 163},
		{[]int{-21, 301, 12, 4, 65, 56, 210, 356, 9, -47}, 164},
		{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 15},
		{[]int{14}, 15},
		{[]int{15}, 15},
		{[]int{1, 11, 10, 2}, 12},
	}

	for _, element := range tests {
		fmt.Printf("One pair: %v. All pairs: %v\n", GetPair(element.arr, element.targetSum), GetAllPairs(element.arr, element.targetSum))
	}
}

// The function returns a pair of numbers from the input array, the sum of which is equal to target sum
// In this case, the target amount cannot be obtained from the same number
func GetPair(arr []int, targetSum int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == targetSum {
				return []int{arr[i], arr[j]}
			}
		}
	}
	return []int{}
}

// The function returns all pairs of numbers from the input array, the sum of which is equal to target sum
// In this case, the target amount cannot be obtained from the same number
func GetAllPairs(arr []int, targetSum int) [][]int {
	output := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == targetSum {
				output = append(output, []int{arr[i], arr[j]})
			}
		}
	}
	return output
}
