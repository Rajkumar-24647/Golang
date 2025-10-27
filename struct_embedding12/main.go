package main

import (
	"fmt"
	"time"
)


//! we can use one struct into another struct

type customer struct {
	name string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {

	newCustomer := customer {
		name: "raj",
		phone: "66461654894156",
	}

	newOrder := order{
		id: "1",
		amount: 30,
		status: "received",
		customer: newCustomer,
	}

	newOrder.customer.name = "swati"

	fmt.Println(newOrder)

}