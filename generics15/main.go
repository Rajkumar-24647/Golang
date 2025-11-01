package main

import (
	"fmt"
	
)

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printSliceString(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//! here we have to repeat the code for int and string slice.But the code inside the functions are same. To remove code duplication...
//? here we can pass nay type  of slice
// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//? to pass specific types
// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
//? now we dont't need of code  duplication	




type Stack[T any] struct {
	elements []T
}

func main() {

	// nums := []int{1,2,3}
	// printSlice(nums)

	// names := []string{"go", "ts", "js"}
	// printSlice(names)


	MyStack := Stack[string]{
		elements: []string{
			"go",
			"ts",
			"js",                        
		},
	} 
	fmt.Println(MyStack)


}
