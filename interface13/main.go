package main

import "fmt"


type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {

	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)

}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)

}

type paypal struct {}

func (p paypal) pay(amount float32) {
	fmt.Println("payment using payapl", amount)
}

func main() {

	// razorpayGw := razorpay{}
	// newPayment := payment {
	// 	gateway: razorpayGw,
	// }


	paypalGw := paypal{}
	newPayment := payment {
		gateway: paypalGw,
	}
	

	newPayment.makePayment(100)

}
