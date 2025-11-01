package main

import "fmt"

//! enums => enumerated types

type OrderStatus int                     
const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func chnageOrderStatus(status OrderStatus){
	fmt.Println("changing order to", status)
}

func main(){

	chnageOrderStatus(Received)

}