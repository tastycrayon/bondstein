package main

import (
	"errors"
	"fmt"
)

var (
	ErrorNotFound = errors.New("no such numbers")
)

func Soulution_1(nums []int, target int) (int, int, error) {
	var i, j int = 0, len(nums) - 1
	for i < j {
		var current int = nums[i] + nums[j]
		if current == target {
			return i, j, nil
		}
		if current < target {
			i++
		} else {
			j--
		}
	}
	return 0, 0, ErrorNotFound
}
func Soulution_2(nums []int, target int) (int, int, error) {
	visitedMap := make(map[int]int, 0)
	for i, current := range nums {
		otherPart := target - current
		if j, ok := visitedMap[otherPart]; ok {
			return i, j, nil
		}
		visitedMap[current] = i
	}
	return 0, 0, ErrorNotFound
}

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(Soulution_1(nums, 6))
	fmt.Println(Soulution_2(nums, 6))
}
