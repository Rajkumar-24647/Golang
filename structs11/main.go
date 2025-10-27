package main

import (
	"fmt"
)

//! order struct
// type order struct {
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time  // nano second precision
// }

// func newOrder(id string, amount float32, status string) *order {
// 		myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

//! method to change status
//? receiver type
//? we have to make it pointer to change status
// func (o *order) changestatus(status string) {
// 	o.status = status
// }

//! another method to get amount
//? here we need not to make it pointer since we're getting the value not changing it
// func (o order) getAmount() float32 {
// 	return o.amount
// }


func main(){

    //! inline structs
	language := struct {
		name string
		isGood bool
	} {"golang", true}

	fmt.Println(language)




	// myOrder := newOrder("1", 40.09, "received")
	// fmt.Println(myOrder.amount)



// 	myOrder1 := order{
// 		id: "1",
// 		amount: 500.00,
// 		status: "received",	
// 	}
// 	//? now want to add createdAt
// 	myOrder1.createdAt = time.Now();
//    fmt.Println(myOrder1.status)
// 	fmt.Println("Order struct:", myOrder1)



// 	myOrder2 := order{
// 		id: "ewfkbewuivbvuiboc",
// 		amount: 65165,
// 		status: "not received",
// 		createdAt: time.Now(),
// 	}

// 	myOrder1.status = "paid"

// 	fmt.Println("myorder2:", myOrder2)
// 	fmt.Println("myorder1:", myOrder1)



//! mathod of changing status

//  	myOrder := order{
// 		id: "1",
// 		amount: 500.00,
// 		status: "received",	
// 	}
// myOrder.changestatus("confirmed")
// fmt.Println(myOrder.getAmount())




//! if  i don't set any field, default value will be zero value.
//? int => 0, float => 0, string => "", bool => false
 	// myOrder := order{
	// 	id: "1",
	// 	status: "received",	
	// }
	// fmt.Println(myOrder)




}