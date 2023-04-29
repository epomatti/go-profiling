package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

type Customer struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

func main() {
	startTime := time.Now()

	var customers []*Customer

	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Could not create profile file: ", err)
		return
	}
	defer f.Close()

	for i := 0; i < 1000000; i++ {
		c := Customer{
			ID:      i,
			Name:    "John Doe",
			Address: "123 Main St",
			Phone:   "555-555-5555",
		}
		customers = append(customers, &c)

		// Pass the customers slice to a function for processing
		processCustomers(customers)

		// periodically clear the slice and allow it to be garbage collected
		if i%1000 == 0 {
			customers = make([]*Customer, 0)
		}
	}

	lastCustomerID := customers[len(customers)-1].ID
	fmt.Printf("ID of the last customer added: %d\n", lastCustomerID)

	pprof.WriteHeapProfile(f)

	endTime := time.Now()
	fmt.Printf("Total time taken: %v\n", endTime.Sub(startTime))
}

// processCustomers takes in a slice of Customer pointers and performs some complex logic on it
func processCustomers(customers []*Customer) {
	for range customers {
		// Do Super complex processes
		// like inserting into database
		// process for data cleanliness
		// currently empty
	}
}
