package main

import "fmt"


//! By value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("in changeNum:", num)
// }

//! By reference
func changeNum(num *int) {
  *num = 5
  fmt.Println("In numChange", *num)
}

func main() {
	num := 1

	changeNum(&num)

	// fmt.Println("Memory address", &num)

	fmt.Println("After change", num)

}