package main

import (
	"fmt"
	
)

func main() {

	// nums := []int{6, 7, 2}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }



	// nums := []int{6, 7, 2}
	// //! here first return index and second returns us value
	// for i,num := range nums {
	// 	fmt.Println(i, num)
	// }



	// m := map[string]string{"age": "28", "price": "50"}

	// for k, val := range m {
	// 	fmt.Println(k, val)
	// }




	//? unicode code
	for i,c := range "golang"{
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}