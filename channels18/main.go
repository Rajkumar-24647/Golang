package main

import (
	// "fmt"
	
)

//! sending data in channels, ex := 1
// func processNum(numChan chan int) {

// 	for num := range numChan {

// 		fmt.Println("processing number", num)

// 		time.Sleep(time.Second)
// 	}

// }

//! receive data ex := 2
// func sum(result chan int, num1 int, num2 int) {
//  numResult := num1 + num2
//  result <- numResult
// }

//! ex := 3
// func task(done chan bool) {
// 	defer func () {
// 		done <- true
// 	}()
// 	fmt.Println("process...", )

// }


//! ex := 4
// func emailsender(emailChan chan string, done chan bool) {
// 	for email := range emailChan{
// 		fmt.Println("seanding emal to", email)
// 		time.Sleep(time.Second)
// 	}
// }

 
func main() {

	
	//! how to declare channels

	// messageChan := make(chan string)
	//! put value in channel
	//? channels are blocking
	// messageChan <- "ping"
	//! receiving value
	// msg := <-messageChan
	// fmt.Println(msg)



    //! ex := 1
	// numChan := make(chan int)
	// go processNum(numChan)
	// numChan <- 5
	// time.Sleep(time.Second * 2)

	// numChan := make(chan int)
	// go processNum(numChan)
	// numChan <- 5
	// for {
	// 	numChan <- rand.IntN(100)
	// }


  //! ex := 2
	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <- result
	// fmt.Println(res)


    //!ex := 3
	// done := make(chan bool)
	// go task(done) 
	// <- done //block



	//! ex := 4
	// emailChan := make(chan string, 100)
	
	// done := make(chan bool)

	// go emailsender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending")

	// close(emailChan)

	// <- done

	// emailChan <- "raj@ex.com"
	// emailChan <- "kumar@ex.com"

	// fmt.Println(<- emailChan)
	// fmt.Println(<- emailChan)





	//! receiving data from multiple channels
	// chan1 := make(chan int)
	// chan2 := make(chan string)

	// go func(){
	// 	chan1 <- 10
	// }()
	// go func(){
	// 	chan2 <- "hey!"
	// }()


	// for i:=0; i<2; i++ {
	// 	select {
	// 	case chan1Val := <- chan1:
	// 		fmt.Println("received data from chan1", chan1Val)
	// 	case chan2Val := <- chan2:
	// fmt.Println("received data from chan2", chan2Val)
	// 	}
	// }
	
	     




}
