package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var output []int
out:
	for i, v1 := range nums {
		for j, v2 := range nums {
			result := v1 + v2
			if result == target && i != j {
				output = append(output, i, j)
				break out
			}
		}
	}
	return output
}

func main() {
	array1 := []int{2, 7, 11, 15}
	array2 := []int{3, 2, 4}
	array3 := []int{3, 3}
	array4 := []int{0, 0, 23, 0, 101}

	fmt.Println(twoSum(array1, 9))
	fmt.Println(twoSum(array2, 6))
	fmt.Println(twoSum(array3, 6))
	fmt.Println(twoSum(array4, 23))
}
