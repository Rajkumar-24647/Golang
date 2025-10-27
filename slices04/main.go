package main

import (
	"fmt"
	// "slices"
)

//! slice -> dynamic arrays
//! most used construct in go
//! useful methods

func main() {
     
	//? uninitialized slice is nil (null)
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(nums == nil)

    //? make() function by default in go. here 2 is initial size
	//? cap() -> tell us capacity
	// var nums = make([]int, 2, 5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))



	// var nums = make([]int, 0, 5)
	//?OR
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	

	//? we can write like this
	// nums := []int{}
	// fmt.Println(nums)
	// fmt.Println(cap(nums))


	//? insert by indexing
	// var nums = make([]int, 3, 5)
	// nums[0] = 2
	// nums[1] = 4
    // fmt.Println(nums)


	//? copy function
    // var nums = make([]int, 0, 5)
	// nums = append(nums, 2)

	// var nums2 = make([]int, len(nums))

	
	// copy(nums2, nums)
	

	// fmt.Println(nums, nums2)
	// fmt.Println(cap(nums))


	//? slice operator
	// var nums = []int{1, 2, 3, 5, 8}

	//! here get all the elements from 0th index to 2th index excluding 2th index
	// fmt.Println(nums[0:4])
	//! return 1,2,3,5

   //? we can also write like this
	// fmt.Println(nums[:4])

	//! print all the next elems from the 1th index including itself
	// fmt.Println(nums[1:])


	//? slice package

	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}
    //! here many properties are present in slices
	// fmt.Println(slices.Equal(nums1,  nums2))



	var nums = [][]int{{2, 3}, {4, 5}}

	fmt.Println(nums)

	
	

}