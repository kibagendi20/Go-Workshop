package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string `json:"n"`
	Age    int    `json:"a"`
	Salary int    `json:"s"`
}

func main() {
	emp := &employee{
		Name:   "Kibagendi",
		Age:    30,
		Salary: 50000,
	}

	empJSON, err := json.MarshalIndent(emp, "", "")

	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))
}
