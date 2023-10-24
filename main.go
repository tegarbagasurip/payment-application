package main 

import "payment-application/delivery"

func main () {
	delivery.Server().Run()
}

// to run the application