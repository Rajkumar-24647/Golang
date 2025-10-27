package main

import "fmt"

//? from this we can only receive the int values
func sum(nums ...int) int {
     
	total := 0
	for _, num := range nums {
		total = total + num
	}

	return  total
}

//? here we can assing any type.
// func sum(nums ...interface{}) int {
     
// 	total := 0
// 	for _, num := range nums {
		
// 	}

// 	return  total
// }

func main() {


	nums := []int{3, 4, 5, 6}
	result := sum(nums...)

	fmt.Println(result)



	// result := sum(4, 6, 2, 9)
	// fmt.Println(result)

}