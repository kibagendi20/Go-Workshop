package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	emp := employee{
		name:   "Wesley",
		age:    20,
		salary: 200,
	}

	empP := &emp
	fmt.Printf("Emp: %+v\n", empP)
	empP = &employee{name: "Enock", age: 25, salary: 2000}
	fmt.Printf("Emp: %+v\n", empP)
}
