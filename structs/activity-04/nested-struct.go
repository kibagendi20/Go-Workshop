package main

import (
	"fmt"
)

type employee struct {
	name    string
	age     int
	salary  int
	address address
}
type address struct {
	city    string
	country string
}

func main() {
	address := address{city: "London", country: "UK"}
	emp := employee{name: "Sam", age: 31, salary: 2000, address: address}

	fmt.Printf("City: %s\n", emp.address.city)
	fmt.Printf("Country: %s\n", emp.address.country)

	//fmt.Printf("City: %s\n", emp.city)
}
