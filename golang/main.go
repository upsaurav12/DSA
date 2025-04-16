package main

import "github.com/upsaurav12/dsa/array"

func main() {

	nums := []int{1, 2, 0, 3, 0, 4, 5}
	// array.PrintArray(nums)
	// max, min := array.FindMinandMaxElements(nums)
	// fmt.Println(max, min)
	// array.ReverseTheArray(nums)
	// check, asc := array.CheckifSorted(nums, false)
	// if check && asc {
	// 	fmt.Println("Array was sorted in ascending order")
	// } else if check && !asc {
	// 	fmt.Println("Array was sorted in descending order")
	// } else {
	// 	fmt.Println("Array was not sorted")
	// }

	// array.RemoveDuplicateFromSortedArray(nums)
	// fmt.Println(removed)

	// array.RotateLeft(nums, 2)

	// array.RotateRight(nums, 2)

	array.MoveAllZeroesToEnd(nums)
}
