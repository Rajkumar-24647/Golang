package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	fmt.Println("doing task", id);
// }

func main() {



	for i := 0; i <= 10; i++ {
		// go task(i)

		func (i int) {
       fmt.Println(i)
		} (i)
	}



	
	//! we can stop our main function for seconds
	// time.Sleep(time.Second * 2);

}