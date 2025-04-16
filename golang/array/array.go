package array

import (
	"fmt"
)

func FindMinandMaxElements(nums []int) (int, int) {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}

	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}

	return max, min
}

func ReverseTheArray(nums []int) []int {

	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)/2-i+1] = nums[len(nums)/2-i+1], nums[i]
	}

	fmt.Println(nums)
	return nums
}

func CheckifSorted(nums []int, asc bool) (bool, bool) {
	for i := 1; i < len(nums); i++ {
		if (asc && nums[i-1] > nums[i]) || (!asc && nums[i-1] < nums[i]) {
			return false, asc
		}
	}
	return true, asc
}

func RotateLeft(nums []int, time int) {

	//O(N)
	n := len(nums)
	time %= n
	result := append(nums[time:], nums[:time]...)
	copy(nums, result)
	fmt.Println(nums)

	/*

		O(N * time)
			count := time
			for count > 0 {
				first := nums[0]
				for i := 0; i < len(nums)-1; i++ {
					nums[i] = nums[i+1]
				}
				nums[len(nums)-1] = first
				count--
			}
			fmt.Println(nums)*/
}

func MoveAllZeroesToEnd(nums []int) {

	//O(N)
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}

	//O(N * LOG(N) + N * zeroes)
	/*
		zeroes := 0
		for _, v := range nums {
			if v == 0 {
				zeroes++
			}
		}

		sort.Ints(nums)

		count := zeroes
		for count > 0 {
			first := nums[0]
			for i := 0; i < len(nums)-1; i++ {
				nums[i] = nums[i+1]
			}
			nums[len(nums)-1] = first
			count--
		}*/
	fmt.Println(nums)
}

func RotateRight(nums []int, time int) {

	//O(N)
	n := len(nums)
	time %= n
	result := append(nums[time:], nums[:n-time]...)
	copy(nums, result)
	fmt.Println(nums)

	/*

		//O(time * N)
			count := time
			for count > 0 {
				last := nums[len(nums)-1]
				for i := len(nums) - 1; i > 0; i-- {
					nums[i] = nums[i-1]
				}
				nums[0] = last
				count--
			}
			fmt.Println(nums)*/
}

func PrintArray(nums []int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}
