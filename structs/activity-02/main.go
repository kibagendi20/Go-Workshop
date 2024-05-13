package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string
	Age    int
	salary int
}

func main() {
	emp := &employee{Name: "Wes", Age: 26}
	empJSON, err := json.Marshal(emp)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Marshal funnction output %s\n", string(empJSON))

	empJSON, err = json.MarshalIndent(emp, "", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Marshal funnction output %s\n", string(empJSON))
}
